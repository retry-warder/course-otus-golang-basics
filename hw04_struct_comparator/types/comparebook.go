package types

import "fmt"

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
			compare = (book1.GetYear() > book2.GetYear())
		}
	case Size:
		{
			compare = book1.GetSize() > book2.GetSize()
		}
	case Rate:
		{
			compare = book1.GetRate() > book2.GetRate()
		}
	default:
		compare = false
	}
	if compare {
		fmt.Println(book1)
	} else {
		fmt.Println(book2)
	}
	return compare
}
