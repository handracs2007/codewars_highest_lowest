package main

import (
	"fmt"
	"strconv"
	"strings"
)

func HighAndLow(in string) string {
	var highest, lowest int
	var numbers = strings.Split(in, " ")

	for i := 0; i < len(numbers); i++ {
		var number, _ = strconv.Atoi(numbers[i])

		if i == 0 {
			highest = number
			lowest = number
		} else {
			if number > highest {
				highest = number
			} else if number < lowest {
				lowest = number
			}
		}
	}

	return fmt.Sprintf("%v %v", highest, lowest)
}

func main() {
	fmt.Println(HighAndLow("1 2 3 4 5"))
	fmt.Println(HighAndLow("1 2 -3 4 5"))
	fmt.Println(HighAndLow("1 9 3 4 -5"))
}
