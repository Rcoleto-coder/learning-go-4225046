package main

import (
	"fmt"
)

func main() {

	str1 := "The quick red fox"
	str2 := "jumped over"
	str3 := "the lazy brown dog."
	aNumber := 42
	stringLength, err := fmt.Println("The value is", aNumber)

	if err == nil {
		fmt.Println("String length:", stringLength)
	}

	fmt.Println(str1, str2, str3)
	fmt.Printf("Value of number: %v\n", aNumber)
	fmt.Printf("Data type: %T\n", aNumber)
}
