package data

// Book 采集书籍基本信息
type Book struct {
	Name       string
	Chapter    string
	Total      string
	Author     string
	Date       string
	BookURL    string
	ChapterURL string
	AuthorURL  string
	IsVIP      bool
}

// Fans 书籍粉丝
type Fans struct {
	Name  string
	URL   string
	Level int16 // 1个，2十，3千，4万
}
