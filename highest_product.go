package highest_product

import (
	"errors"
	"fmt"
	"sort"
)

const NumFactors = 3

/* Input must be of length >= 3 */

// Loops through the numbers one by one and checks if each is among the NumFactor highest
func HighestProduct(numbers []int) (int, error) {
	if len(numbers) < NumFactors {
		errorMsg := fmt.Sprintf("Number list must be longer than %d (%d given)", NumFactors, len(numbers))
		return 0, errors.New(errorMsg)
	}

	// Initialize factor array to the first factors in numbers
	var factors [NumFactors]int
	for j := 0; j < NumFactors; j++ {
		factors[j] = numbers[j]
	}

	var smallestFactor int
	var smallestFactorIndex int
	for i := NumFactors; i < len(numbers); i++ {
		// Determine the smallest stored factor
		smallestFactor = factors[0]
		smallestFactorIndex = 0
		for j := 1; j < NumFactors; j++ {
			if smallestFactor > factors[j] {
				smallestFactor = factors[j]
				smallestFactorIndex = j
			}
		}

		if numbers[i] > smallestFactor {
			factors[smallestFactorIndex] = numbers[i] // Replace the smallest factor
		}
	}

	product := 1
	for i := 0; i < NumFactors; i++ {
		product *= factors[i]
	}
	return product, nil
}

/* Loops through the array NumFactors times and picks out the highest, second highest, third highest, etc... elements */
/* This one got a little dirty to handle som edge cases, but i primarily wanted to show that it is slower then HighestProduct */
func HighestProductSlower(numbers []int) (int, error) {
	if len(numbers) < NumFactors {
		errorMsg := fmt.Sprintf("Number list must be longer than %d (%d given)", NumFactors, len(numbers))
		return 0, errors.New(errorMsg)
	}

	smallestNumber := numbers[0]
	for i := 1; i < len(numbers); i++ {
		if numbers[i] < smallestNumber {
			smallestNumber = numbers[i]
		}
	}

	product := 1
	var currentBiggestNumber int
	var currentBiggestNumberIndex int
	var indexesToSkip [NumFactors]int
	for n := 0; n < NumFactors; n++ {
		indexesToSkip[n] = -1
	}
	for n := 0; n < NumFactors; n++ {
		if inArray(indexesToSkip, 0) {
			currentBiggestNumber = smallestNumber
		} else {
			currentBiggestNumber = numbers[0]
		}
		currentBiggestNumberIndex = 0
		for i := 0; i < len(numbers); i++ {
			if inArray(indexesToSkip, i) {
				continue
			}
			if numbers[i] > currentBiggestNumber {
				currentBiggestNumber = numbers[i]
				currentBiggestNumberIndex = i
			}
		}
		product *= currentBiggestNumber
		indexesToSkip[n] = currentBiggestNumberIndex
	}
	return product, nil
}

func inArray(arr [NumFactors]int, val int) bool {
	for j := 0; j < NumFactors; j++ {
		if val == arr[j] {
			return true
		}
	}
	return false
}

// Very slow
func HighestProductUsingSort(numbers []int) (int, error) {
	if len(numbers) < NumFactors {
		errorMsg := fmt.Sprintf("Number list must be longer than %d (%d given)", NumFactors, len(numbers))
		return 0, errors.New(errorMsg)
	}

	// Make a copy of the input slice to avoid changing the input. This is a big performance hit
	numbersCopy := make([]int, len(numbers))
	copy(numbersCopy, numbers)

	sort.Ints(numbersCopy) // Sorting is no faster than O(n lg(n))

	product := 1
	for i := 0; i < NumFactors; i++ {
		product *= numbersCopy[len(numbersCopy) - 1 - i]
	}

	return product, nil
}
