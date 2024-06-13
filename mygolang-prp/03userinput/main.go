package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "welcome to user input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter any number: ")

	// comma ok || err err

	input, _ := reader.ReadString('\n')
	fmt.Println(input)
	fmt.Printf("type of input %T", input)
	// fmt.Println("\n")

}
