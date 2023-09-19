package monsters

import (
	"fmt"
);


type Monster struct {
	Name string
	HP int
	Attack int
	AttackBonus int
	Image string // This will be the path to the image
}

type Monsters struct {
	Monsters []*Monster
}


func (m *Monster) Print() {
	// Print the monster
	// For now, just print the name
	fmt.Println(m.Name);
}

func (m* Monster) ToString() string {
	return fmt.Sprintf("Name: %s, HP: %d, Attack: %d, AttackBonus: %d, ImagePath: %s", m.Name, m.HP, m.Attack, m.AttackBonus, m.Image);
}


func NewMonster(name string, hp int, attack int, attackBonus int, image string) *Monster {
	return &Monster{
		Name: name,
		HP: hp,
		Attack: attack,
		AttackBonus: attackBonus,
		Image: image,
	}
}

