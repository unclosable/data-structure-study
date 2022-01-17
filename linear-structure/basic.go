package linear_structure

import "fmt"

type coefficient int

type term struct {
	index       int
	coefficient int
}

type linkedTerm struct {
	term
	next *linkedTerm
}

func mkLinkedTerm(terms []term) linkedTerm {
	var reTerm linkedTerm
	flg := &reTerm
	for i, t := range terms {
		flg.index = t.index
		flg.coefficient = t.coefficient
		if i < len(terms)-1 {
			flg.next = &linkedTerm{}
			flg = flg.next
		}
	}
	return reTerm
}

//2.1
func LinearStructureExpress() {
	// as array f(x)=4x**5-3x**2+1
	linearArray := []coefficient{1, 0, -3, 0, 0, 4}
	fmt.Println(linearArray)

	// as indexed array
	// P1 f(x)=9x**12+15x**8+3x**2
	// P2 f(x)=26x**19-4x**8-13x**6+82
	indexedArrayP1 := []term{
		{2, 3},
		{8, 15},
		{12, 9},
	}

	indexedArrayP2 := []term{
		{0, 82},
		{6, -13},
		{8, -4},
		{19, 26},
	}

	indexedArrayP3 := indexedArrayAddition(indexedArrayP1, indexedArrayP2)
	fmt.Println(indexedArrayP3)

	linkedTerm1 := mkLinkedTerm(indexedArrayP1)
	linkedTerm2 := mkLinkedTerm(indexedArrayP2)

	linkedTerm3 := linkedArrayAddition(&linkedTerm1, &linkedTerm2)

	for f := linkedTerm3; f != nil; f = f.next {
		fmt.Println(*f)
	}
}
func indexedArrayAddition(p1 []term, p2 []term) []term {
	if p1 == nil || len(p1) == 0 {
		return p2
	}
	if p2 == nil || len(p2) == 0 {
		return p1
	}
	p1flg, p2flg := 0, 0
	var re []term
	for p1flg < len(p1) || p2flg < len(p2) {
		if p1flg >= len(p1) {
			re = append(re, p2[p2flg])
			p2flg++
		} else if p2flg >= len(p2) {
			re = append(re, p1[p1flg])
			p1flg++
		} else if p1[p1flg].index < p2[p2flg].index {
			re = append(re, p1[p1flg])
			p1flg++
		} else if p1[p1flg].index > p2[p2flg].index {
			re = append(re, p2[p2flg])
			p2flg++
		} else {
			re = append(re, term{index: p1[p1flg].index, coefficient: p1[p1flg].coefficient + p2[p2flg].coefficient})
			p1flg++
			p2flg++
		}
	}
	return re
}
func linkedArrayAddition(p1 *linkedTerm, p2 *linkedTerm) *linkedTerm {
	if p1 == nil && p2 != nil {
		return p2
	} else if p2 == nil && p1 != nil {
		return p1
	} else if p1 == nil && p2 == nil {
		return nil
	}

	var re *linkedTerm
	if p1.index == p2.index {
		re = p1
		p1.coefficient += p2.coefficient
		p1.next = linkedArrayAddition(p1.next, p2.next)
	} else if p1.index > p2.index {
		re = p2
		re.next = linkedArrayAddition(p1, p2.next)
	} else if p1.index < p2.index {
		re = p1
		re.next = linkedArrayAddition(p1.next, p2)
	}
	return re
}
