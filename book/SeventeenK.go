package book

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/yizenghui/sda/code"
	"github.com/yizenghui/sda/data"
)

//SeventeenK 17K
type SeventeenK struct {
	UpdateListURL string
	BookInfoURL   string
	FansRankURL   string
}

//GetUpdate 17K
func (s *SeventeenK) GetUpdate() ([]data.Book, error) {

	var books []data.Book
	var book data.Book
	g, e := goquery.NewDocument(s.UpdateListURL)
	if e != nil {
		return books, e
	}

	// 下列内容于 2017年4月4日 20:50:24 抓取
	g.Find("table tbody tr").Each(func(i int, content *goquery.Selection) {
		// 书名
		book.Name = strings.TrimSpace(content.Find(".td3").Find(".jt").Text())
		// tr有空行
		if book.Name != "xxxx" {

			// 书籍地址
			book.BookURL, _ = content.Find(".td3").Find(".jt").Attr("href")
			// 章节
			book.Chapter = strings.TrimSpace(content.Find(".td4").Find("a").Eq(0).Text())
			// 章节地址
			book.ChapterURL, _ = content.Find(".td4").Find("a").Attr("href")

			// 作者名
			book.Author = strings.TrimSpace(content.Find(".td6").Text())
			// 作者详细页
			book.AuthorURL, _ = content.Find(".td6").Find("a").Attr("href")

			// 字数
			book.Total = strings.TrimSpace(content.Find(".td5").Text())

			// 更新时间
			book.Date = strings.TrimSpace(content.Find(".td7").Text())

			checkIsVIP, _ := content.Find(".td4").Find(".vip").Attr("title")
			if checkIsVIP != "" {
				book.IsVIP = true
			} else {
				book.IsVIP = false
			}

			books = append(books, book)
		}
	})

	return books, nil
}

// GetInfo 获取书籍基础信息(与列表一致)
func (s *SeventeenK) GetInfo() (data.Book, error) {

	var book data.Book
	g, e := goquery.NewDocument(s.BookInfoURL)
	if e != nil {
		return book, e
	}

	html, _ := g.Html()

	book.BookURL = s.BookInfoURL

	// 页面上的更新时间
	dateText := strings.TrimSpace(code.FindString(`<em>更新：(?P<date>[^<]+)</em>`, html, "date"))

	vipChapterMap := code.SelectString(`<span class="time">更新时间：(?P<time>[^<]+)</span>(?P<_>\s+)最新vip章节：<a(?P<_>\s+)href="(?P<url>[^"]+)"(?P<_>\s+)target="_blank">(?P<chapter>[^<]+)</a>`, html)

	chapterMap := code.SelectString(`<span class="time">更新时间：(?P<time>[^<]+)</span>(?P<_>\s+)最新免费章节：<a(?P<_>\s+)href="(?P<url>[^"]+)"(?P<_>\s+)target="_blank">(?P<chapter>[^<]+)</a>`, html)

	// 普通章节更新时间需要与页面更新时间一致
	if t, ok := chapterMap["time"]; ok && t == dateText {

		book.Chapter, _ = chapterMap["chapter"]
		book.ChapterURL, _ = chapterMap["url"]
		book.Date = t
		// fmt.Println(chapterMap)
	} else if t, ok := vipChapterMap["time"]; ok { // && t == dateText
		//  && t == dateText

		book.Chapter, _ = vipChapterMap["chapter"]
		book.ChapterURL, _ = vipChapterMap["url"]
		book.Date = t
		// fmt.Println(vipChapterMap)
	}

	//
	// 书名
	book.Name = strings.TrimSpace(g.Find("h1").Find("a").Text())
	// 作者
	book.Author = strings.TrimSpace(g.Find(".AuthorInfo").Find(".name").Text())
	// 作者页
	book.AuthorURL, _ = g.Find(".AuthorInfo").Find(".name").Attr("href")

	// 字数
	book.Total = strings.TrimSpace(g.Find(".BookData").Find(".red").Text())

	// 如果含有VIP章节字眼，为VIP作品
	checkIsVIP, _ := regexp.MatchString(`最新v(?P<vip>\w+)章节：`, html)
	if checkIsVIP {
		book.IsVIP = true
	} else {
		book.IsVIP = false
	}
	return book, nil
}

// GetFans 获取书籍红包信息
func (s *SeventeenK) GetFans() ([]data.Fans, error) {
	// id
	var rows []data.Fans
	var fans data.Fans
	g, e := goquery.NewDocument(s.FansRankURL)
	if e != nil {
		return rows, e
	}
	g.Find(".PiaoyouTop li").Each(func(i int, content *goquery.Selection) {
		// 粉丝名
		fans.Name = strings.TrimSpace(content.Find("a").Eq(0).Text())

		fans.URL, _ = content.Find("a").Eq(0).Attr("href")

		money := content.Find(".more").Eq(0).Text()

		i64, err := strconv.ParseInt(money, 10, 16)
		if err != nil {
			// fmt.Println(err)
		}

		i16 := int16(i64)

		switch {
		case i16 >= 10000:
			fans.Level = 5
		case i16 >= 1000:
			fans.Level = 4
		case i16 >= 100:
			fans.Level = 3
		case i16 >= 10:
			fans.Level = 2
		default:
			fans.Level = 1
		}

		rows = append(rows, fans)
	})
	return rows, nil
}
