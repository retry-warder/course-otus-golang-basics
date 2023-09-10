package types

type Book struct {
	id     int
	title  string
	author string
	year   int
	size   int
	rate   float64
}

func NewBook(id int, title string, author string, year int, size int, rate float64) *Book {
	book := Book{id, title, author, year, size, rate}
	return &book
}

func (b *Book) GetYear() int {
	return b.year
}

func (b *Book) GetSize() int {
	return b.size
}

func (b *Book) GetRate() float64 {
	return b.rate
}
