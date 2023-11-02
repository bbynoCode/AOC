package main

import (
	"fmt"
	"os"
	"bufio"
)


func main(){

	input, _ := os.Open("input.txt")
	defer input.Close()
	scanner := bufio.NewScanner(input)

	scanner.Scan()
	s := scanner.Text()

	const diffChars = 4

	for i := range s {
		chars := make(map[byte]bool)
		for j:=0; j<diffChars; j++{
			chars[s[i+j]]=true
		}
		if len(chars) == diffChars{
			fmt.Println(i+diffChars)
			break
		}
	}
}

