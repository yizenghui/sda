package code

import (
	"fmt"
	"regexp"
)

// ExplainBookDetailedAddress 将书籍地址解释为书籍详细地址(小说首页)
func ExplainBookDetailedAddress(url string) string {

	// 检查是不是起点地址
	if checkLinkIsQiDian, _ := regexp.MatchString(`qidian.com`, url); checkLinkIsQiDian {
		// 起点详细页
		//http://book.qidian.com/info/1004608738
		InfoBook := `book.qidian.com\/info\/(?P<book_id>\d+)`
		if b, _ := regexp.MatchString(InfoBook, url); b {
			Map := SelectString(InfoBook, url)
			return fmt.Sprintf("http://book.qidian.com/info/%v", Map["book_id"])
		}

		// 起点手机详细页
		//http://m.qidian.com/book/1004608738
		MobileBook := `m.qidian.com\/book\/(?P<book_id>\d+)`
		if b, _ := regexp.MatchString(MobileBook, url); b {
			Map := SelectString(MobileBook, url)
			return fmt.Sprintf("http://book.qidian.com/info/%v", Map["book_id"])
		}

		// 起点手机章节列表页
		//http://m.qidian.com/book/1004608738/catalog
		MobileBookChapterMenu := `m.qidian.com\/book\/(?P<book_id>\d+)\/catalog`
		if b, _ := regexp.MatchString(MobileBookChapterMenu, url); b {
			Map := SelectString(MobileBookChapterMenu, url)
			return fmt.Sprintf("http://book.qidian.com/info/%v", Map["book_id"])
		}

		// 手机章节详细页
		//http://m.qidian.com/book/1004608738/342363924
		MobileBookChapterInfo := `m.qidian.com\/book\/(?P<book_id>\d+)\/(?P<chapter_id>\d+)`
		if b, _ := regexp.MatchString(MobileBookChapterInfo, url); b {
			Map := SelectString(MobileBookChapterInfo, url)
			return fmt.Sprintf("http://book.qidian.com/info/%v", Map["book_id"])
		}

		BookVIPChapterInfo := `vipreader.qidian.com\/chapter\/(?P<book_id>\d+)\/(?P<chapter_id>\d+)`
		if b, _ := regexp.MatchString(BookVIPChapterInfo, url); b {
			Map := SelectString(BookVIPChapterInfo, url)
			return fmt.Sprintf("http://book.qidian.com/info/%v", Map["book_id"])
		}

		// todo http://read.qidian.com/chapter/_AaqI-dPJJ4uTkiRw_sFYA2/-Yjl2ADCXQvM5j8_3RRvhw2
	}

	// 检查是不是纵横地址
	if checkLinkIsZongHeng, _ := regexp.MatchString(`zongheng.com`, url); checkLinkIsZongHeng {
		// 纵横详细页
		// http://book.zongheng.com/book/490607.html
		InfoBook := `book.zongheng.com\/book\/(?P<book_id>\d+).html`
		if b, _ := regexp.MatchString(InfoBook, url); b {
			Map := SelectString(InfoBook, url)
			return fmt.Sprintf("http://book.zongheng.com/book/%v.html", Map["book_id"])
		}

		// 纵横手机详细页
		// http://m.zongheng.com/h5/book?bookid=490607
		MobileBook := `m.zongheng.com\/h5\/book\?bookid=(?P<book_id>\d+)`
		if b, _ := regexp.MatchString(MobileBook, url); b {
			Map := SelectString(MobileBook, url)
			return fmt.Sprintf("http://book.zongheng.com/book/%v.html", Map["book_id"])
		}

		// 纵横手机章节列表页
		// http://m.zongheng.com/h5/chapter/list?bookid=490607
		MobileBookChapterMenu := `m.zongheng.com\/h5\/chapter\/list\?bookid=(?P<book_id>\d+)`
		if b, _ := regexp.MatchString(MobileBookChapterMenu, url); b {
			Map := SelectString(MobileBookChapterMenu, url)
			return fmt.Sprintf("http://book.zongheng.com/book/%v.html", Map["book_id"])
		}

		// 起点手机章节列表页
		//http://m.qidian.com/book/1004608738/342363924
		// http://m.zongheng.com/h5/chapter?bookid=490607&cid=8134632
		MobileBookChapterInfo := `m.zongheng.com\/h5\/chapter\?bookid=(?P<book_id>\d+)&cid=(?P<chapter_id>\d+)`
		if b, _ := regexp.MatchString(MobileBookChapterInfo, url); b {
			Map := SelectString(MobileBookChapterInfo, url)
			return fmt.Sprintf("http://book.zongheng.com/book/%v.html", Map["book_id"])
		}
	}

	// 检查是不是17k地址
	if checkLinkIsSeventeenK, _ := regexp.MatchString(`17k.com`, url); checkLinkIsSeventeenK {
		// 17k详细页
		InfoBook := `17k.com\/book\/(?P<book_id>\d+).html`
		if b, _ := regexp.MatchString(InfoBook, url); b {
			Map := SelectString(InfoBook, url)
			return fmt.Sprintf("http://www.17k.com/book/%v.html", Map["book_id"])
		}

		// 章节列表
		// 17k.com/book/2317974.html
		BookChapterMenu := `17k.com\/list\/(?P<book_id>\d+).html`
		if b, _ := regexp.MatchString(BookChapterMenu, url); b {
			Map := SelectString(BookChapterMenu, url)
			return fmt.Sprintf("http://www.17k.com/book/%v.html", Map["book_id"])
		}

		// 章节详细
		MobileBookChapterInfo := `17k.com\/chapter\/(?P<book_id>\d+)/(?P<chapter_id>\d+).html`
		if b, _ := regexp.MatchString(MobileBookChapterInfo, url); b {
			Map := SelectString(MobileBookChapterInfo, url)
			return fmt.Sprintf("http://www.17k.com/book/%v.html", Map["book_id"])
		}
	}

	return ""
}
