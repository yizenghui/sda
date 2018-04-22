package wechat

import (
	"errors"
	"strings"

	"github.com/PuerkitoBio/goquery"
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
	Ban         bool
	Limit       bool
	Recommend   bool
}

// Find ..
func Find(url string) (article Article, err error) {

	g, e := goquery.NewDocument(url)
	if e != nil {
		return article, nil
	}

	html, _ := g.Html()

	ext, err := html2article.NewFromHtml(html)
	if err != nil {
		return article, err
	}
	art, err := ext.ToArticle()
	if err != nil {
		return article, err
	}
	// fmt.Println(article)

	//parse the article to be readability
	art.Readable(url)

	article.Content = art.Content
	article.ReadContent = art.ReadContent

	article.AppID = strings.TrimSpace(code.FindString(`var user_name = "(?P<user_name>[^"]+)";`, html, "user_name"))

	article.AppName = strings.TrimSpace(code.FindString(`var nickname = "(?P<nickname>[^"]+)";`, html, "nickname"))
	//
	article.Title = strings.TrimSpace(code.FindString(`var msg_title = "(?P<title>[^"]+)";`, html, "title"))

	//
	article.Intro = strings.TrimSpace(code.FindString(`var msg_desc = "(?P<intro>[^"]+)";`, html, "intro"))

	article.Cover = strings.TrimSpace(code.FindString(`var msg_cdn_url = "(?P<cover>[^"]+)";`, html, "cover"))

	article.RoundHead = strings.TrimSpace(code.FindString(`var round_head_img = "(?P<round_head>[^"]+)";`, html, "round_head"))

	article.OriHead = strings.TrimSpace(code.FindString(`var ori_head_img_url = "(?P<ori_head>[^"]+)";`, html, "ori_head"))
	//
	// article.PubAt = strings.TrimSpace(code.FindString(`var publish_time = "(?P<date>[^"]+)"`, html, "date"))

	article.PubAt = strings.TrimSpace(code.FindString(`var ct = "(?P<date>\d+)";`, html, "date"))

	link := strings.TrimSpace(code.FindString(`var msg_link = "(?P<url>[^"]+)";`, html, "url"))

	article.URL = strings.Replace(link, `\x26amp;`, "&", -1)

	link2 := strings.TrimSpace(code.FindString(`var msg_source_url = '(?P<url>[^']+)';`, html, "url"))

	article.SourceURL = strings.Replace(link2, `\x26amp;`, "&", -1)

	article.Author = strings.TrimSpace(code.FindString(`<em class="rich_media_meta rich_media_meta_text">(?P<author>[^<]+)</em>`, html, "author"))

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
	article.Title = strings.Replace(article.Title, `\x26amp;`, `&`, -1)
	article.Title = strings.Replace(article.Title, `\x26gt;`, `>`, -1)
	article.Title = strings.Replace(article.Title, `\x26lt;`, `<`, -1)
	article.Title = strings.Replace(article.Title, `\x0a`, "\n", -1)
	article.Title = strings.Replace(article.Title, `\x26#39;`, `'`, -1)

	article.Intro = strings.Replace(article.Intro, `\x0a`, "\n", -1)
	article.Intro = strings.Replace(article.Intro, `\x26quot;`, `"`, -1)
	article.Intro = strings.Replace(article.Intro, `\x26gt;`, `>`, -1)
	article.Intro = strings.Replace(article.Intro, `\x26lt;`, `<`, -1)
	article.Intro = strings.Replace(article.Intro, `\x26amp;`, `&`, -1)
	article.Intro = strings.Replace(article.Intro, `\x26#39;`, `'`, -1)

	return article, nil
}
