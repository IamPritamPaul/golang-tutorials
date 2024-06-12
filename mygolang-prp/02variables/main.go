package main

import (
	"fmt"
)

func main() {
	// fmt.Println("variables\n")
	var username string = "iampritampaul"
	fmt.Println(username)
	fmt.Printf("type of variable: %T\n\n", username)

	var isLoggedin bool = true
	fmt.Println(isLoggedin)
	fmt.Printf("type of variable: %T\n\n", isLoggedin)

	var smallValue uint8 = 255
	fmt.Println(smallValue)
	fmt.Printf("type of variable: %T\n\n", smallValue)

	var smallFloat float32 = 255.65237547263548263854
	fmt.Println(smallFloat)
	fmt.Printf("type of variable: %T\n\n", smallFloat)

	var bigFloat float64 = 255.65237547263548263854
	fmt.Println(bigFloat)
	fmt.Printf("type of variable: %T\n\n", bigFloat)

	//default values and some aliases
	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("type of variable: %T\n\n", anotherVariable)

	//implicit type
	var website = "https://google.com"
	fmt.Println(website)
	//website = 34
	// will give error

	//no var style
	numberofUsers := 23 // not allowed in global scope
	fmt.Println(numberofUsers)

}
