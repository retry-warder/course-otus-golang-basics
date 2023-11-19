package types

import (
	"encoding/json"
)

type Book struct {
	ID     int     `json:"id"`
	Title  string  `json:"title"`
	Author string  `json:"author,omitempty"`
	Year   int     `json:"year"`
	Size   int     `json:"size"`
	Rate   float64 `json:"rate"`
}

func NewBook(id int, title string, author string, year int, size int, rate float64) Book {
	book := new(Book)
	book.ID = id
	book.Title = title
	book.Author = author
	book.Year = year
	book.Size = size
	book.Rate = rate
	return *book
}

func (p *Book) MarshalJSON() ([]byte, error) {
	return json.Marshal(*p)
}

func (p *Book) UnmarshalJSON(data []byte) error {
	type b Book
	tmp := &b{}
	err := json.Unmarshal(data, tmp)
	if err != nil {
		return err
	}
	*p = Book(*tmp)
	return nil
}
