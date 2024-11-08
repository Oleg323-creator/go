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

const hp uint64 = 100

func (h *Human) Gethp() uint64 {
	if h.CritChance() == 0 {
		resultHp := hp - h.Damage()
		return resultHp
	} else if uint64(rand.Intn(100)) <= h.CritChance() {
		resultHp := hp - (h.Damage() * 2)
		return resultHp
	} else {
		resultHp := hp - h.Damage()
		return resultHp
	}
}

func (h *Human) Attack(enemy *Human) {
	if h.Gethp() <= 0 {
		fmt.Println("enemy died")
	} else {
		fmt.Println("enemy's hp : ", h.Gethp())
		if enemy.Gethp() <= 0 {
			fmt.Println("you died")
		} else {
			fmt.Println("your's hp : ", enemy.Gethp())

		}
	}
}
