package types

const (
	Year = "Year"
	Size = "Size"
	Rate = "Rate"
)

type CompareBook struct {
	typecompare string
	book1       Book
	book2       Book
}

func NewCompareBook(typecompare string, book1 Book, book2 Book) *CompareBook {
	cb := CompareBook{typecompare, book1, book2}
	return &cb
}

func (cb *CompareBook) Compare() bool {
	switch cb.typecompare {
	case Year:
		{
			return cb.book1.getYear() > cb.book2.getYear()
		}
	case Size:
		{
			return cb.book1.getSize() > cb.book2.getSize()
		}
	case Rate:
		{
			return cb.book1.getRate() > cb.book2.getRate()
		}
	default:
		return false
	}
}
