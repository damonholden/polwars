package main

import (
	"fmt"
	"math"
	"strconv"
)

func isNumberCharCodeRange(number int) bool {
	return number > 47 && number < 58

}

func isInLetterCharCodeRange(number int) bool {
	return number > 64 && number < 91
}

func convertCharCodeToInt(number int) int {
	if isNumberCharCodeRange(number) {
		return number - 48
	}

	if isInLetterCharCodeRange(number) {
		return number - 55
	}

	return 100
}

func convertDenaryNumberToRequiredBase(number int, base string) string {
	var output = ""
	for number > 3 {
		output += strconv.Itoa(number % 3)
		number = number / 3
	}

	output += strconv.Itoa(number % 3)

	return output
}

func convertNumber(number string, inputBase int, outputBase int) int {
	var output int

	if inputBase == outputBase {
		return output
	}

	// reverse string fas second for loop goes from left to right
	runes := []rune(number)
	for i, j := 0, len(number)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	number = string(runes)

	// convert to base 10
	for i := 0; i < len(number); i++ {
		output += convertCharCodeToInt(int(number[i])) * int(math.Pow(float64(inputBase), float64(i)))
	}

	if outputBase == 10 {
		return output
	}

	// convert to required base
	outputInCorrectBase := convertDenaryNumberToRequiredBase(output, "3")

	runess := []rune(outputInCorrectBase)
	for i, j := 0, len(outputInCorrectBase)-1; i < j; i, j = i+1, j-1 {
		runess[i], runess[j] = runess[j], runess[i]
	}
	outputInCorrectBase = string(runess)

	new, err := strconv.Atoi(outputInCorrectBase)

	if err == nil {
		return new
	}

	return 1
}

func main() {
	fmt.Println(convertNumber("1010101101002", 2, 3))
}
