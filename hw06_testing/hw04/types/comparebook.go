package hw04_types

type TypeCompare int

const (
	Year TypeCompare = iota + 1
	Size
	Rate
)

type CompareBook struct {
	typecompare TypeCompare
}

func (cb *CompareBook) GetTypeCompare() TypeCompare {
	return cb.typecompare
}

func NewCompareBook(typecompare TypeCompare) *CompareBook {
	cb := CompareBook{typecompare}
	return &cb
}

func (cb *CompareBook) Compare(book1 *Book, book2 *Book) bool {
	var compare bool
	switch cb.GetTypeCompare() {
	case Year:
		{
			compare = (book1.Year() > book2.Year())
		}
	case Size:
		{
			compare = book1.Size() > book2.Size()
		}
	case Rate:
		{
			compare = book1.Rate() > book2.Rate()
		}
	default:
		compare = false
	}
	return compare
}
