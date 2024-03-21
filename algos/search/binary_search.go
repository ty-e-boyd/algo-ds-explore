package main

import "github.com/charmbracelet/log"

func main() {
	set1 := []int{1, 4, 8, 10, 23, 34, 42}
	set2 := []int{4, 12, 18, 34, 39, 66}

	// set, target combo should successfully find no target
	i, nf := BinarySearch(set1, 2)
	if nf {
		log.Warn("Set did not find target")
	} else {
		log.Warnf("Target found in set at index %v", i)
	}

	// set, target combo should successfully find a target
	i, nf = BinarySearch(set2, 18)
	if nf {
		log.Warn("Set did not find target")
	} else {
		log.Warnf("Target found in set at index %v", i)
	}
}

/*
BinarySearch takes a sorted data set and a target, and returns an int and a bool.
If the bool is true, regardless of int, target was not found.
If bool is false, the int is the index where the target was within the set

BigO -> O(log n)
*/
func BinarySearch(s []int, t int) (int, bool) {
	// TODO Implement BinarySearch

	// target not found
	return 0, true
}
