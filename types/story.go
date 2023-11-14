package types

type Story struct {
	Id      string  `json:"id"`
	Title   string  `json:"title" form:"title" binding:"required"`
	Author  string  `json:"author" form:"author" binding:"required"`
	Content string  `json:"content" form:"content" binding:"required"`
	Audio   string  `json:"audio" form:"audio" binding:"required"`
	Image   string  `json:"image" form:"thumbnail" binding:"required"`
	Words   []*Word `json:"words" form:"words" binding:"required"`
}

type DataStore struct {
	Stories []Story `json:"stories"`
}

func NewWordData(word, url string) *Word {
	return &Word{
		Word:     word,
		Location: url,
	}
}
