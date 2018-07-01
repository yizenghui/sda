// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package wechat

import (
	"testing"

	"github.com/yizenghui/sda/code"
)

func Test_GetQiDianUpdate(t *testing.T) {

	// url := "https://mp.weixin.qq.com/s/tKQufLU2i7iJuM7v49jmRQ"
	url := "https://mp.weixin.qq.com/s/NiZ5iszTKEo2dxYo8mbRZg"

	a, _ := Find(url)
	t.Fatal(a.ReadContent)
}

func Test_GetGetVideo(t *testing.T) {

	// url := "https://mp.weixin.qq.com/s/tKQufLU2i7iJuM7v49jmRQ"
	url := "https://mp.weixin.qq.com/s/NiZ5iszTKEo2dxYo8mbRZg"

	a, _ := Find(url)
	t.Fatal(a.Video)
}
func Test_GetGetAudio(t *testing.T) {

	// url := "https://mp.weixin.qq.com/s/tKQufLU2i7iJuM7v49jmRQ"
	url := "https://mp.weixin.qq.com/s?__biz=MjM5MDMyMzg2MA==&mid=2655522346&idx=1&sn=ef39b32b4811ef20181e530e42cf73b1&chksm=bdfa2bd18a8da2c72236146f4cc9ab603b35d5f9f9375e78567e33c9a7126359d8f7703b1c90"

	a, _ := Find(url)
	t.Fatal(a.Audio)
}

//
func Test_GetList(t *testing.T) {

	content := `"Content": "&lt;msg&gt;&lt;appmsg appid=\"\" sdkver=\"0\"&gt;&lt;title&gt;&lt;
	![CDATA[为抗日牺牲的美女护士：用石头砸死日本军官，遗骸数十年才归故里]]&gt;&lt;/title&gt;&lt;des&gt;&lt;
	![CDATA[1938年3月，台儿庄战役爆发。中国方面守军29万人与日军5万人展开血战。战役历时一个月，中方伤亡5万余人，日方伤亡2]]&gt;&lt;
	/des&gt;&lt;action&gt;&lt;/action&gt;&lt;type&gt;5&lt;/type&gt;&lt;showtype&gt;1&lt;/showtype&gt;
	&lt;soundtype&gt;0&lt;/soundtype&gt;&lt;content&gt;&lt;![CDATA[]]&gt;&lt;/content&gt;&lt;contentattr&gt;0&lt;/contentattr&gt;&lt;url&gt;&lt;
	![CDATA[http://mp.weixin.qq.com/s?__biz=MzA3MjkxMTIyNw==&amp;mid=2650840174&amp;idx=1&amp;sn=93a278779739ade1ed56eff489ab3d69&amp;chksm=84e32c7ab394a56c7182ae2beca1824d3a2a6105c31750ee2121d8914d4bd40ae0055a8879a1#rd]]
	&gt;&lt;/url&gt;&lt;lowurl&gt;&lt;![CDATA[]]&gt;&lt;/lowurl&gt;&lt;appattach&gt;&lt;totallen&gt;0&lt;/totallen&gt;&lt;attachid&gt;&lt;/attachid&gt;&lt;fileext&gt;&lt;/fileext&gt;&lt;cdnthumburl&gt;&lt;![CDATA[]]&gt;&lt;/cdnthumburl&gt;&lt;cdnthumbaeskey&gt;&lt;![CDATA[]]&gt;&lt;/cdnthumbaeskey&gt;&lt;aeskey&gt;&lt;![CDATA[]]&gt;&lt;/aeskey&gt;&lt;/appattach&gt;&lt;extinfo&gt;&lt;/extinfo&gt;&lt;sourceusername&gt;&lt;![CDATA[]]&gt;&lt;/sourceusername&gt;&lt;sourcedisplayname&gt;&lt;![CDATA[]]&gt;&lt;/sourcedisplayname&gt;&lt;mmreader&gt;&lt;category type=\"0\" count=\"2\"&gt;&lt;name&gt;&lt;![CDATA[微信公众平台测试号]]&gt;&lt;/name&gt;&lt;topnew&gt;&lt;cover&gt;&lt;
	![CDATA[http://pic3.readfollow.com/aHR0cDovL21tYml6LnFwaWMuY24vbW1iaXpfanBnL2ZZTmcxaWJPUEJ0TkhSYm5pY2tscXFJSzRnbm4waFZreTk3YTFId2oyaDRxZ3pXYmljYmNieUtwTjZlbWlhdFJSa0tXN2pFWEg0NHhVdlpXamczaEF0T2pKUS8wP3d4X2ZtdD1qcGVn]]&gt;&lt;/cover&gt;&lt;width&gt;336&lt;/width&gt;&lt;height&gt;215&lt;/height&gt;&lt;digest&gt;&lt;![CDATA[1938年3月，台儿庄战役爆发。中国方面守军29万人与日军5万人展开血战。战役历时一个月，中方伤亡5万余人，日方伤亡2]]&gt;&lt;/digest&gt;&lt;/topnew&gt;&lt;item&gt;&lt;itemshowtype&gt;0&lt;/itemshowtype&gt;&lt;title&gt;&lt;![CDATA[为抗日牺牲的美女护士：用石头砸死日本军官，遗骸数十年才归故里]]&gt;&lt;/title&gt;&lt;url&gt;&lt;
	![CDATA[http://mp.weixin.qq.com/s?__biz=MzA3MjkxMTIyNw==&amp;mid=2650840174&amp;idx=1&amp;sn=93a278779739ade1ed56eff489ab3d69&amp;chksm=84e32c7ab394a56c7182ae2beca1824d3a2a6105c31750ee2121d8914d4bd40ae0055a8879a1#rd]]
	
	
	&gt;&lt;/url&gt;&lt;shorturl&gt;&lt;![CDATA[]]&gt;&lt;/shorturl&gt;&lt;longurl&gt;&lt;![CDATA[]]&gt;&lt;/longurl&gt;&lt;pub_time&gt;1507696678&lt;/pub_time&gt;&lt;cover&gt;&lt;
	![CDATA[http://pic3.readfollow.com/aHR0cDovL21tYml6LnFwaWMuY24vbW1iaXpfanBnL2ZZTmcxaWJPUEJ0TkhSYm5pY2tscXFJSzRnbm4waFZreTk3YTFId2oyaDRxZ3pXYmljYmNieUtwTjZlbWlhdFJSa0tXN2pFWEg0NHhVdlpXamczaEF0T2pKUS8wP3d4X2ZtdD1qcGVn]]&gt;&lt;/cover&gt;&lt;tweetid&gt;&lt;/tweetid&gt;&lt;digest&gt;&lt;![CDATA[1938年3月，台儿庄战役爆发。中国方面守军29万人与日军5万人展开血战。战役历时一个月，中方伤亡5万余人，日方伤亡2]]&gt;&lt;/digest&gt;&lt;fileid&gt;0&lt;/fileid&gt;&lt;sources&gt;&lt;source&gt;&lt;name&gt;&lt;![CDATA[微信公众平台测试号]]&gt;&lt;/name&gt;&lt;/source&gt;&lt;/sources&gt;&lt;styles&gt;&lt;/styles&gt;&lt;native_url&gt;&lt;/native_url&gt;&lt;del_flag&gt;0&lt;/del_flag&gt;&lt;contentattr&gt;0&lt;/contentattr&gt;&lt;play_length&gt;0&lt;/play_length&gt;&lt;play_url&gt;&lt;/play_url&gt;&lt;player&gt;&lt;![CDATA[]]&gt;&lt;/player&gt;&lt;template_op_type&gt;0&lt;/template_op_type&gt;&lt;weapp_username&gt;&lt;![CDATA[]]&gt;&lt;/weapp_username&gt;&lt;weapp_path&gt;&lt;![CDATA[]]&gt;&lt;/weapp_path&gt;&lt;weapp_version&gt;0&lt;/weapp_version&gt;&lt;weapp_state&gt;0&lt;/weapp_state&gt;&lt;/item&gt;&lt;item&gt;&lt;itemshowtype&gt;0&lt;/itemshowtype&gt;&lt;title&gt;&lt;![CDATA[画如美人—— 读 《状态 ：艺术品的老化 》]]&gt;&lt;/title&gt;&lt;url&gt;&lt;
	![CDATA[http://mp.weixin.qq.com/s?__biz=MzA5MTkyMzgwMA==&amp;mid=2818961058&amp;idx=3&amp;sn=587127aabbc6765b49c7d5967bf81393&amp;chksm=bd85ba958af233835aa00627521a8e3a1adf64cca6bf7a3ee49b8ba5aa1f2d2c92f9cbbaa83c#rd]]
	&gt;&lt;/url&gt;&lt;shorturl&gt;&lt;![CDATA[]]&gt;&lt;/shorturl&gt;&lt;longurl&gt;&lt;![CDATA[]]&gt;&lt;/longurl&gt;&lt;pub_time&gt;1507696678&lt;/pub_time&gt;&lt;cover&gt;&lt;
	![CDATA[http://pic3.readfollow.com/aHR0cDovL21tYml6LnFwaWMuY24vbW1iaXpfanBnL3BQM1AxTHlSYVJ5SWhpYmptTVkyNU9pYW5kYmFVUnJLcmlib28xR0RZWTlkVWRSVVRlQ1g2Y2liR3E1Y2o5aWNpYlV2dDVraWNFdlVJaExPb2xzZlpFZnJ5V1JsQS8wP3d4X2ZtdD1qcGVn]]&gt;&lt;/cover&gt;&lt;tweetid&gt;&lt;/tweetid&gt;&lt;digest&gt;&lt;![CDATA[假如你在卢浮宫隔着熙攘的人群望着列奥达多·达·芬奇的名画《蒙娜丽莎》出神，是否会想象或者担心这个恬静微笑着的女人慢慢变得白发苍苍、眼眶深陷、皱纹横生？]]&gt;&lt;/digest&gt;&lt;fileid&gt;0&lt;/fileid&gt;&lt;sources&gt;&lt;source&gt;&lt;name&gt;&lt;![CDATA[微信公众平台测试号]]&gt;&lt;/name&gt;&lt;/source&gt;&lt;/sources&gt;&lt;styles&gt;&lt;/styles&gt;&lt;native_url&gt;&lt;/native_url&gt;&lt;del_flag&gt;0&lt;/del_flag&gt;&lt;contentattr&gt;0&lt;/contentattr&gt;&lt;play_length&gt;0&lt;/play_length&gt;&lt;play_url&gt;&lt;/play_url&gt;&lt;player&gt;&lt;![CDATA[]]&gt;&lt;/player&gt;&lt;template_op_type&gt;0&lt;/template_op_type&gt;&lt;weapp_username&gt;&lt;![CDATA[]]&gt;&lt;/weapp_username&gt;&lt;weapp_path&gt;&lt;![CDATA[]]&gt;&lt;/weapp_path&gt;&lt;weapp_version&gt;0&lt;/weapp_version&gt;&lt;weapp_state&gt;0&lt;/weapp_state&gt;&lt;/item&gt;&lt;/category&gt;&lt;publisher&gt;&lt;username&gt;&lt;![CDATA[gh_cb5c31e2c2dd]]&gt;&lt;/username&gt;&lt;nickname&gt;&lt;![CDATA[微信公众平台测试号]]&gt;&lt;/nickname&gt;&lt;/publisher&gt;&lt;template_header&gt;&lt;/template_header&gt;&lt;template_detail&gt;&lt;/template_detail&gt;&lt;forbid_forward&gt;0&lt;/forbid_forward&gt;&lt;/mmreader&gt;&lt;thumburl&gt;&lt;
	![CDATA[http://pic3.readfollow.com/aHR0cDovL21tYml6LnFwaWMuY24vbW1iaXpfanBnL2ZZTmcxaWJPUEJ0TkhSYm5pY2tscXFJSzRnbm4waFZreTk3YTFId2oyaDRxZ3pXYmljYmNieUtwTjZlbWlhdFJSa0tXN2pFWEg0NHhVdlpXamczaEF0T2pKUS8wP3d4X2ZtdD1qcGVn]]&gt;&lt;/thumburl&gt;&lt;/appmsg&gt;&lt;fromusername&gt;&lt;![CDATA[gh_cb5c31e2c2dd]]&gt;&lt;/fromusername&gt;&lt;appinfo&gt;&lt;version&gt;0&lt;/version&gt;&lt;appname&gt;&lt;![CDATA[微信公众平台测试号]]&gt;&lt;/appname&gt;&lt;isforceupdate&gt;1&lt;/isforceupdate&gt;&lt;/appinfo&gt;&lt;/msg&gt;",
	`
	// url := "https://mp.weixin.qq.com/s/tKQufLU2i7iJuM7v49jmRQ"
	// url := "https://mp.weixin.qq.com/s/NiZ5iszTKEo2dxYo8mbRZg"

	// data := code.SelectString(`![CDATA[http://mp.weixin.qq.com/s(?P<uri>[^\]]+)]]`, content)

	data := code.SelectString(`http://mp.weixin.qq.com/s(?P<title>[a-zA-Z0-9\?_\-&=;]+)]]`, content)
	// a, _ := Find(url)
	panic(data)
}
