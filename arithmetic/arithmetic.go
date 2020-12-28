package arithmetic

// Add adds two numbers
func Add(nums ...int) int {
	sum := 0
	for _, val := range nums {
		sum += val
	}
	return sum
}

// Subtract subtracts the second operand from the first operand
func Subtract(num1 int, num2 int) int {
	return (num1 - num2)
}
