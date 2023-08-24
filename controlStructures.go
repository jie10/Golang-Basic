package main

import "fmt"

func main() {
	//numbers := []int{1, 2, 3, 4, 5, 6, 7}
	//

	// for loops
	//for i := 0; i < len(numbers); i++ {
	//	fmt.Println(numbers[i])
	//}

	//names := []string{"Foo", "Bar", "Baz"}
	//
	// // range
	//for key, value := range names {
	//	fmt.Printf("the key is %d and the value is %s\n", key, value)
	//}

	//
	//names := []string{"a", "b", "c", "d"}
	//
	//for _, name := range names {
	//	if name == "a" {
	//		return
	//		//break
	//	}
	//}
	//
	//fmt.Println("break out of loop")

	users := map[string]int{
		"Foo": 1,
		"Bar": 2,
		"Baz": 3,
	}

	for key, value := range users {
		fmt.Printf("the key is %s and the value is %d\n", key, value)
	}

	name := "Foo"

	switch name {
	case "Bar":
		fmt.Println("The name is Bar")
	case "Baz":
		fmt.Println("The name is Baz")
	default:
		fmt.Println("The name is default:", name)
	}
}
