package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func numbers(s string) []int {
    var n []int
    for _, f := range strings.Fields(s) {
        i, err := strconv.Atoi(f)
        if err == nil {
            n = append(n, i)
        }
    }
    return n
}

func GetInputSlice() []int {

    scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()
    return numbers(scanner.Text())

}

func main() {
	
	var input []int
    input = GetInputSlice()

	total :=input[0] + input[1]

	fmt.Println(total)

}