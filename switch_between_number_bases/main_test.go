package main

import (
	"errors"
	"testing"
)

func TestBaseTenToSixteen(t *testing.T) {
	expected := "1B"
	output, _ := convertNumber("27", 10, 16)

	if output != expected {
		t.Errorf("Expected %s, got %s", expected, output)
	}
}

func TestBaseTwoToThree(t *testing.T) {
	expected := "21112000"
	output, _ := convertNumber("1010101101001", 2, 3)

	if output != expected {
		t.Errorf("Expected %s, got %s", expected, output)
	}
}

func TestBaseThreeToTwo(t *testing.T) {
	expected := "1010101101001"
	output, _ := convertNumber("21112000", 3, 2)

	if output != expected {
		t.Errorf("Expected %s, got %s", expected, output)
	}
}

func TestBaseTwoToTen(t *testing.T) {
	expected := "3291"
	output, _ := convertNumber("110011011011", 2, 10)

	if output != expected {
		t.Errorf("Expected %s, got %s", expected, output)
	}
}

func TestBaseTenToTwo(t *testing.T) {
	expected := "10001110110010011101011000110000100010010010"
	output, _ := convertNumber("9812345817234", 10, 2)

	if output != expected {
		t.Errorf("Expected %s, got %s", expected, output)
	}
}

func TestConversionOfLargeBaseThirtySixNumberToBaseTwo(t *testing.T) {
	expected := "10000001101111110000111111111111"
	output, _ := convertNumber("ZZZZZZ", 36, 2)

	if output != expected {
		t.Errorf("Expected %s, got %s", expected, output)
	}
}

func TestConversionOfBaseTwoNumberToBaseThirtySix(t *testing.T) {
	expected := "ZZZZZZ"
	output, _ := convertNumber("10000001101111110000111111111111", 2, 36)

	if output != expected {
		t.Errorf("Expected %s, got %s", expected, output)
	}
}

func TestThrownErrorIfInValidCharacter(t *testing.T) {
	expected := errors.New("invalid char code. Must be between 0 and 9 or A and Z")
	_, err := convertNumber(".", 2, 10)
	if err.Error() != expected.Error() {

		t.Errorf("Expected error %s, got %s", expected, err)
	}
}

func TestThrownErrorIfCharacterDoesNotFitInBase(t *testing.T) {
	expected := errors.New("digit is too large for given base")
	_, err := convertNumber("2", 2, 10)
	if err.Error() != expected.Error() {

		t.Errorf("Expected error %s, got %s", expected, err)
	}
}

func TestThrownErrorIfInputBaseIsTooLow(t *testing.T) {
	expected := errors.New("input base must be between 2 and 36")
	_, err := convertNumber("111", 1, 10)
	if err.Error() != expected.Error() {

		t.Errorf("Expected error %s, got %s", expected, err)
	}
}
func TestThrownErrorIfInputBaseIsTooHigh(t *testing.T) {
	expected := errors.New("input base must be between 2 and 36")
	_, err := convertNumber("111", 45, 10)
	if err.Error() != expected.Error() {

		t.Errorf("Expected error %s, got %s", expected, err)
	}
}

func TestThrownErrorIfOutputBaseIsTooLow(t *testing.T) {
	expected := errors.New("output base must be between 2 and 36")
	_, err := convertNumber("111", 2, 1)
	if err.Error() != expected.Error() {

		t.Errorf("Expected error %s, got %s", expected, err)
	}
}
func TestThrownErrorIfOutputBaseIsTooHigh(t *testing.T) {
	expected := errors.New("output base must be between 2 and 36")
	_, err := convertNumber("111", 2, 68)
	if err.Error() != expected.Error() {

		t.Errorf("Expected error %s, got %s", expected, err)
	}
}
