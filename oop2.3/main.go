package main

import (
	"fmt"
	"math/rand"
	"oop2/interfaces"
)

func main() {
	awp := interfaces.Sniper{Damage: uint64(50), CritChance: uint64(50)}
	scout := interfaces.Sniper{Damage: 25, CritChance: 10}
	ak47 := interfaces.Gun{Damage: 30}
	m16 := interfaces.Gun{Damage: 25}

	guns := []interfaces.Weapon{&awp, &scout, &ak47, &m16}
	for i := 0; i < 1; i++ {

		// HUMANS INITIALIZATION
		human := interfaces.NewHuman()
		enemy := interfaces.NewHuman()
		enemy1 := interfaces.NewHuman()

		human.SetWeapon(guns[rand.Intn(len(guns))])
		fmt.Println(human.Guns)
		human.Attack(enemy)
		fmt.Println("ENEMY'S1 HP", human.GetHp()) // ENEMY'S HP

		human.SetWeapon(nil)
		fmt.Println(human.Guns)

		human.SetWeapon(guns[rand.Intn(len(guns))])
		fmt.Println(human.Guns)
		human.Attack(enemy1)
		fmt.Println("ENEMY'S1 HP", human.GetHp()) // ENEMY'S1 HP

	}
}
