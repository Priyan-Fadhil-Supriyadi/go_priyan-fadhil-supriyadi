package calculate

import (
	"testing"
)

func TestAddition(t *testing.T) {
	if Addition(1, 2) != 3 {
		t.Error("Expected 1 (+) 2 to equal 3")
	}
	if Addition(-1, -2) != -3 {
		t.Error("Expected -1 (+) -2 to equal -3")
	}
}

func TestSubtraction(t *testing.T) {
	if Substract(2, 1) != 1 {
		t.Error("Expected 2 (-) 1 to equal 1")
	}
	if Substract(1, -2) != 3 {
		t.Error("Expected 1 (-) -2 to equal 3")
	}
}

func TestMultiplication(t *testing.T) {
	if Mult(1, 2) != 2 {
		t.Error("Expected 1 (*) 2 to equal 2")
	}
	if Mult(-1, 2) != -2 {
		t.Error("Expected -1 (*) 2 to equal -2")
	}
}

func TestDivision(t *testing.T) {
	if Div(4, 2) != 2 {
		t.Error("Expected 4 (/) 2 to equal 2")
	}
	if Div(0, 3) != 0 {
		t.Error("Expected 0 (/) 3 to equal 0")
	}
	if Div(-4, 2) != -2 {
		t.Error("Expected -4 (/) 2 to equal -2")
	}
}
