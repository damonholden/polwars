package main

import (
	"testing"
)

func TestBaseTwoToThree(t *testing.T) {
	expected := "21112000"
	output := convertNumber("1010101101001", 2, 3)

	if output != expected {
		t.Errorf("Expected %s, got %s", expected, output)
	}
}

func TestBaseThreeToTwo(t *testing.T) {
	expected := "1010101101001"
	output := convertNumber("21112000", 3, 2)

	if output != expected {
		t.Errorf("Expected %s, got %s", expected, output)
	}
}

func TestBaseTwoToTen(t *testing.T) {
	expected := "3291"
	output := convertNumber("110011011011", 2, 10)

	if output != expected {
		t.Errorf("Expected %s, got %s", expected, output)
	}
}

func TestBaseTenToTwo(t *testing.T) {
	expected := "10001110110010011101011000110000100010010010"
	output := convertNumber("9812345817234", 10, 2)

	if output != expected {
		t.Errorf("Expected %s, got %s", expected, output)
	}
}