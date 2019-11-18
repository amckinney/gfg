package main

import "fmt"

func main() {
	var (
		coins = []int{50, 25, 10, 5, 1}
		value = 100
	)
	fmt.Printf("I can make change %d ways for %d with %v!\n", countChange(value, coins), value, coins)
}

// countChange returns the number of ways you can
// make change for the given value using the given
// coins.
func countChange(value int, coins []int) int {
	return countChangeRecursive(value, coins, make(map[string]int))
}

// countChangeRecursive is a memoized, recursive solution
// for countChange.
func countChangeRecursive(value int, coins []int, memo map[string]int) int {
	if len(coins) == 0 || value < 0 {
		return 0
	}
	if value == 0 {
		return 1
	}
	if v, ok := memo[key(value, coins)]; ok {
		return v
	}

	// Select the largest coin.
	coin := coins[0]

	// Calculate how many ways if you use the coin.
	use := countChangeRecursive(value-coin, coins, memo)

	// Cache the result so that we avoid redundant work.
	memo[key(value-coin, coins)] = use

	// Then, calculate how many ways if you don't use the coin.
	skip := countChangeRecursive(value, coins[1:], memo)

	return use + skip
}

// key returns the map key for a specific value, coins pair.
func key(value int, coins []int) string {
	return fmt.Sprintf("%d,%v", value, coins)
}
