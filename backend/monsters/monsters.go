package monsters

import (
	"fmt"
	"math/rand"
);


type Monster struct {
	Name string
	HP int
	Attack int
	AttackBonus int
	Image string // This will be the path to the image
}


func (m *Monster) Print() {
	// For now, just print the name
	// We'll want to go ahead and eventually print out the other stats too
	fmt.Printf("Name: %s\nHP: %d\nAttack: %d\nImage Path: '%v'\n", m.Name, m.HP, m.Attack, m.Image);
}

func (m* Monster) ToString() string {
	return fmt.Sprintf("Name: %s, HP: %d, Attack: %d, AttackBonus: %d, ImagePath: '%s'", m.Name, m.HP, m.Attack, m.AttackBonus, m.Image);
}




type Monsters struct {
	Monsters []*Monster
}

func GenerateMonsters() *Monsters {
	// CHANGE THIS IF THE ROOT FOLDER CHANGES
	root_folder := "/src/assets/images/monsters/%v";
	
	monsters_array := []*Monster{
		{
			Name: "Slime",
			HP: 10,
			Attack: 2,
			AttackBonus: 2,
			Image: fmt.Sprintf(root_folder, "Slime.png"),
		},
		{
			Name: "Goobo",
			HP: 10,
			Attack: 2,
			AttackBonus: 2,
			Image: fmt.Sprintf(root_folder, "Gooblin.png"),
		},
		{
			Name: "Manager",
			HP: 10,
			Attack: 2,
			AttackBonus: 2,
			Image: fmt.Sprintf(root_folder, "Manager.png"),
		},
		{
			Name: "Sneezle",
			HP: 10,
			Attack: 2,
			AttackBonus: 2,
			Image: fmt.Sprintf(root_folder, "Sneezle.png"),
		},
		{
			Name: "Loan",
			HP: 10,
			Attack: 2,
			AttackBonus: 2,
			Image: fmt.Sprintf(root_folder, "Loan.png"),
		},
		{
			Name: "Meep",
			HP: 10,
			Attack: 2,
			AttackBonus: 2,
			Image: fmt.Sprintf(root_folder, "Meep.png"),
		},
		{
			Name: "Skeleton",
			HP: 10,
			Attack: 2,
			AttackBonus: 2,
			Image: fmt.Sprintf(root_folder, "Skeleton.png"),
		},
		{
			Name: "Pinchington",
			HP: 10,
			Attack: 2,
			AttackBonus: 2,
			Image: fmt.Sprintf(root_folder, "Pinchington.png"),
		},
		{
			Name: "Squiffer",
			HP: 10,
			Attack: 2,
			AttackBonus: 2,
			Image: fmt.Sprintf(root_folder, "Squiffer.png"),
		},
	};

	return &Monsters{
		Monsters: monsters_array,
	};
}


func NewMonster(possible_monsters *Monsters) *Monster {
	// Choose a random monster from the Monsters struct
	monster_types := possible_monsters.Monsters;
	monster := monster_types[rand.Intn(len(monster_types))];
	return monster;
}