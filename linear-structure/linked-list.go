package linear_structure

type LinkedElement struct {
	data interface{}
	next *LinkedElement
}

type LinkedList struct {
	head *LinkedElement
}

type LinkedListRealization struct {
	//
}

func (_ LinkedElement) MakeEmpty() LinkedList {
	return LinkedList{head: nil}
}

func (_ LinkedListRealization) FindKth(i int, list LinkedList) *LinkedElement {
	if list.head == nil {
		return nil
	}
	var reElement = list.head
	for i > 0 && reElement != nil {
		reElement = reElement.next
		i--
	}
	return reElement
}

func (_ LinkedListRealization) Find(element *LinkedElement, list LinkedList) int {
	reInt := -1
	flg := list.head
	found := false
	for flg != nil && !found {
		reInt += 1
		if flg == element {
			found = true
		} else {
			flg = flg.next
		}
	}
	if found {
		return reInt
	} else {
		return -1
	}
}

func (_ LinkedListRealization) Insert(element *LinkedElement, i int, list LinkedList) {
	h := list.head
	for h.next != nil {
		h = h.next
	}
	h.next = element
}
