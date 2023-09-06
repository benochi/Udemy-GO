package main

//map keys must all be the same type. values must all be the same type.
//Use a map over a struct when properties(values) are closely related
import "fmt"

func main() {
	//ways to create a map
	// var colors map[string]string

	//colors := make(map[string]string)

	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
		"white": "#ffffff",
	}

	//add values to map
	// colors["white"] = "#ffffff"

	//delete off a map
	// delete(colors, "white")

	//iterate over a map
	printMap(colors)
}

func printMap(c map[string]string) {
	for key, value := range c {
		fmt.Println("Hexcode for", key, "is", value)
	}
}
