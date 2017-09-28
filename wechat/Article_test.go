// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package wechat

import (
	"fmt"
	"testing"
)

func Test_GetQiDianUpdate(t *testing.T) {

	// url := "https://mp.weixin.qq.com/s/tKQufLU2i7iJuM7v49jmRQ"
	url := "https://mp.weixin.qq.com/s/NiZ5iszTKEo2dxYo8mbRZg"

	a, _ := Find(url)
	fmt.Println(a)
}
