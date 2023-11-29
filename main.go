package main

import (
	"fmt"
)

// call is a function that takes a variable of any type and returns it.
func call(variable interface{}) interface{} {
	return variable
}

func main() {
	// Example usage of the call function with different types of variables
	number := 42
	text := "Hello, World"
	floatNum := 3.14

	// Call the function with different types of arguments
	resultNumber := call(number)
	resultText := call(text)
	resultFloat := call(floatNum)

	fmt.Println("Number:", resultNumber)
	fmt.Println("Text:", resultText)
	fmt.Println("Float Number:", resultFloat)
}

// func main() {
// 	seqMetadata := metadata.NewMetadata("")

// 	collectibleInfo, err := seqMetadata.GetTokenMetadata(context.Background(), "polygon", "0x631998e91476DA5B870D741192fc5Cbc55F5a52E", []string{"1", "2"})
// 	if err != nil {
// 		fmt.Println("Error fetching token metadata:", err)
// 		return
// 	}

// 	fmt.Println("Collectible Information:", collectibleInfo)
// }
