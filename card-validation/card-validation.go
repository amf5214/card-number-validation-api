package cardvalidation

import (
	"strconv"
)

// Function to validate whether a card number is valid
func ValidateCardNumber(cardNumber string) (isValid bool) {

	// Count the number of characters in the string
	count := len([]rune(cardNumber))

	// If there are not 16 characters then it is automatically not valid
	if count != 16 {
		return false
	}

	// Create an array to hold the number values
	var number = [16]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	// Iterate over the string
	for i := 0; i < count; i++ {

		character, _ := strconv.Atoi(string([]rune(cardNumber)[i]))
		number[i] = character

		if i%2 == 0 {
			factor := 2 * number[i]

			if factor > 9 {
				character0, _ := strconv.Atoi(string([]rune(strconv.Itoa(factor))[0]))
				character1, _ := strconv.Atoi(string([]rune(strconv.Itoa(factor))[1]))
				number[i] = character0 + character1
			} else {
				number[i] = factor
			}
		}
	}

	isValid = sumArray(number)%10 == 0

	return
}

// Function to convert an array of 16 integers into a single string
func ConvertArrayToString(array [16]int) string {
	size := len(array)

	str := ""

	for i := 0; i < size; i++ {
		str = str + strconv.Itoa(array[i])
	}

	return str
}

// Function to sum the values in an array of 16 integers
func sumArray(numbers [16]int) int {
	result := 0
	for i := 0; i < len(numbers); i++ {
		result += numbers[i]
	}
	return result
}
