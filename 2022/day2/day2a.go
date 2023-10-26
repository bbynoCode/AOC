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
    
	var totalPoints int32 = 0
    for _, line := range fileLines {
		move := strings.Split(line, " ")
		// fmt.Println("ELF MOVE: ",moves[0])
		// fmt.Println("My MOVE: " + moves[1])
		// fmt.Println()
		if move[0] == "A"{
			if move[1] == "X"{
				totalPoints += 4
			}else if move[1] == "Y"{
				totalPoints += 8
			}else if move[1] == "Z"{
				totalPoints += 3
			}
		} else if move[0] == "B"{
			if move[1] == "X"{
				totalPoints += 1
			}else if move[1] == "Y"{
				totalPoints += 5
			}else if move[1] == "Z"{
				totalPoints += 9
			}
		}else if move[0] == "C"{
			if move[1] == "X"{
				totalPoints += 7
			}else if move[1] == "Y"{
				totalPoints += 2
			}else if move[1] == "Z"{
				totalPoints += 6
			}
		}
    }
	fmt.Println("Total Points: ",totalPoints)
}

