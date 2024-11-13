package main

import (
	"fmt"
	"math/rand"
	"oop2/interfaces"
)

func main() {
	awp := interfaces.Sniper{Damage: uint64(50), CritChance: uint64(50), Bullet: 1}
	scout := interfaces.Sniper{Damage: 25, CritChance: 10, Bullet: 1}
	ak47 := interfaces.Gun{Damage: 30, Bullet: 3}
	m16 := interfaces.Gun{Damage: 25, Bullet: 5}

	guns := []interfaces.Weapon{&awp, &scout, &ak47, &m16}
	for i := 0; i < 1; i++ {

		// HUMANS INITIALIZATION
		human := interfaces.NewHuman()
		enemy := interfaces.NewHuman()
		enemy1 := interfaces.NewHuman()

		human.SetWeapon(guns[rand.Intn(len(guns))])
		human.Attack(enemy)
		fmt.Println("ENEMY'S HP", enemy.GetHp()) // ENEMY'S HP

		human.SetWeapon(nil)

		human.SetWeapon(guns[rand.Intn(len(guns))])
		human.Attack(enemy1)
		fmt.Println("ENEMY'S1 HP", enemy1.GetHp()) // ENEMY'S1 HP

	}
}
