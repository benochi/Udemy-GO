package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for _, value := range nums {
		if value%2 == 0 {
			fmt.Println(value, "is EVEN!")
		} else {
			fmt.Println(value, "is ODD!")
		}
	}
}
