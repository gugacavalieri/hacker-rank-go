package main

import (
	"testing"
)

/**
 * Test Case: Empty Matrix
 * Input: [][]
 * Expects Return: 0
 */
func TestDiagonalDifferenceEmptyMatrix(t *testing.T) {
	input := [][]int32{}
	expectedResult := int32(0)
	result := diagonalDifference(input)

	if result != expectedResult {
		t.Errorf(`Result = %v, Expected Result = %v`, result, expectedResult)
	}
}

/**
 * Test Case: Hacker Rank #1
 * Input: {11, 2, 4}, 
          {4, 5, 6}, 
          {10, 8, -12}
 * Expects Return: 15
 */
func TestDiagonalDifferenceHackerRankCaseOne(t *testing.T) {
	input := [][]int32{{11, 2, 4}, {4, 5, 6}, {10, 8, -12}}
	expectedResult := int32(15)
	result := diagonalDifference(input)

	if result != expectedResult {
		t.Errorf(`Result = %v, Expected Result = %v`, result, expectedResult)
	}
}

/**
 * Test Case: Matrix with only zeros
 * Input: {0, 0}, 
          {0, 0}
 * Expects Return: 0
 */
 func TestDiagonalDifferenceZeros(t *testing.T) {
	input := [][]int32{{0, 0}, {0, 0}}
	expectedResult := int32(0)
	result := diagonalDifference(input)

	if result != expectedResult {
		t.Errorf(`Result = %v, Expected Result = %v`, result, expectedResult)
	}
}

/**
 * Test Case: Matrix with negative values
 * Input: {-1, -2, -3}, 
          {-4, -5, -6}, 
          {-7, -8, -9}
 * Expects Return: 0
 */
 func TestDiagonalDifferenceNegativeValues(t *testing.T) {
	input := [][]int32{{-1, -2, -3}, {-4, -5, -6}, {-7, -8, -9}}
	expectedResult := int32(0)
	result := diagonalDifference(input)

	if result != expectedResult {
		t.Errorf(`Result = %v, Expected Result = %v`, result, expectedResult)
	}
}

/**
 * Test Case: 4x4 Matrix 
 * Input: {1, 4, 15, 22}, 
          {33, 44, 55, 51}, 
          {50, 31, 23, 45}, 
          {66, 44, 12, 123}
 * Expects Return: 191 - 174 = 17
 */
 func TestDiagonalDifference4By4Matrix(t *testing.T) {
	input := [][]int32{{1, 4, 15, 22}, {33, 44, 55, 51}, {50, 31, 23, 45}, {66, 44, 12, 123}}
	expectedResult := int32(17)
	result := diagonalDifference(input)

	if result != expectedResult {
		t.Errorf(`Result = %v, Expected Result = %v`, result, expectedResult)
	}
}