// This is the main package for this module
package main

/*
 * Complete the 'aVeryBigSum' function below.
 *
 * The function is expected to return a LONG_INTEGER.
 * The function accepts LONG_INTEGER_ARRAY ar as parameter.
 */
func aVeryBigSum(ar []int64) int64 {
	// Write your code here
	sum := int64(0)

	for _, element := range ar {
		sum += element
	}

	return sum
}
