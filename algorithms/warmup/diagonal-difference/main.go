// This is the main package for this module
package main

/*
 * Complete the 'diagonalDifference' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts 2D_INTEGER_ARRAY arr as parameter.
 */

func diagonalDifference(arr [][]int32) int32 {
	leftSum := int32(0)
	rightSum := int32(0)
	squareLength := len(arr)

	// calculate leftSum -> O(n)
	// only sum numbers from the diagonal
	for row, column := 0, 0; row < squareLength; row, column = row+1, column+1 {
		leftSum += arr[row][column]
	}

	// calculate rightSum -> O(n)
	// only sum numbers from the diagonal
	for row, column := 0, squareLength-1; row < squareLength; row, column = row+1, column-1 {
		rightSum += arr[row][column]
	}

	// get absolute value since we can't use the `math` package
	diagonalDiff := leftSum - rightSum
	if diagonalDiff < 0 {
		diagonalDiff = diagonalDiff * -1
	}

	return diagonalDiff
}
