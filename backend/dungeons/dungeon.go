package dungeons
import (
	"fmt"
	"eef.dungeoneer/backend/monsters"	
)


type Room struct {
	Enemy *monsters.Monster
	NextRoom *Room
}


type Dungeon struct {
	FirstRoom *Room
	LastRoom *Room
	CurrentRoom *Room
	NumRooms int
}

func (d *Dungeon) Print() {
	fmt.Println(d.ToString());
}

func (d *Dungeon) ToString() string {
	// Will edit this later
	return d.FirstRoom.Enemy.ToString();
}

func NewRoom() *Room {
	room := &Room{
		Enemy: monsters.NewMonster("Goblin", 10, 2, 2, "goblin.png"),
		NextRoom: nil,
	};
	return room;
}

func NewDungeon() *Dungeon {
	// For now, use the NewRoom function to generate a room -- These values are set in stone at the moment, will change later
	generated_room := NewRoom();
	return &Dungeon{
		FirstRoom: generated_room,
		LastRoom: generated_room.NextRoom,
		CurrentRoom: generated_room,
		NumRooms: 2,
	}
}