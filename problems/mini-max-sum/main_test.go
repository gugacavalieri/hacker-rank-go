package main

import (
	"testing"
)

/**
 * Test Case: Sequencial Array
 * Input: [1, 2, 3, 4, 5]
 * Expects Return: 10, 14
 */
func TestMiniMaxSumDefaultSample(t *testing.T) {
	input := []int32{1, 2, 3, 4, 5}
	expectedMin := int64(10)
	expetedMax := int64(14)
	resultMin, resultMax := miniMaxSum(input)

	if resultMin != expectedMin {
		t.Errorf(`Result = %v, Expected Result = %v`, resultMin, expectedMin)
	}
	if resultMax != expetedMax {
		t.Errorf(`Result = %v, Expected Result = %v`, resultMax, expetedMax)
	}
}

/**
 * Test Case: First Example Array
 * Input: [1, 3, 5, 7, 9]
 * Expects Return: 10, 14
 */
func TestMiniMaxSumFirstExample(t *testing.T) {
	input := []int32{1, 3, 5, 7, 9}
	expectedMin := int64(16)
	expetedMax := int64(24)
	resultMin, resultMax := miniMaxSum(input)

	if resultMin != expectedMin {
		t.Errorf(`Result = %v, Expected Result = %v`, resultMin, expectedMin)
	}
	if resultMax != expetedMax {
		t.Errorf(`Result = %v, Expected Result = %v`, resultMax, expetedMax)
	}
}

/**
 * Test Case: First Example Array
 * Input: [7, 69, 2, 221, 8974]
 * Expects Return: 10, 14
 */
func TestMiniMaxSumSampleTestOne(t *testing.T) {
	input := []int32{7, 69, 2, 221, 8974}
	expectedMin := int64(299)
	expetedMax := int64(9271)
	resultMin, resultMax := miniMaxSum(input)

	if resultMin != expectedMin {
		t.Errorf(`Result = %v, Expected Result = %v`, resultMin, expectedMin)
	}
	if resultMax != expetedMax {
		t.Errorf(`Result = %v, Expected Result = %v`, resultMax, expetedMax)
	}
}

/**
 * Test Case: First Example Array
 * Input: [140638725, 436257910, 953274816, 734065819, 362748590]
 * Expects Return: 1673711044, 2486347135
 */
func TestMiniMaxSumSampleTestThree(t *testing.T) {
	input := []int32{140638725, 436257910, 953274816, 734065819, 362748590}
	expectedMin := int64(1673711044)
	expetedMax := int64(2486347135)
	resultMin, resultMax := miniMaxSum(input)

	if resultMin != expectedMin {
		t.Errorf(`Result = %v, Expected Result = %v`, resultMin, expectedMin)
	}
	if resultMax != expetedMax {
		t.Errorf(`Result = %v, Expected Result = %v`, resultMax, expetedMax)
	}
}
