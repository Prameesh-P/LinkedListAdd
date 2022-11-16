package main

import (
	"fmt"
	"strings"
)

type node struct {
	data int
	next *node
}
type linkidList struct {
	head   *node
	length int
}

func (l *linkidList) add(val int) {
	tail := new(node)
	tail.data = val
	if l.head == nil {
		l.head = tail
		return
	}
	in := l.head
	for ; in.next != nil; in = in.next {
	}
	in.next = l.head
}
func (l linkidList) String() string {
	sb:=strings.Builder{}
	for in:=l.head;in!=nil;in=in.next{
		sb.WriteString(fmt.Sprintf("%d ",in.data))
	}
	return sb.String()
}

func main() {
	ll:=linkidList{}
	ll.add(44)
	ll.add(33)
	fmt.Println(ll)
}