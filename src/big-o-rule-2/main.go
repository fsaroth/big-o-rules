package main

import "fmt"

const SCALING_FACTOR = 1000000

/*
Big O Rule 2: Remove Constants

O(2n) => O(n)
O(n + 10) => O(n)

printFirstItemThenFirstHalfThenSayHiXTimes prints the first item of the given numbers,
followed by printing the first half of the numbers, and then saying "hi" by SCALING_FACTOR times.

Parameters:

	*- numbers: a slice of integers.

Returns:

	*- None.

Time Complexity:

	*- O(n/2 + 1 + SCALING_FACTOR) = O(n)

Space Complexity:

	*- O(1)
*/
func printFirstItemThenFirstHalfThenSayHiXTimes(numbers []int) {
	fmt.Println(numbers[0])
	middleIndex := len(numbers) / 2
	for i := 0; i < middleIndex; i++ {
		fmt.Println(numbers[i])
	}

	for i := 0; i < SCALING_FACTOR; i++ {
		fmt.Println("hi")
	}
}

func main() {
	numbers := []int{1, 2, 3, 4}
	printFirstItemThenFirstHalfThenSayHiXTimes(numbers)
}
