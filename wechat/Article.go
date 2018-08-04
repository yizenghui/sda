package wechat

import (
	"errors"
	"fmt"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/lunny/html2md"
	"github.com/russross/blackfriday"
	"github.com/sundy-li/html2article"
	"github.com/yizenghui/sda/code"
)

//Article struct
type Article struct {
	Title       string
	Author      string
	AppName     string
	AppID       string
	Cover       string
	Intro       string
	Content     string
	ReadContent string
	PubAt       string
	URL         string
	RoundHead   string
	OriHead     string
	SourceURL   string
	WxID        string
	WxIntro     string
	Copyright   string
	Ban         bool
	Limit       bool
	Recommend   bool
	Video       string
	Audio       string
}

// Find ..
func Find(url string) (article Article, err error) {

	g, e := goquery.NewDocument(url)
	if e != nil {
		return article, nil
	}

	html, _ := g.Html()

	contentHTML, err := g.Find("#js_content").Html()
	if err != nil {
		return article, err
	}
	html2 := fmt.Sprintf(`
		<html>
		<head>
		<meta http-equiv="Content-Type" content="text/html; charset=utf-8"/>
		<title>%v</title>
		<body><div class="rich_media_content " id="js_content">
		%v
		</div>
		</body>
		</html>
		`, `NONE TITLE`, contentHTML)
	// panic(html2)
	ext, err := html2article.NewFromHtml(html2)
	if err != nil {
		return article, err
	}
	art, err := ext.ToArticle()
	if err != nil {
		return article, err
	}
	// fmt.Println(article)

	// 视频地址
	article.Video, _ = g.Find("iframe").Eq(0).Attr("data-src")

	if article.Video == `` {
		article.Video, _ = g.Find("video").Eq(0).Attr("src")
	}

	// 音频地址
	audio, _ := g.Find("mpvoice").Eq(0).Attr("voice_encode_fileid")

	if audio != `` {
		article.Audio = fmt.Sprintf("https://res.wx.qq.com/voice/getvoice?mediaid=%v", audio)
	} else {
		article.Audio, _ = g.Find("audio").Eq(0).Attr("src")
	}

	//parse the article to be readability
	art.Readable(url)

	article.Content = art.Content
	article.ReadContent = art.ReadContent
	// article.ReadContent, _ = g.Find("#js_content").Html()

	article.AppID = strings.TrimSpace(code.FindString(`var user_name = "(?P<user_name>[^"]+)";`, html, "user_name"))

	article.AppName = strings.TrimSpace(code.FindString(`var nickname = "(?P<nickname>[^"]+)";`, html, "nickname"))
	//
	article.Title = strings.TrimSpace(code.FindString(`var msg_title = "(?P<title>[^"]+)";`, html, "title"))

	//
	article.Intro = strings.TrimSpace(code.FindString(`var msg_desc = "(?P<intro>[^"]+)";`, html, "intro"))

	article.WxID = strings.TrimSpace(code.FindString(`<label class="profile_meta_label">微信号</label>(?P<intro>[\s]+)<span class="profile_meta_value">(?P<wxid>[^"]+)</span>`, html, "wxid"))

	article.WxIntro = strings.TrimSpace(code.FindString(`<label class="profile_meta_label">功能介绍</label>(?P<intro>[\s]+)<span class="profile_meta_value">(?P<wxintro>[^"]+)</span>`, html, "wxintro"))

	article.Cover = strings.TrimSpace(code.FindString(`var msg_cdn_url = "(?P<cover>[^"]+)";`, html, "cover"))

	article.RoundHead = strings.TrimSpace(code.FindString(`var round_head_img = "(?P<round_head>[^"]+)";`, html, "round_head"))

	article.OriHead = strings.TrimSpace(code.FindString(`var ori_head_img_url = "(?P<ori_head>[^"]+)";`, html, "ori_head"))
	//
	// article.PubAt = strings.TrimSpace(code.FindString(`var publish_time = "(?P<date>[^"]+)"`, html, "date"))

	article.PubAt = strings.TrimSpace(code.FindString(`var ct = "(?P<date>\d+)";`, html, "date"))

	article.Copyright = strings.TrimSpace(code.FindString(`var _copyright_stat = "(?P<copyright>\d+)";`, html, "copyright"))
	// var _copyright_stat = "2";

	link := strings.TrimSpace(code.FindString(`var msg_link = "(?P<url>[^"]+)";`, html, "url"))

	article.URL = strings.Replace(link, `\x26amp;`, "&", -1)

	link2 := strings.TrimSpace(code.FindString(`var msg_source_url = '(?P<url>[^']+)';`, html, "url"))

	article.SourceURL = strings.Replace(link2, `\x26amp;`, "&", -1)

	article.Author = strings.TrimSpace(code.FindString(`<em class="rich_media_meta rich_media_meta_text">(?P<author>[^<]+)</em>`, html, "author"))

	// data-src="https://v.qq.com/iframe/preview.html?vid=p0689redfaq&amp;width=500&amp;height=375&amp;auto=0"></iframe>

	if strings.Contains(article.SourceURL, string("readfollow.com")) {
		article.Recommend = true
	}
	if strings.Contains(html, string("ban.readfollow.com")) {
		article.Ban = true
	}
	if strings.Contains(html, string("limit.readfollow.com")) {
		article.Limit = true
	}

	// fmt.Println(article)
	if article.AppName == "" {
		return article, errors.New("无法获取文章信息")
	}

	// 处理特殊字符
	article.URL = strings.Replace(article.URL, `http://`, `https://`, -1)
	article.URL = strings.Replace(article.URL, `#rd`, "&scene=27#wechat_redirect", 1)

	article.Title = strings.Replace(article.Title, `\x26quot;`, `"`, -1)
	article.Title = strings.Replace(article.Title, `\x0a`, "\n", -1)
	article.Title = strings.Replace(article.Title, `\x26gt;`, `>`, -1)
	article.Title = strings.Replace(article.Title, `\x26lt;`, `<`, -1)
	article.Title = strings.Replace(article.Title, `\x26amp;`, `&`, -1)
	article.Title = strings.Replace(article.Title, `\x26#39;`, `'`, -1)

	article.Intro = strings.Replace(article.Intro, `\x26quot;`, `"`, -1)
	article.Intro = strings.Replace(article.Intro, `\x0a`, "\n", -1)
	article.Intro = strings.Replace(article.Intro, `\x26gt;`, `>`, -1)
	article.Intro = strings.Replace(article.Intro, `\x26lt;`, `<`, -1)
	article.Intro = strings.Replace(article.Intro, `\x26amp;`, `&`, -1)
	article.Intro = strings.Replace(article.Intro, `\x26#39;`, `'`, -1)

	return article, nil
}

//MarkDownFormatContent 通过markdown语法格式化内容
func MarkDownFormatContent(content string) string {
	html2md.AddConvert(func(content string) string {
		// Pre code blocks
		re := regexp.MustCompile(`<span\b[^>]*>([\s\S]*)</span>`)
		content = re.ReplaceAllStringFunc(content, func(innerHTML string) string {
			matches := re.FindStringSubmatch(innerHTML)
			return matches[1]
		})
		return content
	})
	md := html2md.Convert(content)
	input := []byte(md)
	unsafe := blackfriday.MarkdownCommon(input)
	return string(unsafe[:])
	// contentBytes := bluemonday.UGCPolicy().SanitizeBytes(unsafe)
	// return strings.TrimSpace(fmt.Sprintf(`%v`, string(contentBytes[:])))
}
