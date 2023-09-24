package main

import (
	"container/list"
	"context"
	"fmt"
	"log"
	"eef.dungeoneer/backend/dungeons"
	"eef.dungeoneer/backend/entities"
)

// App struct
type App struct {
	ctx context.Context
	num_rooms int
	current_dungeon *dungeons.Dungeon
  curr_player *entities.Player
	text_queue list.List
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	a.num_rooms = 0
	a.current_dungeon = dungeons.NewDungeon()
  a.curr_player = CreatePlayer()
	a.text_queue = list.List{}
}

// GenerateDungeon returns a string value of the generated dungeon [This is just for testing purposes]
func (a *App) GenerateDungeon() string{
	if(a.current_dungeon != nil) {
		a.current_dungeon = dungeons.NewDungeon();
	}
	a.current_dungeon.GenerateDungeon();
	return a.current_dungeon.ToString();
}

func (a *App) WipeDungeon() {
	if(a.current_dungeon != nil) {
		a.current_dungeon.WipeDungeon();
	}
}

func (a *App) GetRoomStats() *dungeons.Room {
	if(a.current_dungeon != nil) {
		return a.current_dungeon.CurrentRoom;
	}
	return nil;
}

func (a *App) ProgressInDungeon(){
	if(a.current_dungeon != nil) {
		a.current_dungeon.CurrentRoom = a.current_dungeon.CurrentRoom.NextRoom;
		// If the current room is nil, then the dungeon is finished
		// Need to create a dummy room & monster so the frontend knows we're done
		if (a.current_dungeon.CurrentRoom == nil) {
			a.current_dungeon.CurrentRoom = &dungeons.Room{
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
}

func (a *App) DungeonToString() string{
	if(a.current_dungeon != nil) {
		return a.current_dungeon.ToString();
	}
	return "No dungeon generated yet";
}

func (a *App) AddTextToQueue(text string) {
	a.text_queue.PushBack(text);
	log.Default().Println("Added text to queue: " + text);
}

func (a *App) CheckTextQueue() string {
	if(a.text_queue.Len() > 0) {
		return a.text_queue.Remove(a.text_queue.Front()).(string);
	}
	return "";
}

func CreatePlayer() *entities.Player{
  // Eventually, want to go and replace player creation w/ reading a savel file
  player := entities.NewPlayer("Eef", 20, 5);
  return player
}


// Combat Logic
type ResultPacket struct{
  NextRoom bool
  // Will add more here
};


func (a *App) CombatTurn(decision string) *ResultPacket{
  // Debug logging -- for sanity checking when testing
  log_output := fmt.Sprintf("%v has decided to %v", a.curr_player.Name, decision);
  log.Print(log_output);
  
  // Prep the return packet for the outcome
  return_packet := ResultPacket{
    NextRoom: false,
  }

  // Set up the variables we'll use
  current_enemy := a.current_dungeon.CurrentRoom.Enemy;
  player := a.curr_player;

  switch decision {
    case "attack":
      log.Print("Attacking...");
      start_health := current_enemy.Base.Health;
      current_enemy.Base.Health = entities.Attack(&player.Entity, &current_enemy.Base);
      dmg_dealt := start_health - current_enemy.Base.Health;
      a.AddTextToQueue(fmt.Sprintf("\n> %v does %d damage to the %v!", player.Entity.Name, dmg_dealt, current_enemy.Base.Name));


      // Eventually have the monster fight back
      // Room progression Logic
      if(current_enemy.Base.Health <= 0) {
        a.ProgressInDungeon();
        return_packet.NextRoom = true;
      }

    case "block":
      log.Print("Blocking...");
      // Will do this logic later :)

    case "heal":
      log.Print("Healing...");
      // Will do this logic later :)

    default:
      // Catch anything
      log.Print("Unkown case detected. Ignoring.");
  }
  return &return_packet;
}
