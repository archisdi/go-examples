package main

import "fmt"

func main() {
	// var colors map[string]string

	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#00FF00",
		"white": "#FFFFFF",
		"black": "#000000",
	}

	// colors := make(map[string]string)
	// colors["red"] = "#ff0000"
	// colors["green"] = "#00FF00"

	// delete(colors, "red")

	// fmt.Println(colors)
	printMap(colors)
}

func printMap(mapData map[string]string) {
	for color, hex := range mapData {
		fmt.Println("hex code for", color, "is", hex)
	}
}
