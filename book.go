package sda

import (
	"fmt"
	"regexp"

	"github.com/yizenghui/sda/book"
	"github.com/yizenghui/sda/code"
	"github.com/yizenghui/sda/data"
)

//GetUpdateBookByListURL 获取更新列表内容 支持起点、纵横、17K
func GetUpdateBookByListURL(url string) ([]data.Book, error) {

	// 起点列表
	checkLinkIsQiDian, _ := regexp.MatchString(`http:\/\/a.qidian.com\/\?orderId=5&page=(?P<p>\d+)&style=2`, url)
	if checkLinkIsQiDian {
		// fmt.Println("checkLinkIsQiDian", checkLinkIsQiDian)
		qidian := book.QiDian{UpdateListURL: url}
		return qidian.GetUpdate()
	}

	// 纵横男生网
	checkLinkIsZongHeng, _ := regexp.MatchString(`http:\/\/book.zongheng.com\/store\/c0\/c0\/b0\/u0\/p(?P<p>\d+)\/v9\/s9\/t0\/ALL.html`, url)
	if checkLinkIsZongHeng {
		// fmt.Println("checkLinkIsZongHeng", checkLinkIsZongHeng)
		zongheng := book.ZongHeng{UpdateListURL: url}
		return zongheng.GetUpdate()
	}

	//17K
	checkLinkIsSeventeenK, _ := regexp.MatchString(`http:\/\/all.17k.com\/lib\/book\/(?P<p>[0-9_]+).html`, url)
	if checkLinkIsSeventeenK {
		// fmt.Println("checkLinkIsSeventeenK", checkLinkIsSeventeenK)
		sk := book.SeventeenK{UpdateListURL: url}
		return sk.GetUpdate()
	}

	var books []data.Book
	return books, nil
}

//FindBookBaseByBookURL 通过书籍URL信息获取书籍的基本信息
func FindBookBaseByBookURL(url string) (data.Book, error) {
	// 检查修正URL
	url = code.ExplainBookDetailedAddress(url)
	// 该url能够匹配到内容
	if url != "" {
		// 起点列表
		checkLinkIsQiDian, _ := regexp.MatchString(`book.qidian.com\/info\/(?P<book_id>\d+)`, url)
		if checkLinkIsQiDian {
			// fmt.Println("checkLinkIsQiDian", checkLinkIsQiDian)
			qidian := book.QiDian{BookInfoURL: url}
			return qidian.GetInfo()
		}

		// 纵横男生网
		checkLinkIsZongHeng, _ := regexp.MatchString(`book.zongheng.com\/book\/(?P<book_id>\d+).html`, url)
		if checkLinkIsZongHeng {
			// fmt.Println("checkLinkIsZongHeng", checkLinkIsZongHeng)
			zongheng := book.ZongHeng{BookInfoURL: url}
			return zongheng.GetInfo()
		}

		//17K
		checkLinkIsSeventeenK, _ := regexp.MatchString(`www.17k.com\/book\/(?P<book_id>\d+).html`, url)
		if checkLinkIsSeventeenK {
			// fmt.Println("checkLinkIsSeventeenK", checkLinkIsSeventeenK)
			sk := book.SeventeenK{BookInfoURL: url}
			return sk.GetInfo()
		}

	}
	var book data.Book
	return book, nil
}

//FindBookFansByBookURL 通过书籍URL信息获取书籍的基本信息
func FindBookFansByBookURL(url string) ([]data.Fans, error) {
	// 检查修正URL
	url = code.ExplainBookDetailedAddress(url)
	// 该url能够匹配到内容
	if url != "" {
		// 起点列表
		checkLinkIsQiDian, _ := regexp.MatchString(`book.qidian.com\/info\/(?P<book_id>\d+)`, url)
		if checkLinkIsQiDian {
			Map := code.SelectString(`book.qidian.com\/info\/(?P<book_id>\d+)`, url)
			qidian := book.QiDian{FansRankURL: fmt.Sprintf("http://book.qidian.com/fansrank/%v", Map["book_id"])}
			return qidian.GetFans()
		}

		// 纵横男生网
		checkLinkIsZongHeng, _ := regexp.MatchString(`book.zongheng.com\/book\/(?P<book_id>\d+).html`, url)
		if checkLinkIsZongHeng {
			// fmt.Println("checkLinkIsZongHeng", checkLinkIsZongHeng)
			Map := code.SelectString(`book.zongheng.com\/book\/(?P<book_id>\d+).html`, url)
			zongheng := book.ZongHeng{FansRankURL: fmt.Sprintf("http://book.zongheng.com/donate/%v.html", Map["book_id"])}
			return zongheng.GetFans()
		}

		//17K
		checkLinkIsSeventeenK, _ := regexp.MatchString(`www.17k.com\/book\/(?P<book_id>\d+).html`, url)
		if checkLinkIsSeventeenK {
			// fmt.Println("checkLinkIsSeventeenK", checkLinkIsSeventeenK)
			Map := code.SelectString(`www.17k.com\/book\/(?P<book_id>\d+).html`, url)
			sk := book.SeventeenK{FansRankURL: fmt.Sprintf("http://www.17k.com/book/voteRank.action?bookId=%v&type=hb", Map["book_id"])}
			return sk.GetFans()
		}

	}
	var fans []data.Fans
	return fans, nil
}
