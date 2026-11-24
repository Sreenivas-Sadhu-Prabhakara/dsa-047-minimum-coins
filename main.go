package main

import "fmt"

// denominations of Indian currency, largest first.
var denominations = []int{2000, 500, 200, 100, 50, 20, 10, 5, 2, 1}

// MinCoins returns the fewest notes/coins making up amount, greedily taking
// the largest denomination that fits. Valid because this denomination
// system is canonical: every denomination divides cleanly into the ones
// above it in a way that keeps greed optimal.
func MinCoins(amount int) []int {
	out := []int{}
	for _, d := range denominations {
		for amount >= d {
			amount -= d
			out = append(out, d)
		}
	}
	return out
}

func main() {
	fmt.Println(MinCoins(49))
	fmt.Println(MinCoins(2896))
}
