 package main

import (
	"fmt"
	
)

// mostra a lista com o elemento sword destacado
func ToStr(l *DList[int], sword *DNode[int]) string {
	resultado := "["

	for it := l.Front(); it != l.End(); it = it.next{
		if it == sword{
			resultado += fmt.Sprintf(" %d>", it.Value)
		}else{
			resultado += fmt.Sprintf(" %d", it.Value)
		}
	}

	resultado += " ]"
	return resultado
}
// move para frente na lista circular
func Next(l *DList[int], it *DNode[int]) *DNode[int] {
	prox := it.next

	if prox == l.End(){
		return l.Front()
	}
	return prox
}

func main() {
	var qtd, chosen int
	fmt.Scan(&qtd, &chosen)
	//fmt.Println(qtd, chosen)
	l := NewDList[int]()
	for i := 1; i <= qtd; i++ {
		l.PushBack(i)
	}
	sword := l.Front()
	for range chosen - 1 {
		sword = Next(l, sword)
	}
	for range qtd - 1 {
		fmt.Println(ToStr(l, sword))
		l.Erase(Next(l, sword))
		sword = Next(l, sword)
	}
	fmt.Println(ToStr(l, sword))
}
