package main

import "fmt"

/*
Big o Rule 4: Drop Non Dominants

O(n^2 + 3n + 100 + n/2) => O(n^2)

printAllNumbersThenAllPairSums prints all the numbers in the given slice and
then prints the sum of each pair of numbers.

Time complexity: O(n^2)
Space complexity: O(K) where K = Constant

Parameters:

	numbers: A slice of integers to print and calculate the sum of each pair of

Returns:

	nil

Example:

	printAllNumbersThenAllPairSums([]int{1, 2, 3, 4, 5})
*/
func printAllNumbersThenAllPairSums(numbers []int) {
	fmt.Println("these are the numbers: ")
	for _, number := range numbers {
		fmt.Println(number)
	}

	fmt.Println("and these are their sums: ")
	for _, firstNumber := range numbers {
		for _, secondNumber := range numbers {
			fmt.Println(firstNumber + secondNumber)
		}
	}
}

func main() {
	printAllNumbersThenAllPairSums([]int{1, 2, 3, 4, 5})
}
