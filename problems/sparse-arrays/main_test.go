package main

import (
	"reflect"
	"testing"
)

/**
 * Test Case: Hacker Rank Test Input Ten
 * Input: ["aba", "baba", "aba", "xzxb"], ["aba", "xzxb", "ab"]
 * Expects Return: [2, 1, 0]
 */
func TestHackerRankTestInputTen(t *testing.T) {
	inputStrings := []string{"aba", "baba", "aba", "xzxb"}
	inputQueries := []string{"aba", "xzxb", "ab"}
	expectedResult := []int32{2, 1, 0}
	result := matchingStrings(inputStrings, inputQueries)

	if !reflect.DeepEqual(result, expectedResult) {
		t.Errorf(`Result = %v, Expected Result = %v`, result, expectedResult)
	}
}

/**
 * Test Case: Hacker Rank Test Input Eleven
 * Input: ["def", "de", "fgh"], ["de", "lmn", "fgh"]
 * Expects Return: [1, 0, 1]
 */
func TestHackerRankTestInputEleven(t *testing.T) {
	inputStrings := []string{"def", "de", "fgh"}
	inputQueries := []string{"de", "lmn", "fgh"}
	expectedResult := []int32{1, 0, 1}
	result := matchingStrings(inputStrings, inputQueries)

	if !reflect.DeepEqual(result, expectedResult) {
		t.Errorf(`Result = %v, Expected Result = %v`, result, expectedResult)
	}
}

/*
*
  - Test Case: Hacker Rank Test Input Twelve
  - Input: ["abcde", "sdaklfj", "asdjf", "na", "basdn", "sdaklfj", "asdjf", "na", "asdjf", "na", "basdn", "sdaklfj", "asdjf"],
    ["abcde", "sdaklfj", "asdjf", "na", "basdn"]
  - Expects Return: [1, 3, 4, 3, 2]
*/
func TestHackerRankTestInputTwelve(t *testing.T) {
	inputStrings := []string{"abcde", "sdaklfj", "asdjf", "na", "basdn", "sdaklfj", "asdjf", "na", "asdjf", "na", "basdn", "sdaklfj", "asdjf"}
	inputQueries := []string{"abcde", "sdaklfj", "asdjf", "na", "basdn"}
	expectedResult := []int32{1, 3, 4, 3, 2}
	result := matchingStrings(inputStrings, inputQueries)

	if !reflect.DeepEqual(result, expectedResult) {
		t.Errorf(`Result = %v, Expected Result = %v`, result, expectedResult)
	}
}
