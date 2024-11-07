package main

import "fmt"

/*
	Space Complexity: How much more memory is required by the algorithm as the input size grows.

	1. Most primitives (booleans, numbers, undefined, null) are constant space.
	2. Strings require O(n) space (where n is the string length).
	3. Reference types are generally O(n), where n is the length (for arrays) or the number of keys (for objects).
	4. Space complexity in a recursive function is O(n), where n is the depth of the recursive call.
	5. Logarithmic space complexity is generally O(log(n)). An example of this is the binary search algorithm.

*/

/*
booo function prints "booo!" n times.
Time complexity: O(n)
Space complexity: O(1)

Parameters:

	n: A slice of integers to print "booo!" n times

Returns:

	nil

Example:

	booo([]int{1, 2, 3, 4, 5})
*/
func booo(n []int) {
	for i := 0; i < len(n); i++ {
		fmt.Println("booo!")
	}
}

/*
array function creates an array of "hi" n times.
Time complexity: O(n)
Space complexity: O(n)

Parameters:

	n: An integer to create an array of "hi" n times

Returns:

	[]string: An array of "hi" n times

Example:

	arrayOfHiNTimes(6)
*/

func arrayOfHiNTimes(n int) []string {
	var hiArr []string
	for i := 0; i < n; i++ {
		hiArr = append(hiArr, "hi")
	}
	return hiArr
}

func main() {
	booo([]int{1, 2, 3, 4, 5})
	fmt.Println(arrayOfHiNTimes(6))
}
