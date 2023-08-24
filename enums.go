package main

import "fmt"

// enums

// weapon type
// axe
// sword
// wooden stick
// knife

type WeaponType int

func (w WeaponType) String() string {
	switch w {
	case Sword:
		return "SWORD"
	case Axe:
		return "AXE"
	case Knife:
		return "KNIFE"
	default:
		panic("not found")
	}
}

const (
	//Axe         WeaponType = 1
	//Sword       WeaponType = 2
	//WoodenStick WeaponType = 3
	//Knife       WeaponType = 4
	Axe WeaponType = iota // auto increment
	Sword
	WoodenStick
	Knife
)

func getDamage(weaponType WeaponType) int {
	switch weaponType {
	case Axe:
		return 100
	case Sword:
		return 90
	case WoodenStick:
		return 1
	case Knife:
		return 40
	default:
		panic("weapon does not exist")
	}
}

func main() {
	fmt.Printf("damage of the given weapon (%s) (%d):\n", Axe, getDamage(Axe))
	fmt.Printf("damage of the given weapon (%s) (%d):\n", Sword, getDamage(Sword))
	fmt.Printf("damage of the given weapon (%s) (%d):\n", Knife, getDamage(Knife))
}
