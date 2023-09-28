package dungeons
import (
	"fmt"
	"math/rand"
	"eef.dungeoneer/backend/entities"	
)


type Room struct {
	Enemy *entities.Monster
	NextRoom *Room
}


type Dungeon struct {
	FirstRoom *Room
	LastRoom *Room
	CurrentRoom *Room
	NumRooms int
	MonsterLog *entities.Monsters
}

func NewDungeon() *Dungeon {
	dungeon := &Dungeon{
		FirstRoom: nil,
		LastRoom: nil,
		CurrentRoom: nil,
		NumRooms: 0,
		MonsterLog: entities.GenerateMonsters(),
	};
	return dungeon;
}

func (d *Dungeon) GenerateDungeon() {
	// Make sure the dungeon is empty first
	d.WipeDungeon();
	// Generate a random number of rooms

	d.NumRooms = rand.Intn(10);
	// Eventually, we'll make the min rooms different
	for d.NumRooms < 3 {
        d.NumRooms = rand.Intn(10)
    }

	// Loop through the number of rooms and generate a new room for each one and assign it to the dungeon
	for i := 0; i < d.NumRooms; i++ {
		// Create a new room
		room := &Room{
			Enemy: d.MonsterLog.Monsters[rand.Intn(len(d.MonsterLog.Monsters))],
			NextRoom: nil,
		};
		// If this is the first room, assign it to the first room
		if d.FirstRoom == nil {
			d.FirstRoom = room;
		} else {
			// Otherwise, assign it to the last room
			d.LastRoom.NextRoom = room;
		}
		// Assign the last room to the new room
		d.LastRoom = room;
		// If this is the first room, assign it to the current room
		if d.CurrentRoom == nil {
			d.CurrentRoom = room;
		}
	}

}

func (d *Dungeon) WipeDungeon() {
	if d.FirstRoom == nil {
		// Loop through the dungeon and delete all the rooms
		for d.FirstRoom != nil {
			next := d.FirstRoom.NextRoom
			d.FirstRoom = nil
			d.FirstRoom = next
			d.NumRooms--
		}
	}
	d.LastRoom = nil
	d.CurrentRoom = nil
}

func (d *Dungeon) Print() {
	fmt.Println(d.ToString());
}


func (d *Dungeon) ProgressInDungeon(){
  d.CurrentRoom = d.CurrentRoom.NextRoom;
  // If the current room is nil, then the dungeon is finished
  // Need to create a dummy room & monster so the frontend knows we're done
  if (d.CurrentRoom == nil) {
    d.CurrentRoom = &Room{
      Enemy: &entities.Monster{
        Base: entities.Entity{
          Name: "Finished",
          Health: 0,
          Damage: 0,
        },
        AttackBonus: 0,
        Image: "",
      },
      NextRoom: nil,
    }
  }

}	

func (d *Dungeon) ToString() string {
	// Loop through the dungeon and print out the rooms
	fmt.Println("");
	room := d.FirstRoom;
	dungeon_string := "=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-\n";
	for room != nil {
		dungeon_string += fmt.Sprintf("Room: %s\n", room.Enemy.ToString());
		room = room.NextRoom;
	}
	dungeon_string += "=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-\n";
	return dungeon_string;
}
