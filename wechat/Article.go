package wechat

import (
	"strings"

	"github.com/GanEasy/sdax/code"
	"github.com/PuerkitoBio/goquery"
)

//Article struct
type Article struct {
	Title   string
	Author  string
	AppName string
	AppID   string
	Cover   string
	Intro   string
	PubAt   string
	URL     string
}

// Find ..
func Find(url string) (article Article, err error) {

	g, e := goquery.NewDocument(url)
	if e != nil {
		return article, nil
	}

	html, _ := g.Html()

	article.AppID = strings.TrimSpace(code.FindString(`var user_name = "(?P<user_name>[^"]+)";`, html, "user_name"))

	article.AppName = strings.TrimSpace(code.FindString(`var nickname = "(?P<nickname>[^"]+)";`, html, "nickname"))
	//
	article.Title = strings.TrimSpace(code.FindString(`var msg_title = "(?P<title>[^"]+)";`, html, "title"))

	//
	article.Intro = strings.TrimSpace(code.FindString(`var msg_desc = "(?P<intro>[^"]+)";`, html, "intro"))

	article.Cover = strings.TrimSpace(code.FindString(`var msg_cdn_url = "(?P<cover>[^"]+)";`, html, "cover"))

	//
	article.PubAt = strings.TrimSpace(code.FindString(`var publish_time = "(?P<date>[^"]+)"`, html, "date"))

	link := strings.TrimSpace(code.FindString(`var msg_link = "(?P<url>[^"]+)";`, html, "url"))

	article.URL = strings.Replace(link, `\x26amp;`, "&", -1)

	article.Author = strings.TrimSpace(code.FindString(`<em class="rich_media_meta rich_media_meta_text">(?P<author>[^<]+)</em>`, html, "author"))

	return article, nil
}
