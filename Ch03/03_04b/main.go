package main

import (
	"fmt"
	"sort"
)

func main() {
	// This is an array
	// var colors = [3]string{"Red", "Green", "Blue"}
	// This is a slice,like vectors but in GO it is a portion of the array that can be manipulated at runtime
	// var sliceOfColors = []string{"Red", "Green", "Blue"}

	var colors = make([]string, 0, 3)
	colors = append(colors, "Red", "Green", "Blue")
	fmt.Println(colors)
	colors = append(colors, "Purple", "Yellow")
	fmt.Println(colors)

	colors = remove(colors, 0)
	fmt.Println(colors)

	sort.Strings(colors)
	fmt.Println(colors)
}

// To remove itens from a slice
func remove(slice []string, i int) [] string {
	return append(slice[:i], slice[i+1:]...)
}
