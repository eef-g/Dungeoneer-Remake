package entities

import (
	"fmt"
	"math/rand"
);


type Monster struct {
	Base Entity
  AttackBonus int // This is the value that scaeles with the player level 
	Image string // This will be the path to the image
}

func (m *Monster) Print() {
	// For now, just print the name
	// We'll want to go ahead and eventually print out the other stats too
	fmt.Printf("Name: %s\nHP: %d\nAttack: %d\nImage Path: '%v'\n", m.Base.Name, m.Base.Health, m.Base.Damage, m.Image);
}

func (m* Monster) ToString() string {
	return fmt.Sprintf("Name: %s, HP: %d, Attack: %d, AttackBonus: %d, ImagePath: '%s'", m.Base.Name, m.Base.Health, m.Base.Damage, m.AttackBonus, m.Image);
}

type Monsters struct {
	Monsters []*Monster
}


func CraftMonster(name string, health int, damage int, damage_bonus int, image_path string) *Monster{
	root_folder := "/src/assets/images/monsters/%v";
  monster := &Monster{
    Base: Entity{
      Name: name,
      Health: health,
      Damage: damage,
    },
    AttackBonus: damage_bonus,
    Image: fmt.Sprintf(root_folder, image_path),
  };
  return monster;
}

func GenerateMonsters() *Monsters {
  // Populate the monsters_array function by calling the CraftMonster function for each monster
	monsters_array := []*Monster{
    CraftMonster("Slime", 10, 2, 2, "Slime.png"),
    CraftMonster("Goobo", 10, 2, 2, "Gooblin.png"),
    CraftMonster("Manager", 10, 2, 2, "Manager.png"),
    CraftMonster("Sneezle", 10, 2, 2, "Sneezle.png"),
    CraftMonster("Loan", 10, 2, 2, "Loan.png"),
    CraftMonster("Meep", 10, 2, 2, "Meep.png"),
    CraftMonster("Skeleton", 10, 2, 2, "Skeleton.png"),
    CraftMonster("Pinchington", 10, 2, 2, "Pinchington.png"),
    CraftMonster("Drooplet", 10, 2, 2, "Drooplet.png"),
  }

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
