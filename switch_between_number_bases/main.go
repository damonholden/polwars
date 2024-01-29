package main

import (
	"math"
	"strconv"
	"strings"
)

func convertNumber(number string, inputBase int, outputBase int) string {

	if inputBase == outputBase {
		return number
	}

	reversedNumber := reverseNumber(number)

	baseTenOutput := convertNumberToBaseTen(reversedNumber, inputBase)

	if outputBase == 10 {
		return baseTenOutput
	}

	// convert to required base
	outputInCorrectBase := convertDenaryNumberToRequiredBase(baseTenOutput, outputBase)

	runess := []rune(outputInCorrectBase)
	for i, j := 0, len(outputInCorrectBase)-1; i < j; i, j = i+1, j-1 {
		runess[i], runess[j] = runess[j], runess[i]
	}
	outputInCorrectBase = string(runess)

	// remove leading zeros as not needed in any base
	outputInCorrectBase = strings.TrimLeft(outputInCorrectBase, "0")

	return outputInCorrectBase

}

func reverseNumber(number string) string {
	runes := []rune(number)
	for i, j := 0, len(number)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func convertNumberToBaseTen(reversedNumber string, base int) string {
	var convertedNumber int
	for i := 0; i < len(reversedNumber); i++ {
		convertedNumber += convertCharCodeToInt(int(reversedNumber[i])) * int(math.Pow(float64(base), float64(i)))
	}
	return strconv.Itoa(convertedNumber)
}

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

func convertDenaryNumberToRequiredBase(number string, base int) string {
	numberAsInt, error := strconv.Atoi(number)

	if error != nil {
		return "Error"
	}

	var output = ""

	for numberAsInt > 1 {
		output += strconv.Itoa(numberAsInt % base)
		numberAsInt = numberAsInt - numberAsInt%base
		numberAsInt = numberAsInt / base
	}

	output += strconv.Itoa(numberAsInt % base)

	return output
}
