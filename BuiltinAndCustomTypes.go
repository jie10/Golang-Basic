package main

// Structs
/*

// general types
var {
	floatVar32 float32 = 0.1
	floatVar64 float64 = 0.1
	name string = "Foo"
	intVar32 int32 = 1
	intVar64 int64 = 123123
	intVar int = -1
	uintVar uint = 1
	uint32Var uint32 = 1
	uint64Var uint64 = 10
	uint8Var uint8 = 0x1
	byteVar byte = 0x2
	runVar rune = 'a'
}
type Player struct {
	name        string
	health      int
	attackPower float64
}

func (player Player) getHealth() int {
	return player.health
}

// optional
func getHealth(player Player) int {
	return player.health
}

func main() {
	player := Player{
		name:        "Captain Jack",
		health:      101,
		attackPower: 45.1,
	}
	//fmt.Printf("this is the player: %+v\n", player)
	//fmt.Printf("this is the player health: %d\n", getHealth(player))
	fmt.Printf("health: %v\n", player.health)
	fmt.Printf("this is the player health: %d\n", player.getHealth())
}
*/

// maps
/*
func main() {
	//users := map[string]int{} // empty map

	users := make(map[string]int)

	users["foo"] = 10
	users["bar"] = 11
	//fmt.Printf("users: %+v\n", users)

	// for
	for key, value := range users {
		fmt.Printf("the key: %s and the value: %d\n", key, value)
	}

	//delete users
	delete(users, "foo")
	fmt.Print(users)

	//if condition
	age, ok := users["zoo"]
	if !ok {
		fmt.Println("user not found")
	} else {
		fmt.Println("user found: ", age)
	}
	//fmt.Printf("age: %d\n", age)
}
*/

// slices
/*
func main() {
	//numbers := [2]int{}
	numbers := []int{}
	//otherNumbers := make([]int, 3)
	//fmt.Println(otherNumbers)

	numbers = append(numbers, 1)
	numbers = append(numbers, 10)
	fmt.Println(numbers)
}
*/

// custom types
/*
type Weapon string

func getWeapon(weapon Weapon) string {
	return string(weapon)
}
*/
