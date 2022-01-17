package linear_structure

import "fmt"

type ArrayElement struct {
	value interface{}
}

func (e ArrayElement) PrintValue() {
	fmt.Println(e.value)
}

type ArrayList struct {
	datas []*ArrayElement
	index int
}

func (a ArrayList) PrintList() {
	fmt.Println(a.datas)
}

type ArrayListRealization struct {
	// golang's interface is sucks
}

func (a ArrayListRealization) MakeEmpty() List {
	return ArrayList{index: -1, datas: make([]*ArrayElement, 10)}
}

func (a ArrayListRealization) FindKth(k int, l ArrayList) *ArrayElement {
	if k > a.Length(l) {
		return nil
	}
	return l.datas[k]
}

func (a ArrayListRealization) Find(e *ArrayElement, l ArrayList) int {
	for i := 0; i < l.index; i++ {
		// well ~
		if e == l.datas[i] {
			return i
		}
	}
	return -1
}

func (a ArrayListRealization) Insert(e *ArrayElement, i int, l ArrayList) {
	if l.index == len(l.datas)-1 {
		//full
		return
	}
	if i < 1 || i > l.index+2 {
		//index illegal
		return
	}
	for j := l.index; j >= i-1; j-- {
		l.datas[j+1] = l.datas[j]
	}
	l.datas[i-1] = e
	l.index++
}

func (a ArrayListRealization) Delete(i int, l ArrayList) {
	if i < 1 || i > l.index+1 {
		// i illegal
		return
	}
	for j := i; j <= l.index; j++ {
		l.datas[j-1] = l.datas[j]
	}
	l.index--
}

func (a ArrayListRealization) Length(l ArrayList) int {
	return l.index + 1
}
