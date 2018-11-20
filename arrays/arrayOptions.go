package arrays

import "fmt"

func ListOfOptionsForArrays() {
	fmt.Println(" Lets explore Arrays in go language")
	fmt.Println(" Enter 1. to  Find all elements that appear more than n/k times in an array")
	fmt.Println(" Enter 2. to  Count minimum number of subsets (or subsequences) with consecutive numbers")
	fmt.Println("Enter your choice : ")
	var input int
	_, err := fmt.Scanf("%d", &input)
	if err == nil {
		OptionsOfArrys(input)
	}
}
