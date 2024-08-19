package main

/*
logAllPairsOfArray prints all pairs of elements in the given array.
It iterates over the array and for each element, it iterates again to print the pair.

Time complexity: O(n^2) =>
Quadratic time, the time it takes to run the function grows quadratically with the size of the input data.

Space complexity: O(1)

Parameters:

array: a slice of integers

Return:

none

Example:

logAllPairsOfArray([]int{1, 2, 3, 4, 5})
*/
func logAllPairsOfArray(array []int) {
	for i := 0; i < len(array); i++ {
		for j := 0; j < len(array); j++ {
			println(array[i], array[j])
		}
	}
}

func main() {
	array := []int{1, 2, 3, 4, 5}

	logAllPairsOfArray(array)
}

/*
O(n * n) = O(n^2) - Quadratic time,
the time it takes to run the function grows quadratically with the size of the input data.
*/
