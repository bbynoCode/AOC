package main
 
import (
    "bufio"
    "fmt"
    "os"
    // "strings"
)
 
func main() {
	totalCal := 0
    const s = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	charMap := make(map[rune]int)
	for idx, char := range s {
		charMap[char] = idx + 1 // Add 1 to make it 1-52 instead of 0-51
	}

    filePath := "input.txt"
    readFile, _ := os.Open(filePath)
	defer readFile.Close()
  
    fileScanner := bufio.NewScanner(readFile)

    for fileScanner.Scan() {
        //fileLines = append(fileLines, fileScanner.Text())
		elfOne := createSetOfItems(fileScanner.Text())
		fileScanner.Scan()
		elfTwo := createSetOfItems(fileScanner.Text())
		fileScanner.Scan()
		elfThree := createSetOfItems(fileScanner.Text())

		for item := range elfOne {
			if elfTwo[item] && elfThree[item] {
				totalCal += charMap[item]
				break
			}
		}
    }
    fmt.Printf("The items add up to %d\n",totalCal)
}

func createSetOfItems(items string)(set map[rune]bool){
	set = make(map[rune]bool)
	for _, item := range items{
		set[item] = true
	}
	return
}