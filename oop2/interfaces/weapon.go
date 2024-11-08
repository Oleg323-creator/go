package interfaces

import (
	"fmt"
	"math/rand"
)

type Weapon interface {
	Damage() uint64
	CritChance() uint64
}

type AWP struct {
	damageA     uint64
	critChanceA uint64
}

func (a *AWP) Damage() uint64 {
	return a.damageA
}

func (a *AWP) CritChance() uint64 {
	return a.critChanceA
}

type Scout struct {
	damageS     uint64
	critChanceS uint64
}

func (s *Scout) Damage() uint64 {
	return s.damageS
}

func (s *Scout) CritChance() uint64 {
	return s.critChanceS
}

type AK47 struct {
	damageAK     uint64
	critChanceAK uint64
}

func (a *AK47) Damage() uint64 {
	return a.damageAK
}

func (a *AK47) CritChance() uint64 {
	return a.critChanceAK
}

type M16 struct {
	damageM     uint64
	critChanceM uint64
}

func (m *M16) Damage() uint64 {
	return m.damageM
}

func (m *M16) CritChance() uint64 {
	return m.critChanceM
}

type Human struct {
	hp uint64
	Weapon
}

const hp uint64 = 100

func (h *Human) Attack() {
	if h.CritChance() == 0 {
		resultHp := hp - h.Damage()
		if resultHp <= 0 {
			fmt.Println("enemy has dead")
		} else {
			fmt.Println(resultHp)
		}
	} else if uint64(rand.Intn(100)) <= h.CritChance() {
		resultHp := hp - (h.Damage() * 2)
		if resultHp <= 0 {
			fmt.Println("enemy has dead")
		} else {
			fmt.Println(resultHp)
		}
	} else {
		resultHp := hp - h.Damage()
		if resultHp <= 0 {
			fmt.Println("enemy has dead")
		} else {
			fmt.Println(resultHp)
		}
	}

}
