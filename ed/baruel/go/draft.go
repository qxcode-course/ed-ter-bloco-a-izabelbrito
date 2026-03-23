package main

import "fmt"

func main(){

    var n, qtd int

    fmt.Scan(&n, &qtd)

    album := make([]bool, n+1)

    rep := false

    for i := 0; i < qtd; i++{
        var fig int

        fmt.Scan(&fig)

        if album[fig] == true{
            if rep == true{
                fmt.Print(" ")
            }
            fmt.Print(fig)

            rep = true
        }else{
            album[fig] = true
        }

    }

	if rep == false{
		fmt.Print("N")
	}
	fmt.Println()

	falta := false

	for i := 1; i <= n; i++{
		if album[i] == false {
			if falta == true{
				fmt.Print(" ")
			}
			fmt.Print(i)
			falta = true
		}
    }

	if falta == false{
		fmt.Print("N")
	}
	fmt.Println()
}