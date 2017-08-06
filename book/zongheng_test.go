// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package book

import (
	"fmt"
	"testing"

	"github.com/yizenghui/sda/data"
)

func Test_GetZongHengUpdate(t *testing.T) {

	var url string
	var books []data.Book

	url = "http://book.zongheng.com/store/c0/c0/b0/u0/p1/v9/s9/t0/ALL.html"
	z := ZongHeng{UpdateListURL: url}
	books, _ = z.GetUpdate()

	fmt.Println(books)

}

func Test_GetZongHengInfo(t *testing.T) {
	var book data.Book
	var url string
	url = "http://book.zongheng.com/book/490607.html"
	z := ZongHeng{BookInfoURL: url}
	book, _ = z.GetInfo()
	fmt.Println(book)
}

func Test_GetZongHengFans(t *testing.T) {
	var url string
	url = "http://book.zongheng.com/book/490607.html"
	z := ZongHeng{BookInfoURL: url}
	fans, _ := z.GetFans()
	fmt.Println(fans)
}
