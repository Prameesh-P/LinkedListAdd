package main

import (
	"fmt"
	"strings"
)

func main() {
	l := LinkidList{}
	l.add(1)
	l.add(2)
	l.add(3)
	l.add(4)
	l.add(5)
	l.add(6)
	l.add(7)
	fmt.Println(l)
}

type node struct {
	data int
	next *node
}
type LinkidList struct {
	head   *node
	Length int
}

func (l *LinkidList) add(value int) {
	tail := new(node)
	tail.data = value
	if l.head == nil {
		l.head = tail

	} else {
		in := l.head
		for ; in.next != nil; in = in.next {
		}
		in.next = tail
	}
}
func (l LinkidList) String() string {
	sb := strings.Builder{}
	for in := l.head; in != nil; in = in.next {
		sb.WriteString(fmt.Sprintf("%d", in.data))
	}

	return sb.String()
}
