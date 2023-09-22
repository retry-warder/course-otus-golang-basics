package binarysearch

type BSS struct {
	id   int
	data string
}

func (bss *BSS) SetID(id int) {
	bss.id = id
}

func (bss *BSS) SetData(data string) {
	bss.data = data
}

func NewBSS(id int, data string) *BSS {
	bss := new(BSS)
	bss.SetID(id)
	bss.SetData(data)
	return bss
}

type BSA []BSS

func BinarySearch(bsa BSA, id int) (bool, string) {
	var mid, exp int
	min := 0
	high := len(bsa) - 1
	for min <= high {
		mid = (min + high) / 2
		exp = mid
		findval := bsa[exp]
		if findval.id == id {
			return true, findval.data
		}
		if findval.id > id {
			high = mid - 1
		} else {
			min = mid + 1
		}
	}
	return false, ""
}
