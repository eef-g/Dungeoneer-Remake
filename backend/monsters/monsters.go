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
	// Print the monster
	// For now, just print the name
	fmt.Println(m.Name);
}

func (m* Monster) ToString() string {
	return fmt.Sprintf("Name: %s, HP: %d, Attack: %d, AttackBonus: %d, ImagePath: %s", m.Name, m.HP, m.Attack, m.AttackBonus, m.Image);
}




type Monsters struct {
	Monsters []*Monster
}

func GenerateMonsters() *Monsters {
	monsters_array := []*Monster{
		&Monster{
			Name: "Slime",
			HP: 10,
			Attack: 2,
			AttackBonus: 2,
			Image: "slime.png",
		},
		&Monster{
			Name: "Goblin",
			HP: 10,
			Attack: 2,
			AttackBonus: 2,
			Image: "goblin.png",
		},
		&Monster{
			Name: "Orc",
			HP: 10,
			Attack: 2,
			AttackBonus: 2,
			Image: "orc.png",
		},
		&Monster{
			Name: "Ogre",
			HP: 10,
			Attack: 2,
			AttackBonus: 2,
			Image: "ogre.png",
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