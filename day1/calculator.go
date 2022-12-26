package calculator

import (
	"errors"
	"math"
)

const ERR_DIVISION_BY_ZERO = "Division by zero"
const SQUARE_ROOT_INVALID_FOR_NEGATIVE_NUMBERS = "Square root is invalid for negative numbers"

func roundOff(num float64, decimal float64) float64 {
	scale := math.Pow(10, decimal)
	return math.Round((num * scale)) / scale
}
func add(num1 float32, num2 float32) float32 {
	result := num1 + num2
	return result
}
func mult(num1 float32, num2 float32) float32 {
	result := num1 * num2
	return result
}
func divide(num1 float64, num2 float64) (float64, error) {
	result := num1 / num2
	if num2 == 0 {
		return 0, errors.New(ERR_DIVISION_BY_ZERO)
	}
	return result, nil
}
func Sin(angle float64) float64 {
	result := math.Sin(angle)
	return roundOff(result, 2)
}
func Cos(angle float64) float64 {
	result := math.Cos(angle)
	return roundOff(result, 2)
}
func Tan(angle float64) (float64, error) {
	result, err := divide(Sin(angle), Cos(angle))
	return roundOff(result, 2), err
}
func SquareRoot(num float64) (float64, error) {
	if num < 0 {
		return 0, errors.New(SQUARE_ROOT_INVALID_FOR_NEGATIVE_NUMBERS)
	}
	result := math.Sqrt(num)
	return result, nil
}
