package main

import (
	"testing"
)

/**
 * Test Case: Hacker Rank Sample Input
 * Input: 07:05:45PM
 * Expects Return: 19:05:45
 */
func TestHackerRankSampleInput(t *testing.T) {
	input := "07:05:45PM"
	expectedResult := "19:05:45"
	result := timeConversion(input)

	if result != expectedResult {
		t.Errorf(`Result = %v, Expected Result = %v`, result, expectedResult)
  }
}

/**
 * Test Case: Example one
 * Input: 09:03:03AM
 * Expects Return: 09:03:03
 */
 func TestExampleOne(t *testing.T) {
	input := "09:03:03AM"
	expectedResult := "09:03:03"
	result := timeConversion(input)

	if result != expectedResult {
		t.Errorf(`Result = %v, Expected Result = %v`, result, expectedResult)
  }
}

/**
 * Test Case: Example Two
 * Input: 00:00:00AM
 * Expects Return: 00:00:00
 */
 func TestExampleTwo(t *testing.T) {
	input := "00:00:00AM"
	expectedResult := "00:00:00"
	result := timeConversion(input)

	if result != expectedResult {
		t.Errorf(`Result = %v, Expected Result = %v`, result, expectedResult)
  }
}

/**
 * Test Case: Example Three
 * Input: 12:01:00PM
 * Expects Return: 12:01:00
 */
 func TestExampleThree(t *testing.T) {
	input := "12:01:00PM"
	expectedResult := "12:01:00"
	result := timeConversion(input)

	if result != expectedResult {
		t.Errorf(`Result = %v, Expected Result = %v`, result, expectedResult)
  }
}

/**
 * Test Case: Example Four
 * Input: 12:01:00AM
 * Expects Return: 00:01:00
 */
 func TestExampleFour(t *testing.T) {
	input := "12:01:00AM"
	expectedResult := "00:01:00"
	result := timeConversion(input)

	if result != expectedResult {
		t.Errorf(`Result = %v, Expected Result = %v`, result, expectedResult)
  }
}
