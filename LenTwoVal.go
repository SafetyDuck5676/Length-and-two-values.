package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Print("Please enter the array size:")
	nStr := Userinput()
	n, _ := strconv.Atoi(nStr)
	fmt.Print("Please enter the first value:")
	firstValue := Userinput()
	fmt.Print("Please enter the second value:")
	secondValue := Userinput()
	a := make([]string, n) // This makes a array with the length of n
	for i := 0; i < len(a); i++ {
		if i%2 != 0 { // If its an odd number assign it input 2
			a[i] = secondValue
		} else { // Else assign it input 1
			a[i] = firstValue
		}
	}
	fmt.Println(a)
}

func Userinput() string {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	return input
}
