package main

import (
	"fmt"
)

func main() {
	colors := map[string]string{
		"red":   "#ff000",
		"green": "#4bf745",
		"white": "#ffffff",
	}

	printmp(colors)
}

func printmp(m map[string]string) {
	for color, hex := range m {
		fmt.Println("hex code for ", color, " is ", hex)
	}
}
