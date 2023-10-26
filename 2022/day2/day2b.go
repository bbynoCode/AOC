package main
 
import (
    "bufio"
    "fmt"
    "os"
    "strings"
)
 
func main() {
 
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
    
	var totalPoints int64 = 0
    for _, line := range fileLines {
		move := strings.Split(line, " ")

		if move[0] == "A"{ // Elf chooses Rock
			if move[1] == "X"{
				totalPoints += 3
			}else if move[1] == "Y"{
				totalPoints += 4
			}else if move[1] == "Z"{
				totalPoints += 8
			}
		} else if move[0] == "B"{ // Elf Chooses Paper
			if move[1] == "X"{
				totalPoints += 1
			}else if move[1] == "Y"{
				totalPoints += 5
			}else if move[1] == "Z"{
				totalPoints += 9
			}
		}else if move[0] == "C"{ // Elf Chooses Scissors
			if move[1] == "X"{
				totalPoints += 2
			}else if move[1] == "Y"{
				totalPoints += 6
			}else if move[1] == "Z"{
				totalPoints += 7
			}
		}
    }
	fmt.Println("Total Points: ",totalPoints)
}

