package main

import "fmt"

func main() {
	input := "geeksforgeeks"
	removed := removeDuplicates(input)
	fmt.Printf("Before: %s, After: %s\n", input, removed)
}

func removeDuplicates(s string) string {
	return recursiveRemove(s, "")
}

func recursiveRemove(total, result string) string {
	if len(total) == 0 {
		return result
	}
	if len(result) == 0 {
		return recursiveRemove(total[1:], string(total[0]))
	}

	consider, previous := total[0], result[len(result)-1]
	if consider == previous {
		return recursiveRemove(total[1:], result)
	}

	return recursiveRemove(total[1:], result+string(total[0]))
}
