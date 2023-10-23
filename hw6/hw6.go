package hw6

type ListInter interface {
	Len() int
	Front() *Node
	Back() *Node
	PushFront(v int) *Node
	PushBack(v int) *Node
	Remove(i int)
	MoveToFront(i int)
}

type Node struct {
	Value int
	Next  *Node
	Prev  *Node
}
type List struct {
	Len  int
	Head *Node
	Last *Node
}

func NewList() *List {
	return new(List)
}

func (l *List) PushBack(val int) *Node {
	l.Len += 1
	if l.Head == nil {
		l.Head = &Node{
			Value: val,
		}
		l.Last = &Node{
			Value: val,
		}
		return l.Head
	}
	curr := l.Head
	for curr.Next != nil {
		curr = curr.Next
	}
	elem := &Node{
		Value: val,
		Prev:  curr,
	}
	curr.Next = elem
	l.Last = elem
	return l.Last
}

func (l *List) PushFront(val int) *Node {
	l.Len += 1
	if l.Head == nil {
		l.Head = &Node{
			Value: val,
		}
		return l.Head
	}
	curr := l.Head
	elem := &Node{
		Value: val,
		Next:  curr,
	}
	curr.Prev = elem
	l.Head = elem
	return l.Head
}

func (l *List) Print() {
	i := 0
	curr := l.Head
	for {
		/*fmt.Println(i)
		fmt.Printf("curr: %+v\n", curr)*/
		if curr.Next == nil {
			break
		}
		curr = curr.Next
		i += 1
	}
}
func (l *List) LenList() int {
	//fmt.Printf("len is %d\n", l.Len)
	return l.Len

}
func (l *List) Front() *Node {
	//	fmt.Printf("first is %v\n", l.Head.Value)
	if l.Len == 0 {
		return nil
	}
	return l.Head
}
func (l *List) Back() *Node {
	//	fmt.Printf("last is %v\n", l.Last.Value)
	if l.Len == 0 {
		return nil
	}
	return l.Last
}
func (l *List) Remove(i int) {
	c := 1
	if i != l.Len {
		CurrR := l.Head
		for c < i {
			CurrR = CurrR.Next
			c++
		}
		CurrR.Prev.Next = CurrR.Next
		CurrR.Next.Prev = CurrR.Prev
	}
	if i == l.Len {
		l.Last.Prev.Next = nil
		l.Last = l.Last.Prev
		l.Last.Next = nil
	}
	l.Len--
}

func (l *List) MoveToFront(i int) {
	c := 1
	if i != 1 && i != l.Len {
		fsecond := l.Head
		curr := l.Head
		for c < i {
			curr = curr.Next
			c++
		}
		currtofront := curr

		curr.Prev.Next = curr.Next
		curr.Next.Prev = curr.Prev

		currtofront.Next = fsecond
		currtofront.Prev = nil
		fsecond.Prev = currtofront
		l.Head = currtofront
	}
	if i == l.Len {
		currtofront := l.Last
		fsecond := l.Head
		l.Last = l.Last.Prev
		l.Last.Next = nil
		currtofront.Next = fsecond
		currtofront.Prev = nil
		fsecond.Prev = currtofront
		l.Head = currtofront
	}
}
