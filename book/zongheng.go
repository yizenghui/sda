package book

import (
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/yizenghui/sda/code"
	"github.com/yizenghui/sda/data"
)

//ZongHeng 纵横
type ZongHeng struct {
	UpdateListURL string
	BookInfoURL   string
	FansRankURL   string
}

// ZongHengFansLevelMate 纵横小说网粉丝级别映射
var ZongHengFansLevelMate = map[string]int16{
	"http://static.zongheng.com/v2012/images/star_16.gif": 5,

	"http://static.zongheng.com/v2012/images/star_13.gif": 4,
	"http://static.zongheng.com/v2012/images/star_12.gif": 4,
	"http://static.zongheng.com/v2012/images/star_11.gif": 4,

	"http://static.zongheng.com/v2012/images/star_10.gif": 3,
	"http://static.zongheng.com/v2012/images/star_9.gif":  3,
	"http://static.zongheng.com/v2012/images/star_8.gif":  3,
	"http://static.zongheng.com/v2012/images/star_7.gif":  3,
	"http://static.zongheng.com/v2012/images/star_6.gif":  3,
	"http://static.zongheng.com/v2012/images/star_5.gif":  3,

	"http://static.zongheng.com/v2012/images/star_4.gif": 2,
	"http://static.zongheng.com/v2012/images/star_3.gif": 2,
	"http://static.zongheng.com/v2012/images/star_2.gif": 2,

	"http://static.zongheng.com/v2012/images/star_1.gif": 1,
}

//GetUpdate 纵横
func (z *ZongHeng) GetUpdate() ([]data.Book, error) {

	var books []data.Book
	var book data.Book
	g, e := goquery.NewDocument(z.UpdateListURL)
	if e != nil {
		return books, e
	}

	// 下列内容于
	g.Find(".main_con li").Each(func(i int, content *goquery.Selection) {
		// 书名
		book.Name = strings.TrimSpace(content.Find(".chap").Find(".fs14").Text())
		// li有空行
		if book.Name != "" {

			// 书籍地址
			book.BookURL, _ = content.Find(".chap").Find(".fs14").Attr("href")
			// 章节
			book.Chapter = strings.TrimSpace(content.Find(".chap").Find("a").Eq(1).Text())
			// 章节地址
			book.ChapterURL, _ = content.Find(".chap").Find("a").Eq(1).Attr("href")

			// 作者名
			book.Author = strings.TrimSpace(content.Find(".author").Text())
			// 作者详细页
			book.AuthorURL, _ = content.Find(".author").Find("a").Attr("href")

			// 字数
			book.Total = strings.TrimSpace(content.Find(".number").Text())

			// 更新时间
			book.Date = strings.TrimSpace(content.Find(".time").Text())

			checkIsVIP, _ := content.Find(".chap").Find(".vip").Attr("title")
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
func (z *ZongHeng) GetInfo() (data.Book, error) {

	var book data.Book
	g, e := goquery.NewDocument(z.BookInfoURL)
	if e != nil {
		return book, e
	}

	html, _ := g.Html()

	book.BookURL = z.BookInfoURL
	book.ChapterURL, _ = g.Find(".update").Find(".cont").Find("a").Attr("href")

	chapterContent, _ := g.Find(".update").Find(".cont").Find("a").Html()

	chapterName := strings.TrimSpace(code.FindString(`：(?P<chapter>[^<]+)<p>`, chapterContent, "chapter"))
	if chapterName != "" {
		book.Chapter = chapterName
	} else {
		book.Chapter = g.Find(".update").Find(".cont").Find("a").Text()
	}

	dateText := strings.TrimSpace(g.Find(".update").Find(".uptime").Text())

	date := strings.TrimSpace(code.FindString(`·(?P<date>[^<]+)·`, dateText, "date"))
	if date != "" {
		book.Date = date
	} else {
		book.Date = dateText
	}

	//
	// 书名
	book.Name = code.FindString(`<meta name="og:novel:book_name" content="(?P<name>[^"]+)"/>`, html, "name")
	// 作者
	book.Author = code.FindString(`<meta name="og:novel:author" content="(?P<author>[^"]+)"/>`, html, "author")
	// 作者
	book.AuthorURL = code.FindString(`<meta name="og:novel:author_link" content="(?P<author_url>[^"]+)"/>`, html, "author_url")

	book.Total = code.FindString(`<em>·</em>字数：<span title="(?P<t>\d+)字">(?P<total>\d+)</span>字`, html, "total")

	checkIsVIP, _ := regexp.MatchString(`<em class="(?P<vip>\w+)" title="VIP作品"></em>`, html)
	if checkIsVIP {
		book.IsVIP = true
	} else {
		book.IsVIP = false
	}
	return book, nil
}

// GetFans 获取前100粉丝级别
func (z *ZongHeng) GetFans() ([]data.Fans, error) {
	// id
	var rows []data.Fans
	var fans data.Fans
	g, e := goquery.NewDocument(z.FansRankURL)
	if e != nil {
		return rows, e
	}
	g.Find(".tabcontainer li").Each(func(i int, content *goquery.Selection) {
		// 书名
		fans.Name = strings.TrimSpace(content.Find("a").Eq(0).Text())

		fans.URL, _ = content.Find("a").Eq(0).Attr("href")

		fansLevelImg, _ := content.Find("img").Eq(0).Attr("src")

		if l, ok := ZongHengFansLevelMate[fansLevelImg]; ok {
			fans.Level = l
		}
		rows = append(rows, fans)
	})
	return rows, nil
}
