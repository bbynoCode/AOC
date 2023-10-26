package main
 
import (
    "bufio"
    "fmt"
    "os"
    "strconv"
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
    
    var total int64 = 0
    var highest int64 = 0
    var highest2 int64 = 0
    var highest3 int64 = 0
    var temp int64 = 0

    for _, line := range fileLines {

        if line == ""{
            //fmt.Println("Total: ",total)
            //fmt.Println("\n") 
            if total > highest {
                highest3 = highest2
                highest2 = highest
                highest = total
            } else if total > highest2 {
                highest3 = highest2
                highest2 = total
            } else if total > highest3 {
                highest3 = total
            }
            total = 0
        } else {
            temp, err = strconv.ParseInt(line, 10, 64)
            total += temp
            //fmt.Println(line)
        }  
    }
    fmt.Println("Total calories of top 3: ",highest + highest2 + highest3)
}