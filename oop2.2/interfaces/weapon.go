package interfaces

import (
	"fmt"
	"math/rand"
)

type Weapon interface {
	Damage() uint64
	CritChance() uint64
}

type Sniper struct {
	damage     uint64
	critChance uint64
}

func (s *Sniper) Damage() uint64 {
	return s.damage
}

func (s *Sniper) CritChance() uint64 {
	return s.critChance
}

type Gun struct {
	damage uint64
}

func (g *Gun) Damage() uint64 {
	return g.damage
}

func (g *Gun) CritChance() uint64 {
	return 0
}

type Human struct {
	hp uint64
	Weapon
}

func (h *Human) GetHp() uint64 {
	return h.hp
}

func (h *Human) Attack(enemy *Human) {
	// counting crit chance
	dmg := h.Damage()
	if h.CritChance() == 0 {
		h.hp -= dmg
	} else if uint64(rand.Intn(100)) <= h.CritChance() {
		h.hp -= dmg * 2
	} else {
		h.hp -= dmg
	}
	//
	if h.hp <= 0 {
		fmt.Println("enemy died")
	} else {
		fmt.Println("enemy's hp : ", h.hp)
		if enemy.hp <= 0 {
			fmt.Println("you died")
		} else {
			fmt.Println("your's hp : ", enemy.hp)

		}
	}
}
