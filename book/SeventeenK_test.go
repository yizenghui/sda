// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package book

import (
	"fmt"
	"testing"

	"github.com/yizenghui/sda/data"
)

func Test_GetSeventeenKUpdate(t *testing.T) {

	var url string
	var books []data.Book

	url = "http://all.17k.com/lib/book/2_0_0_0_0_0_0_0_1.html"
	s := SeventeenK{UpdateListURL: url}
	books, _ = s.GetUpdate()
	fmt.Println(books)

}

func Test_GetSeventeenKInfo(t *testing.T) {
	var book data.Book
	var url string
	url = "http://www.17k.com/book/2573233.html"
	sk := SeventeenK{BookInfoURL: url}
	book, _ = sk.GetInfo()
	fmt.Println(book)
}
