package types

type Story struct {
	Id           int    `db:"id" `
	Title        string `json:"title" form:"title" db:"title" binding:"required"`
	Author       string `json:"author" form:"author" db:"author" binding:"required"`
	Story        string `json:"story" form:"story" db:"content" binding:"required"`
	Words        string `form:"words" binding:"required"`
	AudioUrl     string `form:"audio" db:"audio" binding:"required"`
	ThumbnailUrl string `form:"thumbnail" db:"thumbnail" binding:"required"`
}
