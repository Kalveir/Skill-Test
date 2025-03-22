package main

import (
	"fmt"
)

// Parent Struct Character
type Character struct {
	Name    string
	Damage  int
	Defense int
}

func (c Character) ShowStats() {
	fmt.Printf("%s - Damage: %d, Defense: %d\n", c.Name, c.Damage, c.Defense)
}

// Child Struct: Hero
type Hero struct {
	Character
	SpecialSkill string
}

func (h Hero) UseSkill() {
	fmt.Printf("%s menggunakan skill: %s!\n", h.Name, h.SpecialSkill)
}

// Child Struct: Villain
type Villain struct {
	Character
	EvilPlan string
}

func (v Villain) ExecutePlan() {
	fmt.Printf("%s menjalankan rencana: %s!\n", v.Name, v.EvilPlan)
}

func main() {
	hero := Hero{Character{"Ginzo", 10, 50}, "Hiken - Tsukikage"}
	villain := Villain{Character{"Muno Wonda", 30, 700}, "Menaklukan Ashura World"}

	hero.ShowStats()
	hero.UseSkill()

	fmt.Println()

	villain.ShowStats()
	villain.ExecutePlan()
}

