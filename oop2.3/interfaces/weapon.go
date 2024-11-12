package interfaces

import (
	"fmt"
	"log"
	"math/rand"
)

type Weapon interface {
	GetDamage() uint64
	GetCritChance() uint64
}

// ADDING Sniper TO Weapon INTERFACE

type Sniper struct {
	Damage     uint64
	CritChance uint64
}

func (s *Sniper) GetDamage() uint64 {
	return s.Damage
}

func (s *Sniper) GetCritChance() uint64 {
	return s.CritChance
}

// ADDING Guns TO Weapon INTERFACE

type Gun struct {
	Damage uint64
}

func (g *Gun) GetDamage() uint64 {
	return g.Damage
}

func (g *Gun) GetCritChance() uint64 {
	return 0
}

// Added Guns as field
type Human struct {
	Hp   uint64
	Guns Weapon
}

// Generating humans
func NewHuman() *Human {
	return &Human{
		Hp:   100,
		Guns: nil,
	}
}

// Setting weapon
func (h *Human) SetWeapon(TypeWeapon Weapon) {
	if TypeWeapon == nil {
		h.Guns = TypeWeapon
		log.Println("Can't shoot without gun")
	} else {
		h.Guns = TypeWeapon
	}

}

func (h *Human) GetHp() uint64 {
	return h.Hp
}

func (h *Human) Attack(enemy *Human) {
	// counting crit chance
	if h.Guns == nil {
		log.Println("Can't shoot without gun")
		return
	}

	dmg := h.Guns.GetDamage()
	if h.Guns.GetCritChance() == 0 {
		h.Hp -= dmg
	} else if uint64(rand.Intn(100)) <= h.Guns.GetCritChance() {
		h.Hp -= dmg * 2
	} else {
		h.Hp -= dmg
	}

	if h.Hp <= 0 {
		fmt.Println("enemy died")
	} else {
		fmt.Println("enemy's hp : ", h.Hp)
		if enemy.Hp <= 0 {
			fmt.Println("you died")
		} else {
			fmt.Println("your's hp : ", enemy.Hp)

		}
	}

	//	return ТЫ СКАЗАЛ РЕТУРНИТЬ ИЗ ФУНКЦИИ СНАЧАЛА, А ПОТОМ СКАЗАЛ СНАЧАЛА СТРЕЛЯТЬ В ОДНОГО
	//	ПОТОМ ВЫБРОСИТЬ И СТРЕЛЬНУТЬ В ДРУГОГО Я ТАК ПОНЯЛ УЖЕ НЕ НАДО ЗДЕСЬ РЕТУРН

}
