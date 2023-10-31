package main
 
import (
    "bufio"
    "fmt"
    "os"
    "strings"
)
 
func main() {
 
    const s = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	charMap := make(map[rune]int)
    
	for idx, char := range s {
		charMap[char] = idx + 1 // Add 1 to make it 1-52 instead of 0-51
	}

    filePath := "input.txt"
    readFile, err := os.Open(filePath)
  
    if err != nil {
        fmt.Println(err)
    }

    fileScanner := bufio.NewScanner(readFile)
    fileScanner.Split(bufio.ScanLines)
    var fileLines []string
  
    for fileScanner.Scan() {
        fileLines = append(fileLines, fileScanner.Text())
    }
    readFile.Close()
    
    totalCal := 0

    for _, line := range fileLines {
        for _, char := range line[:len(line)/2] {
            if strings.ContainsRune(line[(len(line)/2):], char) {
                totalCal += charMap[char]
                // Print debug statements
                //fmt.Printf("%s and %s both have %c\n", line[:len(line)/2], line[(len(line)/2):], char )
                //fmt.Printf("Adding %d. New Total: %d\n", charMap[char], totalCal)
                break
            }
        }
    }
    fmt.Printf("The items in both pockets add up to %d\n",totalCal)
}
