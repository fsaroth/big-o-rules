package main

import "fmt"

var boxes = [5]int32{1, 2, 3, 4, 5}

func main() {
	logFirstTwoBoxes(boxes)
}

/*
logFirstTwoBoxes prints the first two elements of the given array of boxes.
Time complexity: O(1) + O(1) = O(2) => O(1)

Parameters:

	*- boxes: array of int32 elements

Return:

	*- void
*/
func logFirstTwoBoxes(boxes [5]int32) {
	fmt.Println(boxes[0]) // O(1)
	fmt.Println(boxes[1]) // O(1)
}
