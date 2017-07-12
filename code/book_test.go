package code

import (
	"fmt"
	"testing"
)

func Test_ExplainDetailedAddress(t *testing.T) {

	urls := []string{
		"http://vipreader.qidian.com/chapter/1004608738/377630007",
		"http://book.qidian.com/info/1004608738?from=readfollow",
		"http://m.qidian.com/book/1004608738/?from=readfollow",
		"http://m.qidian.com/book/1004608738/catalog?from=readfollow",
		"http://m.qidian.com/book/1004608738/342363924?from=readfollow",
		"http://m.zongheng.com/h5/chapter/list?bookid=490607",
		"http://m.zongheng.com/h5/book?bookid=490607",
		"http://book.zongheng.com/book/490607.html?from=readfollow",
		"http://m.zongheng.com/h5/chapter?bookid=490607&cid=8134632",
		"http://www.17k.com/chapter/2317974/27502630.html",
		"http://h5.17k.com/chapter/2317974/27502630.html",
		"http://www.17k.com/list/2317974.html",
		"http://h5.17k.com/list/2317974.html",
		"http://www.17k.com/book/2317974.html",
		"http://h5.17k.com/book/2317974.html",
	}
	for k, v := range urls {
		url := ExplainDetailedAddress(v)
		fmt.Println(k, v, url)
	}

}
