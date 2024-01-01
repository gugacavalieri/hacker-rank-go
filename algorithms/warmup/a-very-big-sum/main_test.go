package main

import (
	"testing"
)

/**
 * Test Case: Empty Array
 * Input: []
 * Expects Return: 0
 */
func TestAVeryBigSumEmptyArray(t *testing.T) {
	input := []int64{}
	expectedResult := int64(0)
	result := aVeryBigSum(input)

	if result != expectedResult {
		t.Errorf(`Result = %v, Expected Result = %v`, result, expectedResult)
	}
}

/**
 * Test Case: Array with One Element
 * Input: [1]
 * Expects Return: 1
 */
func TestAVeryBigSumOneElement(t *testing.T) {
	input := []int64{1}
	expectedResult := int64(1)
	result := aVeryBigSum(input)

	if result != expectedResult {
		t.Errorf(`Result = %v, Expected Result = %v`, result, expectedResult)
	}
}

/**
 * Test Case: Array with Three Elements
 * Input: [1,2,3]
 * Expects Return: 6
 */
func TestAVeryBigSumThreeElements(t *testing.T) {
	input := []int64{1, 2, 3}
	expectedResult := int64(6)
	result := aVeryBigSum(input)

	if result != expectedResult {
		t.Errorf(`Result = %v, Expected Result = %v`, result, expectedResult)
	}
}

/**
 * Test Case: Array with Large elements
 * Input: [1000000001, 1000000002, 1000000003, 1000000004, 1000000005]
 * Expects Return: 5000000015
 */
func TestAVeryBigSumLargeElements(t *testing.T) {
	input := []int64{1000000001, 1000000002, 1000000003, 1000000004, 1000000005}
	expectedResult := int64(5000000015)
	result := aVeryBigSum(input)

	if result != expectedResult {
		t.Errorf(`Result = %v, Expected Result = %v`, result, expectedResult)
	}
}

/**
 * Test Case: Array with Negative Numbers
 * Input: [5, -5]
 * Expects Return: 0
 */
func TestAVeryBigSumNegativeNumbers(t *testing.T) {
	input := []int64{5, -5}
	expectedResult := int64(0)
	result := aVeryBigSum(input)

	if result != expectedResult {
		t.Errorf(`Result = %v, Expected Result = %v`, result, expectedResult)
	}
}

/**
 * Test Case: Array with Two Negative Numbers
 * Input: [5, -5, -5]
 * Expects Return: -5
 */
func TestAVeryBigSumTwoNegativeNumbers(t *testing.T) {
	input := []int64{5, -5, -5}
	expectedResult := int64(-5)
	result := aVeryBigSum(input)

	if result != expectedResult {
		t.Errorf(`Result = %v, Expected Result = %v`, result, expectedResult)
	}
}
