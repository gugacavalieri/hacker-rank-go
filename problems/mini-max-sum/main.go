// This is the main package for this module
package main

import (
	"fmt"
)

/*
 * Complete the 'miniMaxSum' function below.
 *
 * The function accepts INTEGER_ARRAY arr as parameter.
 *
 * Here, we sum all the numbers in the array then we subtract the min and max numbers since we always need 4/5 numbers.
 * Complexity: O(n)
 */
func miniMaxSum(arr []int32) (int64, int64) {

	minNumber := arr[0]
	maxNumber := int32(0)
	totalSum := int64(0)

	for _, element := range arr {
		if element < minNumber {
			minNumber = element
		}
		if element > maxNumber {
			maxNumber = element
		}
		totalSum += int64(element)
	}

	var minSum = int64(totalSum) - int64(maxNumber)
	var maxSum = int64(totalSum) - int64(minNumber)

	fmt.Println(minSum, maxSum)

	return minSum, maxSum
}
