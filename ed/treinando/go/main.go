package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func tostr(vet []int) string {
	if len(vet) == 0 {
		return "[]"
	}
	var rec func(v []int) string
	rec = func(v []int) string {
		if len(v) == 1 {
			return strconv.Itoa(v[0])
		}
		return strconv.Itoa(v[0]) + ", " + rec(v[1:])
	}
	return "[" + rec(vet) + "]"
}

func tostrrev(vet []int) string {
	if len(vet) == 0 {
		return "[]"
	}
	var rec func(v []int) string
	rec = func(v []int) string {
		if len(v) == 1 {
			return strconv.Itoa(v[0])
		}
		return rec(v[1:]) + ", " + strconv.Itoa(v[0])
	}
	return "[" + rec(vet) + "]"
}

func reverse(vet []int) {
	if len(vet) <= 1 {
		return
	}
	vet[0], vet[len(vet) - 1] = vet[len(vet) - 1], vet[0]

	reverse(vet[1 : len(vet)-1])
}

func sum(vet []int) int {
	if len(vet) == 0 {
		return 0
	}
	return vet[0] + sum(vet[1:])
}

func mult(vet []int) int {
	if len(vet) == 0 {
		return 1
	}
	return vet[0] * mult(vet[1:])
}

func min(vet []int) int {
	if len(vet) == 0 {
		return -1
	}
	var rec func(v []int) (int, int)
	rec = func(v []int) (int, int) {
		if len(v) == 1 {
			return v[0], 0
		}
		valR, indR := rec(v[1:])
		if v[0] <= valR {
			return v[0], 0
		}
		return valR, indR+1
	}
	_, indice := rec(vet)
	return indice
}

func main() {
	var vet []int
	scanner := bufio.NewScanner(os.Stdin)
	for {
		if !scanner.Scan() {
			break
		}
		line := scanner.Text()
		args := strings.Fields(line)
		fmt.Println("$" + line)

		switch args[0] {
		case "end":
			return
		case "read":
			vet = nil
			for _, arg := range args[1:] {
				if val, err := strconv.Atoi(arg); err == nil {
					vet = append(vet, val)
				}
			}
		case "tostr":
			fmt.Println(tostr(vet))
		case "torev":
			fmt.Println(tostrrev(vet))
		case "reverse":
			reverse(vet)
		case "sum":
			fmt.Println(sum(vet))
		case "mult":
			fmt.Println(mult(vet))
		case "min":
			fmt.Println(min(vet))
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
