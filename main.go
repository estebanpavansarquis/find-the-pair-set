package main

import (
	"fmt"
)

func main() {

	// FindThePairSet function call
	findThePairSetInput := GetFindThePairSetInput()
	fmt.Println("My input for FindThePairSet is: ", findThePairSetInput)
	fmt.Println("Result of running FindThePairSet: ", FindThePairSet(findThePairSetInput.NumberList, findThePairSetInput.Target))

}
