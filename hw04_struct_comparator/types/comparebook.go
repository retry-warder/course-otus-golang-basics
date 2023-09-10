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
			return cb.book1.GetYear() > cb.book2.GetYear()
		}
	case Size:
		{
			return cb.book1.GetSize() > cb.book2.GetSize()
		}
	case Rate:
		{
			return cb.book1.GetRate() > cb.book2.GetRate()
		}
	default:
		return false
	}
}
