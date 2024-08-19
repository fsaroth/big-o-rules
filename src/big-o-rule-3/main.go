package main

import "fmt"

/*
Big O Rule 3: Different terms for inputs

O(a + b) => O(a) + O(b)

compareTwoBoxes prints the contents of two boxes.
It iterates over each element in boxes1 and prints it.
Then, it iterates over each element in boxes2 and prints it.

Time complexity: O(a + b)
Space complexity: O(1)

Parameters:

	*- boxes1: a slice of integers
	*- boxes2: a slice of integers

Return:

	*- none

Example:

	compareTwoBoxes([]int{1, 2, 3, 4, 5}, []int{6, 7, 8, 9, 10})
*/
func compareTwoBoxes(boxes1, boxes2 []int) {
	for _, box1 := range boxes1 {
		fmt.Println(box1)
	}
	for _, box2 := range boxes2 {
		fmt.Println(box2)
	}
}

func main() {
	boxes1 := []int{1, 2, 3, 4, 5}
	boxes2 := []int{6, 7, 8, 9, 10}

	compareTwoBoxes(boxes1, boxes2)
}
