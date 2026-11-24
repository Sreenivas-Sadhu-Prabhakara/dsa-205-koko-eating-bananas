package main

import "fmt"

// MinEatingSpeed returns the smallest eating speed k (bananas per hour)
// that finishes all piles within h hours, eating from one pile per hour.
// Binary search on the answer: hours needed is monotone decreasing in k.
func MinEatingSpeed(piles []int, h int) int {
	lo, hi := 1, 1
	for _, p := range piles {
		if p > hi {
			hi = p // eating faster than the largest pile gains nothing
		}
	}
	for lo < hi {
		k := lo + (hi-lo)/2
		if hoursNeeded(piles, k) <= h {
			hi = k
		} else {
			lo = k + 1
		}
	}
	return lo
}

func hoursNeeded(piles []int, k int) int {
	hours := 0
	for _, p := range piles {
		hours += (p + k - 1) / k // ceil division: a partial pile still costs the hour
	}
	return hours
}

func main() {
	fmt.Println(MinEatingSpeed([]int{3, 6, 7, 11}, 8))
	fmt.Println(MinEatingSpeed([]int{30, 11, 23, 4, 20}, 5))
}
