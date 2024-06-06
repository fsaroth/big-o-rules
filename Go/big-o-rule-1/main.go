package main

var nemo = []string{"nemo"}
var everyone = []string{
	"dory", "bruce", "marlin", "nemo", "gill",
	"bloat", "nigel", "squirt", "darla", "hank",
}
var large = make([]string, 1000000)

/*
findNemo function is responsible for finding the value "nemo" in the slice.
The time complexity of this function is O(n).

Arguments:

	*- slice []string: slice of strings to be searched for the value "nemo".

Returns:

	*- void
*/
func findingNemo(slice []string) {
	for i := 0; i < len(slice); i++ {
		if slice[i] == "nemo" {
			println("Found Nemo!")
			// break
			/*
				Even if we have the break statement,
				the time complexity of this function is still O(n).
				(Worst Case)
			*/
		}
	}
}

/*
fillSliceWithString fills the given slice with the specified value.
The time complexity of this function is O(n).

Arguments:

	*- slice []string: slice of strings to be filled with the value.
	*- value string: value to be filled in the slice.

Returns:

	*- void
*/
func fillSliceWithString(slice []string, value string) {
	for i := 0; i < len(slice); i++ {
		slice[i] = value
	}
}

func main() {
	fillSliceWithString(large, "nemo")
	findingNemo(large)
}
