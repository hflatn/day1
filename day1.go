package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)



func FindNumbers(message []string) int {
	// var sb strings.Builder
	// sb.WriteString("hello")
	var all_numbers int
	for _, word := range message {
		var foundNumber strings.Builder
		for _, letter := range word{
			if unicode.IsDigit(letter) {
				foundNumber.WriteRune(letter)
			}
		}
		// Strings are read only slices of bytes so thats why you have to convert them back into strings
		// When you say numberStr[0] that retuns bytes that you can conver to a String 
		numberStr := foundNumber.String()
		firstChar := string(numberStr[0])
		lastChar := string(numberStr[len(numberStr)-1])
		firstAndLast := firstChar + lastChar
		to_add, err := strconv.Atoi(firstAndLast)


		if err != nil {
			fmt.Println("Error converting string to integer:", err)
			continue
		}
		all_numbers += to_add
	}
	
	return all_numbers
}

func main() {
	gfsdgds()
	message := []string{"1abc2", "pqr3stu8vwx", "a1b2c3d4e5f", "treb7uchet"}
	result := FindNumbers(message)
	fmt.Println("Sum of all numbers in message: ", result)

}