package main

import  "oop2/interfaces"

func main() {
	awp := AWP{50, 50}
	scout := Scout{25, 10}
	ak47 := AK47{30, 0}
	m16 := M16{25, 0}

	guns := []Weapon{&awp, &scout, &ak47, &m16}
	for i := range guns {
		human := Human{hp, guns[i]}
		human.Attack()
	}

}
