package calculator

import (
	"math"
	"testing"
)

func TestSuccessfulAdditionTwoNumbers(t *testing.T) {
	result := add(10, 5)
	if result != 15 {
		t.Error("wrong")
	}
}
func TestSuccessfulAdditionTwoNumbersWithOppositeSign(t *testing.T) {
	result := add(-10, 5.12)
	if result != -4.88 {
		t.Error("wrong")
	}
}
func TestSuccessfulAdditionTwoNegativeNumbers(t *testing.T) {
	result := add(-2, -12)
	if result != -14 {
		t.Error("wrong")
	}
}

func TestSuccessfulMultiplicationTwoNumbers(t *testing.T) {
	result := mult(10, 5)
	if result != 50 {
		t.Error("wrong")
	}
}
func TestSuccessfulMultiplicationTwoNegativeNumbers(t *testing.T) {
	result := mult(-2, -2)
	if result != 4 {
		t.Error("wrong")
	}
}

func TestSuccessfulDivisionTwoNumbers(t *testing.T) {
	result, _ := divide(10, 5)
	if result != 2 {
		t.Error("wrong")
	}
}
func TestDoesNotAllowDivisionByZero(t *testing.T) {
	_, err := divide(10, 0)
	if err.Error() != ERR_DIVISION_BY_ZERO {
		t.Error("Division by zero allowed")
	}
}

func TestSin0Degrees(t *testing.T) {
	result := Sin(0)
	if result != 0 {
		t.Errorf("fail %f", result)
	}
}
func TestSin90Degrees(t *testing.T) {
	NinetyDegrees, _ := divide(math.Pi, 2)
	result := Sin(NinetyDegrees)
	if result != 1 {
		t.Errorf("fail %f", result)
	}
}
func TestSin45Degrees(t *testing.T) {
	FortyFiveDegrees, _ := divide(math.Pi, 4)
	result := Sin(FortyFiveDegrees)
	if result != 0.71 {
		t.Errorf("fail %f", result)
	}
}
func TestSinOfInfinityDegrees(t *testing.T) {
	result := Sin(math.Inf(+1))
	if !math.IsNaN(result) {
		t.Errorf("fail %f", result)
	}
}
func TestSinOfNegativeDegrees(t *testing.T) {
	result := Sin(-math.Pi / 2)
	if result != -1 {
		t.Errorf("fail %f", result)
	}
}

func TestCos0degrees(t *testing.T) {
	result := Cos(0)
	if result != 1 {
		t.Errorf("fail %f", result)
	}
}
func TestCos90Degree(t *testing.T) {
	NinetyDegrees, _ := divide(math.Pi, 2)
	result := Cos(NinetyDegrees)
	if result != 0 {
		t.Errorf("fail %f", result)
	}
}
func TestCos45Degree(t *testing.T) {
	FortyFiveDegrees, _ := divide(math.Pi, 4)
	result := Cos(FortyFiveDegrees)
	if result != 0.71 {
		t.Errorf("fail %f", result)
	}
}
func TestCosOfInfinityDegrees(t *testing.T) {
	result := Cos(math.Inf(+1))
	if !math.IsNaN(result) {
		t.Errorf("fail %f", result)
	}
}
func TestCosOfNegativeDegrees(t *testing.T) {
	result := Cos(-math.Pi)
	if result != -1 {
		t.Errorf("fail %f", result)
	}
}

func TestTanOf0Degrees(t *testing.T) {
	result, err := Tan(0)
	if result != 0 && err != nil {
		t.Errorf("fail")
	}
}
func TestTan45Degrees(t *testing.T) {
	FortyFiveDegrees, _ := divide(math.Pi, 4)
	result, err := Tan(FortyFiveDegrees)
	if result != 1 && err != nil {
		t.Errorf("fail")
	}
}
func TestTan90Degrees(t *testing.T) {
	NinetyDegrees, _ := divide(math.Pi, 2)
	_, err := Tan(NinetyDegrees)
	if err.Error() != ERR_DIVISION_BY_ZERO {
		t.Errorf("TAN 90 degrees result division by zero error")

	}
}
func TestTanOfInfinityDegree(t *testing.T) {
	result, err := Tan(math.Inf(+1))
	if !math.IsNaN(result) && err != nil {
		t.Errorf("fail %f", result)
	}
}

func TestSquareRootOfPositiveNumber(t *testing.T) {
	result, err := SquareRoot(4)

	if result != 2 && err != nil {
		t.Error("Fail")
	}
}
func TestSquareRootOfNegativeNumber(t *testing.T) {
	_, err := SquareRoot(-4)

	if err.Error() != SQUARE_ROOT_INVALID_FOR_NEGATIVE_NUMBERS {
		t.Error("Fail")
	}
}
