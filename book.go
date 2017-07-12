package sda

import (
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
	url = code.ExplainDetailedAddress(url)
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
