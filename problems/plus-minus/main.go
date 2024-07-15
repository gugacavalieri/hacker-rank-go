// This is the main package for this module
package main

import (
	"fmt"
)

/*
 * Complete the 'plusMinus' function below.
 *
 * The function accepts INTEGER_ARRAY arr as parameter.
 *
 * This function will count the positive, negative and zero numbers in a array
 * Then it will calculate it's ratio
 * Complexity: O(n) -> Single Linear Loop through the array
 */
func plusMinus(arr []int32) (string, string, string) {
	arrLength := len(arr)
	arrPositiveCount := 0
	arrNegativeCount := 0
	arrZeroCount := 0

	/* interate to find counts */
	for _, element := range arr {
		if element > 0 {
			arrPositiveCount += 1
		}
		if element < 0 {
			arrNegativeCount += 1
		}
		if element == 0 {
			arrZeroCount += 1
		}
	}

	/* calculate ratios */
	arrPositiveRatio := float32(arrPositiveCount) / float32(arrLength)
	arrNegativeRatio := float32(arrNegativeCount) / float32(arrLength)
	arrZeroRatio := float32(arrZeroCount) / float32(arrLength)

	/* print ratios */
	printPrecision := "%.6f\n"
	fmt.Printf(printPrecision, arrPositiveRatio)
	fmt.Printf(printPrecision, arrNegativeRatio)
	fmt.Printf(printPrecision, arrZeroRatio)

	return fmt.Sprintf(printPrecision, arrPositiveRatio), fmt.Sprintf(printPrecision, arrNegativeRatio), fmt.Sprintf(printPrecision, arrZeroRatio)
}
