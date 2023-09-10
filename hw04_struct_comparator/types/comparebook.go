package types

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
	switch cb.GetTypeCompare() {
	case Year:
		{
			return book1.GetYear() > book2.GetYear()
		}
	case Size:
		{
			return book1.GetSize() > book2.GetSize()
		}
	case Rate:
		{
			return book1.GetRate() > book2.GetRate()
		}
	default:
		return false
	}
}
