package main

import (
	"math/rand"
)

func main() {
	awp := Sniper{50, 50}
	scout := Sniper{25, 10}
	ak47 := Gun{30, 0}
	m16 := Gun{25, 0}

	guns := []Weapon{&awp, &scout, &ak47, &m16}
	for i := 0; i < 5; i++ {
		human := Human{hp, guns[rand.Intn(len(guns))]}
		enemy := Human{hp, guns[rand.Intn(len(guns))]}
		human.Attack(&enemy)
	}
}
