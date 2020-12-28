package arithmetic

import "errors"

// Add adds two numbers
func Add(nums ...int) (int, error) {
	if len(nums) < 2 {
		return 0, errors.New("A minimum of 2 numbers is required")
	}
	sum := 0
	for _, val := range nums {
		sum += val
	}
	return sum, nil
}

// Subtract subtracts the second operand from the first operand
func Subtract(num1 int, num2 int) int {
	return num1 - num2
}
