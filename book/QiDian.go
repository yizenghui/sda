package book

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/yizenghui/sda/code"
	"github.com/yizenghui/sda/data"
)

//QiDian 起点
type QiDian struct {
	UpdateListURL string
	BookInfoURL   string
}

// NewQiDian 起点资源
func NewQiDian(url string) QiDian {

	qidian := QiDian{}
	// 起点列表
	checkLinkIsQiDianList, _ := regexp.MatchString(`a.qidian.com\/\?orderId=5&page=(?P<p>\d+)&style=2`, url)
	if checkLinkIsQiDianList {
		qidian.UpdateListURL = url
	}

	// 起点详细页
	regInfoBook := `book.qidian.com\/info\/(?P<book_id>\d+)`
	checkLinkIsQiDianInfo, _ := regexp.MatchString(regInfoBook, url)
	if checkLinkIsQiDianInfo {
		Map := code.SelectString(regInfoBook, url)
		qidian.BookInfoURL = fmt.Sprintf("http://book.qidian.com/info/%v", Map["book_id"])
	}

	// regQiDianChapter := `\/\/book.qidian.com\/info\/(?P<book_id>\d+)`
	// checkLinkIsQiDian, _ := regexp.MatchString(repQiDian, url)
	// if checkLinkIsQiDian {
	// 	Map := code.SelectString(repQiDian, url)
	// 	qidian.BookInfoURL = fmt.Sprintf("http://book.qidian.com/info/%v", Map["book_id"])
	// }

	return qidian
}

//GetUpdate 起点
func (q *QiDian) GetUpdate() ([]data.Book, error) {
	var books []data.Book
	var book data.Book
	g, e := goquery.NewDocument(q.UpdateListURL)
	if e != nil {
		return books, e
	}

	// 下列内容于 2017年4月4日 20:50:24 抓取
	g.Find(".rank-table-list tbody tr").Each(func(i int, content *goquery.Selection) {
		// 书详细页
		book.BookURL, _ = content.Find(".name").Attr("href")

		book.ChapterURL, _ = content.Find(".chapter").Attr("href")
		// 书名
		book.Name = strings.TrimSpace(content.Find(".name").Text())
		// 章节
		book.Chapter = strings.TrimSpace(content.Find(".chapter").Text())
		// 作者
		book.Author = strings.TrimSpace(content.Find(".author").Text())
		// 作者详细页
		book.AuthorURL, _ = content.Find(".author").Attr("href")
		// 小说更新时间
		book.Date = strings.TrimSpace(content.Find(".date").Text())
		// 字数
		book.Total = strings.TrimSpace(content.Find(".total").Text())

		checkLinkIsJobInfo, _ := regexp.MatchString(`vip(?P<reader>\w+).qidian.com`, book.ChapterURL)
		if checkLinkIsJobInfo {
			book.IsVIP = true
		} else {
			book.IsVIP = false
		}

		books = append(books, book)
	})

	return books, nil
}

// GetInfo 获取书籍基础信息(与列表一致)
func (q *QiDian) GetInfo() (data.Book, error) {

	var book data.Book
	g, e := goquery.NewDocument(q.BookInfoURL)
	if e != nil {
		return book, e
	}

	id, _ := g.Find("#bookImg").Attr("data-bid")

	book.BookURL = fmt.Sprintf("//book.qidian.com/info/%v", id)

	book.ChapterURL, _ = g.Find(".book-info-detail").Find(".update").Find(".cf").Find(".blue").Attr("href")

	book.Chapter = strings.TrimSpace(g.Find(".book-info-detail").Find(".update").Find(".cf").Find(".blue").Text())

	book.Date = strings.TrimSpace(g.Find(".book-info-detail").Find(".update").Find(".cf").Find(".time").Text())

	book.Name = strings.TrimSpace(g.Find(".book-info").Find("h1").Find("em").Text())
	book.Author = strings.TrimSpace(g.Find(".book-info").Find(".writer").Text())

	book.AuthorURL, _ = g.Find(".book-info").Find(".writer").Attr("href")

	book.Total = strings.TrimSpace(g.Find(".book-info").Find("p").Eq(2).Find("em").Eq(0).Text())

	checkLinkIsJobInfo, _ := regexp.MatchString(`vip(?P<reader>\w+).qidian.com`, book.ChapterURL)
	if checkLinkIsJobInfo {
		book.IsVIP = true
	} else {
		book.IsVIP = false
	}
	return book, nil
}
