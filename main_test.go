package main

import (
	"reflect"
	"testing"
)

func TestMinCoins(t *testing.T) {
	cases := []struct {
		amount int
		want   []int
	}{
		{49, []int{20, 20, 5, 2, 2}},
		{121, []int{100, 20, 1}},
		{0, []int{}},
		{1, []int{1}},
		{2896, []int{2000, 500, 200, 100, 50, 20, 20, 5, 1}},
	}
	for _, c := range cases {
		got := MinCoins(c.amount)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("MinCoins(%d) = %v, want %v", c.amount, got, c.want)
		}
		sum := 0
		for _, d := range got {
			sum += d
		}
		if sum != c.amount {
			t.Errorf("MinCoins(%d) sums to %d", c.amount, sum)
		}
	}
}
