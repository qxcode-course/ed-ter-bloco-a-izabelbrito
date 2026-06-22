package main

import (
	"fmt"
)
func main() {

	fila:= NewQueue[rune]()

	for i := 0; i < 16; i++ {
		fila.Enqueue(rune('A' + i))	
	}

	for i := 0; i < 15; i++ {
		var gols_time1, gols_time2 int

		if _, err := fmt.Scan(&gols_time1, &gols_time2); err != nil{
			break
		}

		time1 := fila.Dequeue()
		time2 := fila.Dequeue()

		if gols_time1 > gols_time2{
			fila.Enqueue(time1)
		}else{
			fila.Enqueue(time2)
		}	
	}

	campeao := fila.Dequeue()
	fmt.Printf("%c\n", campeao)
}
