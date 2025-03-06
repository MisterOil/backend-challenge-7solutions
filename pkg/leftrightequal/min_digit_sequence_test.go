package leftrightequal

import (
	"testing"
)

func TestMinDigitSequenceCase1(t *testing.T) {
	dataInput := "LLRR="
	findMinDigit, _ := MinDigitSequence(dataInput)
	expected := "210122"
	if findMinDigit != expected {
		t.Errorf("Expected %s, got %s", expected, findMinDigit)
	}
}

func TestMinDigitSequenceCase2(t *testing.T) {
	dataInput := "==RLL"
	findMinDigit, _ := MinDigitSequence(dataInput)
	expected := "000210"
	if findMinDigit != expected {
		t.Errorf("Expected %s, got %s", expected, findMinDigit)
	}
}

func TestMinDigitSequenceCase3(t *testing.T) {
	dataInput := "=LLRR"
	findMinDigit, _ := MinDigitSequence(dataInput)
	expected := "221012"
	if findMinDigit != expected {
		t.Errorf("Expected %s, got %s", expected, findMinDigit)
	}
}

func TestMinDigitSequenceCase4(t *testing.T) {
	dataInput := "RRL=R"
	findMinDigit, _ := MinDigitSequence(dataInput)
	expected := "012001"
	if findMinDigit != expected {
		t.Errorf("Expected %s, got %s", expected, findMinDigit)
	}
}
