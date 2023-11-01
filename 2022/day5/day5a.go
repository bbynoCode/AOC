package main

import (
	"bufio"
	"fmt"
	"os"
)

type stack struct {
	elements []rune
}

func (stackBox *stack) push(box rune){
	stackBox.elements = append(stackBox.elements, box)
}

func (stackBox *stack) pop()(box rune){
	box = stackBox.elements[len(stackBox.elements)-1]
	stackBox.elements = stackBox.elements[:len(stackBox.elements)-1]
	return
}

func (stackBox *stack) addToBottom(box rune){
	stackBox.elements = append([]rune{box}, stackBox.elements...)
}

func (stackBox stack) String()string{
	var str string
	for _, box := range stackBox.elements{
		str += string(box)+" "
	}
	return str
}

func main(){
	
	input,_ := os.Open("input.txt")
	defer input.Close()
	scanner := bufio.NewScanner(input)

	stacks := make([]stack, 9)

	//Parsing the input
	scanner.Scan()
	for (scanner.Text()!=" 1   2   3   4   5   6   7   8   9 "){
		for i, box := range scanner.Text(){
			if box != ' ' && box!='[' && box!=']'{
				stacks[i/4].addToBottom(box)
			}
		}
		scanner.Scan()
	}
	scanner.Scan()

	for scanner.Scan(){
		var toMove, from, to int
		fmt.Sscanf(scanner.Text(), "move %d from %d to %d", &toMove, &from, &to)
		for move :=0; move < toMove; move++{
			stacks[to-1].push(stacks[from-1].pop())
		}
	}

	for _, stackBox := range stacks{
		fmt.Print(string(stackBox.pop()))
	}
	fmt.Printf("\n")
}