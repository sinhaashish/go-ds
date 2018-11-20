package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/go-ds/arrays"
)

func main() {

	fmt.Println(" Welcome to the world of Go DataStructure")
	fmt.Println(" Enter A for Arrays")
	fmt.Println(" Enter B for LinkedList")
	fmt.Println(" Enter C for Stack")
	fmt.Println(" Enter D for Queue")
	fmt.Println(" Enter E for Trees")
	fmt.Println(" Enter F for Graph")
	fmt.Println("Enter your choice : ")
	reader := bufio.NewReader(os.Stdin)
	char, _, _ := reader.ReadRune()
	switch char {
	case 'A':
		arrays.ListOfOptionsForArrays()
	case 'B':
		fmt.Println(" Lets explore LinkedList in go language ")
	case 'C':
		fmt.Println(" Lets explore Stack in go language")
	case 'D':
		fmt.Println(" Lets explore Queue in go language")
	case 'E':
		fmt.Println(" Lets explore Trees in go language")
	case 'F':
		fmt.Println(" Lets explore Graph in go language")
	default:
		fmt.Println(" Your choice does not matches the given options .Please try again")
	}
}
