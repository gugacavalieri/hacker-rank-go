package main

import (
	"testing"
)

/**
 * Test Case: Empty Array
 * Input: []
 * Expects Return: 0, 0, 0
 */
func TestHackerRankEmptyArray(t *testing.T) {
	input := []int32{}
	expectedPositive := "NaN\n"
	expectedNegative := "NaN\n"
	expectedZero := "NaN\n"
	resultPositive, resultNegative, resultZero := plusMinus(input)

	if resultPositive != expectedPositive {
		t.Errorf(`Result = %v, Expected Result = %v`, resultPositive, expectedPositive)
	}
	if resultNegative != expectedNegative {
		t.Errorf(`Result = %v, Expected Result = %v`, resultNegative, expectedNegative)
	}
	if resultZero != expectedZero {
		t.Errorf(`Result = %v, Expected Result = %v`, resultZero, expectedZero)
	}
}

/**
 * Test Case: Hacker Rank Example #1
 * Input: [1, 1, 0, -1, -1]
 * Expects Return: 0.400000, 0.400000, 0.200000
 */
func TestHackerRankExampleOne(t *testing.T) {
	input := []int32{1, 1, 0, -1, -1}
	expectedPositive := "0.400000\n"
	expectedNegative := "0.400000\n"
	expectedZero := "0.200000\n"
	resultPositive, resultNegative, resultZero := plusMinus(input)

	if resultPositive != expectedPositive {
		t.Errorf(`Result = %v, Expected Result = %v`, resultPositive, expectedPositive)
	}
	if resultNegative != expectedNegative {
		t.Errorf(`Result = %v, Expected Result = %v`, resultNegative, expectedNegative)
	}
	if resultZero != expectedZero {
		t.Errorf(`Result = %v, Expected Result = %v`, resultZero, expectedZero)
	}
}

/**
 * Test Case: Hacker Rank Test Case #0
 * Input: [-4, 3, -9, 0, 4, 1]
 * Expects Return: 0.500000, 0.333333, 0.166667
 */
func TestHackerRankTestCaseZero(t *testing.T) {
	input := []int32{-4, 3, -9, 0, 4, 1}
	expectedPositive := "0.500000\n"
	expectedNegative := "0.333333\n"
	expectedZero := "0.166667\n"
	resultPositive, resultNegative, resultZero := plusMinus(input)

	if resultPositive != expectedPositive {
		t.Errorf(`Result = %v, Expected Result = %v`, resultPositive, expectedPositive)
	}
	if resultNegative != expectedNegative {
		t.Errorf(`Result = %v, Expected Result = %v`, resultNegative, expectedNegative)
	}
	if resultZero != expectedZero {
		t.Errorf(`Result = %v, Expected Result = %v`, resultZero, expectedZero)
	}
}

/**
 * Test Case: Hacker Rank Test Case #1
 * Input: [1, 2, 3, -1, -2, -3, 0, 0]
 * Expects Return: 0.375000, 0.375000, 0.250000
 */
func TestHackerRankTestCaseOne(t *testing.T) {
	input := []int32{1, 2, 3, -1, -2, -3, 0, 0}
	expectedPositive := "0.375000\n"
	expectedNegative := "0.375000\n"
	expectedZero := "0.250000\n"
	resultPositive, resultNegative, resultZero := plusMinus(input)

	if resultPositive != expectedPositive {
		t.Errorf(`Result = %v, Expected Result = %v`, resultPositive, expectedPositive)
	}
	if resultNegative != expectedNegative {
		t.Errorf(`Result = %v, Expected Result = %v`, resultNegative, expectedNegative)
	}
	if resultZero != expectedZero {
		t.Errorf(`Result = %v, Expected Result = %v`, resultZero, expectedZero)
	}
}

/**
 * Test Case: Hacker Rank Test Case #4
 * Input: [55, 48, 48, 45, 91, 97, 45, 1, 39, 54, 36, 6, 19, 35, 66, 36, 72, 93, 38, 21, 65, 70, 36, 63, 39, 76, 82, 26, 67, 29, 24, 82, 62, 53, 1, 50, 47, 65, 67, 19, 66, 90, 77]
 * Expects Return: 1.000000, 0.000000, 0.000000
 */
func TestHackerRankTestCaseFour(t *testing.T) {
	input := []int32{55, 48, 48, 45, 91, 97, 45, 1, 39, 54, 36, 6, 19, 35, 66, 36, 72, 93, 38, 21, 65, 70, 36, 63, 39, 76, 82, 26, 67, 29, 24, 82, 62, 53, 1, 50, 47, 65, 67, 19, 66, 90, 77}
	expectedPositive := "1.000000\n"
	expectedNegative := "0.000000\n"
	expectedZero := "0.000000\n"
	resultPositive, resultNegative, resultZero := plusMinus(input)

	if resultPositive != expectedPositive {
		t.Errorf(`Result = %v, Expected Result = %v`, resultPositive, expectedPositive)
	}
	if resultNegative != expectedNegative {
		t.Errorf(`Result = %v, Expected Result = %v`, resultNegative, expectedNegative)
	}
	if resultZero != expectedZero {
		t.Errorf(`Result = %v, Expected Result = %v`, resultZero, expectedZero)
	}
}
