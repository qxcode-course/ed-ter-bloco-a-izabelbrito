package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Node struct{
	Value int
	next *Node
	prev *Node
}

type LList struct{
	root *Node
	size int
}

func NewLList() *LList{
	sentinel := &Node{}
	sentinel.next = sentinel
	sentinel.prev = sentinel
	return &LList{
		root: sentinel,
		size: 0,
	}
}

func (l *LList) Size() int{
	return l.size
}

func (l *LList) String() string{
	if l.root.next == l.root{
		return "[]"
	}
	var sb strings.Builder
	sb.WriteString("[")

	node := l.root.next
	sb.WriteString(strconv.Itoa(node.Value))
	node = node.next

	for node != l.root{
		sb.WriteString(", ")
		sb.WriteString(strconv.Itoa(node.Value))
		node = node.next
	}

	sb.WriteString("]")
	return sb.String()
}

func (l *LList) PushFront(value int){
	newNode := &Node{Value: value}

	primeiro := l.root.next

	newNode.prev = l.root
	newNode.next = primeiro

	l.root.next = newNode
	primeiro.prev = newNode

	l.size ++
}

func (l *LList) PushBack(value int){
	newNode := &Node{Value: value}

	ultimo := l.root.prev

	newNode.prev = ultimo
	newNode.next = l.root

	ultimo.next = newNode
	l.root.prev = newNode

	l.size ++
} 

func (l *LList) PopFront(){
	if l.size == 0 {
		return
	}

	primeiro := l.root.next
	segundo := primeiro.next

	l.root.next = segundo
	segundo.prev = l.root

	l.size --
}

func (l *LList) PopBack(){
	if l.size == 0 {
		return
	}
	
	ultimo := l.root.prev
	penultimo := ultimo.prev

	l.root.prev = penultimo
	penultimo.next = l.root

	l.size --
}

func (l *LList) Clear(){
	l.root.next = l.root
	l.root.prev = l.root
	l.size --
}


func main() {
	scanner := bufio.NewScanner(os.Stdin)
	ll := NewLList()

	for {
		fmt.Print("$")
		if !scanner.Scan() {
			break
		}
		line := scanner.Text()
		fmt.Println(line)
		args := strings.Fields(line)

		if len(args) == 0 {
			continue
		}

		cmd := args[0]

		switch cmd {
		case "show":
			fmt.Println(ll.String())
		case "size":
			fmt.Println(ll.Size())
		case "push_back":
				for _, v := range args[1:] {
					num, _ := strconv.Atoi(v)
				 	ll.PushBack(num)
			}
		case "push_front":
				for _, v := range args[1:] {
					num, _ := strconv.Atoi(v)
					ll.PushFront(num)
			}
		case "pop_back":
			ll.PopBack()
		case "pop_front":
			ll.PopFront()
		case "clear":
			ll.Clear()
		case "end":
			return
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
