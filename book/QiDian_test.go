// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package book

import (
	"fmt"
	"testing"

	"github.com/yizenghui/sda/data"
)

func Test_GetQiDianUpdate(t *testing.T) {
	var url string
	var books []data.Book
	url = "http://a.qidian.com/?orderId=5&page=1&style=2"
	q := QiDian{UpdateListURL: url}
	books, _ = q.GetUpdate()
	fmt.Println(books)
}

func Test_GetQiDianInfo(t *testing.T) {
	var book data.Book
	var url string
	url = "http://book.qidian.com/info/1004608738"
	q := QiDian{BookInfoURL: url}
	book, _ = q.GetInfo()
	fmt.Println(book)
}
