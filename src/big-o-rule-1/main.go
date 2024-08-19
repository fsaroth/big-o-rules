package main

import "fmt"

/*
Big O Rule 1: Worst Case

The worst-case scenario is the scenario in which the algorithm takes the longest time to execute.
It is important to consider the worst-case scenario when analyzing the time complexity of an algorithm.

findNemo searches for the string "a" in the given slice of strings.
If "a" is found, it prints "Found Nemo!" to the console.

Time complexity: O(n)
Space complexity: O(1)

Parameters:

	nemo: a slice of strings

Return:

	nil

Example:

	findNemo([]string{"nemo", "dory", "bruce", "marlin", "gill", "bloat", "nigel", "squirt", "darla", "hank"})
*/
func findNemo(nemo []string) {
	for i := 0; i < len(nemo); i++ {
		if nemo[i] == "nemo" {
			fmt.Println("Found Nemo!")
		}
	}
}

func main() {
	nemo := []string{
		"nemo", "dory", "bruce", "marlin", "gill", "bloat", "nigel", "squirt", "darla", "hank",
	}
	findNemo(nemo)
	large := make([]string, 1000000)
	for i := 0; i < len(large); i++ {
		large[i] = "nemo"
	}
	findNemo(large)
}
