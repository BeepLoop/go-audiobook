package types

type Words struct {
	Id   int    `db:"id"`
	Word string `db:"word"`
	Url  string `db:"url"`
}

func NewWordData(word, url string) *Words {
	return &Words{
		Word: word,
		Url:  url,
	}
}
