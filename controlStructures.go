package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7}

	for i := 0; i < len(numbers); i++ {
		fmt.Println(numbers[i])
	}
}
