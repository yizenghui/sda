// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sda

import (
	"fmt"
	"testing"

	"github.com/yizenghui/sda/data"
)

func Test_FindBookBaseByBookURL(t *testing.T) {

	urls := []string{
		"http://vipreader.qidian.com/chapter/1004608738/377630007",
		// "http://book.qidian.com/info/1004608738?from=readfollow",
		// "http://m.qidian.com/book/1004608738/?from=readfollow",
		// "http://m.qidian.com/book/1004608738/catalog?from=readfollow",
		// "http://m.qidian.com/book/1004608738/342363924?from=readfollow",
		"http://m.zongheng.com/h5/chapter/list?bookid=490607",
		// "http://m.zongheng.com/h5/book?bookid=490607",
		// "http://book.zongheng.com/book/490607.html?from=readfollow",
		// "http://m.zongheng.com/h5/chapter?bookid=490607&cid=8134632",
		// "http://www.17k.com/chapter/2317974/27502630.html",
		// "http://h5.17k.com/chapter/2317974/27502630.html",
		// "http://www.17k.com/list/2317974.html",
		// "http://h5.17k.com/list/2317974.html",
		// "http://www.17k.com/book/2317974.html",
		"http://h5.17k.com/book/2317974.html",
	}
	for k, v := range urls {
		book, _ := FindBookBaseByBookURL(v)
		fmt.Println(k, v, book)
	}

}

func Test_GetUpdateBookByListURL(t *testing.T) {

	var url string
	var books []data.Book

	url = "http://a.qidian.com/?orderId=5&page=1&style=2"
	books, _ = GetUpdateBookByListURL(url)
	fmt.Println(books)

	url = "http://book.zongheng.com/store/c0/c0/b0/u0/p1/v9/s9/t0/ALL.html"
	books, _ = GetUpdateBookByListURL(url)
	fmt.Println(books)

	url = "http://all.17k.com/lib/book/2_0_0_0_0_0_0_0_1.html"
	books, _ = GetUpdateBookByListURL(url)
	fmt.Println(books)
}
