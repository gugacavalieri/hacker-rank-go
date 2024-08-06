// This is the main package for this module
package main

/*
 * Complete the 'matchingStrings' function below.
 *
 * The function is expected to return an INTEGER_ARRAY.
 * The function accepts following parameters:
 *  1. STRING_ARRAY strings
 *  2. STRING_ARRAY queries
 *
 * Time complexity: O(mn)
 */
func matchingStrings(strings []string, queries []string) []int32 {
	queryCounter := []int32{}
	queryCounterMap := map[string]int32{}

	// count queries
	for _, query := range queries {

		// check if map already contains key
		// that means query already has a count
		_, mapContainsKey := queryCounterMap[query]
		if mapContainsKey {
			continue
		}

		// if not, let's calculate the count for the current query
		currentCount := int32(0)
		for _, currentString := range strings {
			if currentString == query {
				currentCount++
			}
		}
		queryCounterMap[query] = currentCount
	}

	// build return array using counterMap
	for _, query := range queries {
		queryCounter = append(queryCounter, queryCounterMap[query])
	}

	return queryCounter
}
