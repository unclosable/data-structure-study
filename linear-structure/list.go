package linear_structure

type Element interface {
	PrintValue()
}
type List interface {
	PrintList()
}
type ListRealization interface {
	MakeEmpty() List
	FindKth(k int, l List) Element
	Find(e Element, l List) int
	Insert(e Element, i int, l List)
	Delete(i int, l List)
	Length(l List) int
}
