package main

import (
	"fmt"
	"strings"
)

func main() {
l:=linkidList{}
l.add(12)
l.add(2)
l.add(1)
fmt.Println(l)
}

type node struct {
	data int
	next *node
}
type linkidList struct {
	head *node
	len  int
}
func (l *linkidList) add (val int){
	tail:=new(node)
	tail.data=val
	if l.head==nil{
		l.head=tail
		return
	}
	in:=l.head
	for ; in.next!=nil; in=in.next{}
	in.next=tail
}
func (l linkidList)String ()string  {
	sb:= strings.Builder{}
	for in:=l.head;in!=nil;in=in.next{
		sb.WriteString(fmt.Sprintf("%d ",in.data))
	}
	return sb.String()
}