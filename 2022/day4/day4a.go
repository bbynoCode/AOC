package main
 
import (
    "bufio"
    "fmt"
    "os"
    "strings"
	"strconv"
)
 
func main() {

	var overlap int
    filePath := "input.txt"
    readFile, _ := os.Open(filePath)
	defer readFile.Close()
  
    fileScanner := bufio.NewScanner(readFile)

    for fileScanner.Scan(){
		elves := strings.Split(fileScanner.Text(), ",")
		// fmt.Println(elves)
		elfOne := strings.Split(elves[0], "-")
		elfTwo := strings.Split(elves[1], "-")

		elfOneLow, _ := strconv.Atoi(elfOne[0])
		elfOneHigh, _ := strconv.Atoi(elfOne[1])
		elfTwoLow, _ := strconv.Atoi(elfTwo[0])
		elfTwoHigh, _ := strconv.Atoi(elfTwo[1])
		//elfOne is inside of elfTwo
		if elfOneLow >= elfTwoLow && elfOneHigh <= elfTwoHigh {
			overlap += 1

		//elfTwo is inside of elfOne
		} else if elfOneLow <= elfTwoLow && elfOneHigh >= elfTwoHigh {
			overlap += 1
		}
	}
	fmt.Printf("Number of elves that overlap: %d\n",overlap)
}






