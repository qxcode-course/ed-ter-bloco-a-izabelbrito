package main
import (
	"bufio"
	"container/list"
	"fmt"
	"os"
)
func main() {
    scanner := bufio.NewScanner(os.Stdin)

    for scanner.Scan(){
		textoEntrada := scanner.Text()

        if textoEntrada == ""{
			continue
		}

    l := list.New()

    cursor := l.PushBack('|')

    for _, char := range textoEntrada{
        switch char{
        case 'R' :
            l.InsertBefore('\n', cursor)
        case '<':
            ladoEsquerdo := cursor.Prev()
            if ladoEsquerdo != nil{
                l.MoveBefore(cursor, ladoEsquerdo)
            }
        case '>':
				ladoDireito := cursor.Next()
				if ladoDireito != nil{
					l.MoveAfter(cursor, ladoDireito)
				}
        case 'B':
            ladoEsquerdo := cursor.Prev()
            if ladoEsquerdo != nil{
                l.Remove(ladoEsquerdo)
            }
        case 'D':
            ladoDireito := cursor.Next()
            if ladoDireito != nil{
                l.Remove(ladoDireito)
            }
        default:
            l.InsertBefore(char,cursor)
        }
    }

    for node:= l.Front(); node != nil; node = node.Next(){
        fmt.Printf("%c", node.Value.(rune))
    }

    fmt.Println()
    }
}