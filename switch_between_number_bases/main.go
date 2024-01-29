package main

import (
	"math"
	"strconv"
	"strings"
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

func convertDenaryNumberToRequiredBase(number int, base int) string {
	var output = ""
	for number > 1 {

		output += strconv.Itoa(number % base)
		number = number - number%base
		number = number / base
	}

	output += strconv.Itoa(number % base)

	return output
}

func convertNumber(number string, inputBase int, outputBase int) string {
	var output int

	if inputBase == outputBase {
		return number
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

		strOutput := strconv.Itoa(output)

		return strOutput
	}

	// convert to required base
	outputInCorrectBase := convertDenaryNumberToRequiredBase(output, outputBase)

	runess := []rune(outputInCorrectBase)
	for i, j := 0, len(outputInCorrectBase)-1; i < j; i, j = i+1, j-1 {
		runess[i], runess[j] = runess[j], runess[i]
	}
	outputInCorrectBase = string(runess)

	// remove leading zeros as not needed in any base
	outputInCorrectBase = strings.TrimLeft(outputInCorrectBase, "0")

	return outputInCorrectBase

}
