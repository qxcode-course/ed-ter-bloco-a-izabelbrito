package main
import "fmt"
func main() {

    var h, p, f, d int

    fmt.Scan(&h, &p, &f, &d)

    if d == 1{
        if f < h{
            if p > f && p < h{
                fmt.Println("N")
            }else{
                fmt.Println("S")
            }
        }else{
            if p > f || p < h{
                fmt.Println("N")
            }else{
                fmt.Println("S")
            }
        }
	} else { 
        if f > h {
			if p < f && p > h {
				fmt.Println("N")
			} else {
				fmt.Println("S")
			}
		} else {
			if p < f || p > h {
				fmt.Println("N")
			} else {
				fmt.Println("S")
			}
		}
    
    }
}