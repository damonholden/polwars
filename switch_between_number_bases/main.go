package main

import (
	"errors"
	"math"
	"strconv"
	"strings"
)

func convertNumber(number string, inputBase int, outputBase int) (string, error) {
	if inputBase == outputBase {
		return number, nil
	}

	if inputBase != 10 {
		numberConvertedToBaseTen, error := convertNumberToBaseTen(number, inputBase)

		if error != nil {
			return "", error
		}

		number = numberConvertedToBaseTen
	}

	if outputBase == 10 {
		return number, nil
	}

	numberConvertedToOutputBase := convertBaseTenNumberToRequiredBase(number, outputBase)

	return numberConvertedToOutputBase, nil

}

func reverseNumber(number string) string {
	runes := []rune(number)
	for i, j := 0, len(number)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func convertNumberToBaseTen(number string, base int) (string, error) {
	var numberConvertedToBaseTen int

	for i, j := len(number)-1, 0; i >= 0; i, j = i-1, j+1 {
		charAsBaseTenValue, error := convertCharCodeToBaseTenValue(int(number[i]))

		if error != nil {
			return "", error
		}

		multiplier := int(math.Pow(float64(base), float64(j)))

		numberConvertedToBaseTen += charAsBaseTenValue * multiplier
	}
	return strconv.Itoa(numberConvertedToBaseTen), nil
}

func charCodeIsInNumberCharCodeRange(charCode int) bool {
	return charCode > 47 && charCode < 58
}

func charCodeIsInLetterCharCodeRange(charCode int) bool {
	return charCode > 64 && charCode < 91
}

func convertCharCodeToBaseTenValue(charCode int) (int, error) {
	if charCodeIsInNumberCharCodeRange(charCode) {
		return charCode - 48, nil
	}

	if charCodeIsInLetterCharCodeRange(charCode) {
		return charCode - 55, nil
	}

	return 0, errors.New("invalid char code. Must be between 0 and 9 or A and Z")
}

func convertBaseTenNumberToRequiredBase(number string, base int) string {
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

	output = reverseNumber(output)
	output = strings.TrimLeft(output, "0")

	return output
}
