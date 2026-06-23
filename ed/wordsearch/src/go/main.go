package main

import (
	"bufio"
	"fmt"
	"os"
)

// Não mude a assinatura desta função, ela é a função chamada pelo LeetCode
func exist(grid [][]byte, word string) bool {
	
	for i := range grid{
		for j := range grid[i] {
			if dfs(grid, word, i, j, 0){
				return true
			}	
		}	
	}
	return false
}

func dfs(grid [][]byte, word string, l, c, index int) bool{
	if index == len(word) - 1{
		return l >= 0 && l < len(grid) && 
		c >= 0 && c < len(grid[l]) &&
		grid[l][c] == word[index]
	}
	
	if l < 0 || l >= len(grid) || c < 0 ||
	c >= len(grid[l]) || grid[l][c] != word[index]{
		return false
	}

	temp := grid[l][c]
	grid[l][c] = '#'

	if dfs(grid, word, l + 1, c, index + 1) ||
	dfs(grid, word, l, c + 1, index + 1) ||
	dfs(grid, word, l - 1, c, index + 1) ||
	dfs(grid, word, l, c - 1, index + 1){
		return true
	}

	grid[l][c] = temp
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
