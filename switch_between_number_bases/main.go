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

	err := checkForErrorsInInput(number, inputBase, outputBase)

	if err != nil {
		return "", err
	}

	if inputBase != 10 {
		number = convertNumberToBaseTen(number, inputBase)
	}

	if outputBase == 10 {
		return number, nil
	}

	numberConvertedToOutputBase := convertBaseTenNumberToRequiredBase(number, outputBase)

	return numberConvertedToOutputBase, nil
}

func checkForErrorsInInput(number string, inputBase int, outputBase int) error {
	if !baseIsValid(inputBase) {
		return errors.New("input base must be between 2 and 36")
	}

	if !baseIsValid(outputBase) {
		return errors.New("output base must be between 2 and 36")
	}

	charErr := checkIfAllCharactersCanBeHandled(number)

	if charErr != nil {
		return charErr
	}

	numberErr := checkIfNumberIsValidForBase(number, inputBase)

	if numberErr != nil {
		return numberErr
	}

	return nil
}

func baseIsValid(base int) bool {
	if base < 2 || base > 36 {
		return false
	}

	return true
}

func checkIfNumberIsValidForBase(number string, inputBase int) error {
	for _, char := range number {
		if convertCharCodeToBaseTenValue(int(char)) >= inputBase {
			return errors.New("digit is too large for given base")
		}
	}

	return nil
}

func checkIfAllCharactersCanBeHandled(number string) error {
	for _, char := range number {
		if !charCodeIsInLetterCharCodeRange(int(char)) && !charCodeIsInNumberCharCodeRange(int(char)) {
			return errors.New("invalid char code. Must be between 0 and 9 or A and Z")
		}
	}

	return nil
}

func convertNumberToBaseTen(number string, base int) string {
	var numberConvertedToBaseTen int

	for i, j := len(number)-1, 0; i >= 0; i, j = i-1, j+1 {
		charAsBaseTenValue := convertCharCodeToBaseTenValue(int(number[i]))

		multiplier := int(math.Pow(float64(base), float64(j)))

		numberConvertedToBaseTen += charAsBaseTenValue * multiplier
	}
	return strconv.Itoa(numberConvertedToBaseTen)
}

func convertCharCodeToBaseTenValue(charCode int) int {
	if charCodeIsInNumberCharCodeRange(charCode) {
		return charCode - 48
	}

	return charCode - 55
}

func convertBaseTenNumberToChar(number int) string {
	if number < 10 {
		return strconv.Itoa(number)
	}

	character := 'A' + rune(number-10)

	return string(character)
}

func charCodeIsInNumberCharCodeRange(charCode int) bool {
	return charCode > 47 && charCode < 58
}

func charCodeIsInLetterCharCodeRange(charCode int) bool {
	return charCode > 64 && charCode < 91
}

func convertBaseTenNumberToRequiredBase(number string, base int) string {
	numberAsInt, err := strconv.Atoi(number)

	if err != nil {
		return "Error"
	}

	var output = ""

	for numberAsInt > 1 {
		output += convertBaseTenNumberToChar(numberAsInt % base)
		numberAsInt = numberAsInt - numberAsInt%base
		numberAsInt = numberAsInt / base
	}

	output += strconv.Itoa(numberAsInt % base)

	output = reverseNumber(output)
	output = strings.TrimLeft(output, "0")

	return output
}

func reverseNumber(number string) string {
	runes := []rune(number)
	for i, j := 0, len(number)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
