package main

import (
	"bufio"
	"fmt"
	"os"
)

// Não mude a assinatura desta função, ela é a função chamada pelo LeetCode
func exist(grid [][]byte, word string) bool {
	_, _ = grid, word
	return false
}
var dfs func(r, c, indice int) bool
dfs = func(r, c, indice int) bool{
	if indice == len(word){
		return true
	}
}

if r < 0 || c < 0 || r >= len(grid) || c >= len(grid[0]) || grid[c][r] != word[indice]{
	return false
}


func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	var word string
	fmt.Sscanf(scanner.Text(), "%s", &word)
	grid := make([][]byte, 0)
	for scanner.Scan() {
		grid = append(grid, []byte(scanner.Text()))
	}
	if exist(grid, word) {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}
