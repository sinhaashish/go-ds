package arrays

import (
	"fmt"
	"sort"
)

// It gives the options to choose form different problems in arrys.
func OptionsOfArrys(choice int) {
	switch choice {
	case 1:
		ElementMoreThanNKTimes()
	case 2:
		CountMinimumNumberOfSubsets()
	default:
		fmt.Println(" Your choice does not matches the given options .Please try again")
	}

}

// Find all elements that appear more than n/k times in an array
func ElementMoreThanNKTimes() {
	inputArray := []int{3, 1, 2, 2, 1, 2, 3, 3}
	fmt.Println(" The input array is ", inputArray)

}

// Count minimum number of subsets (or subsequences) with consecutive numbers
func CountMinimumNumberOfSubsets() {
	inputArray := []int{100, 56, 5, 6, 102, 58, 101, 57, 7, 103, 59}
	sort.Ints(inputArray)
	count := 1
	for i := 0; i < len(inputArray)-1; i++ {
		if inputArray[i]+1 != inputArray[i+1] {
			count++
		}
	}
	fmt.Print(" \n The input array is ", inputArray)
	fmt.Println(" \n The input array is ", count)
}
