package main

import (
	"container/list"
	"context"
	"log"
	"eef.dungeoneer/backend/dungeons"
	"eef.dungeoneer/backend/entities"
)

// App struct
type App struct {
	ctx context.Context
	num_rooms int
	current_dungeon *dungeons.Dungeon
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

func (a * App) AddTextToQueue(text string) {
	a.text_queue.PushBack(text);
	log.Default().Println("Added text to queue: " + text);
}

func (a* App) CheckTextQueue() string {
	if(a.text_queue.Len() > 0) {
		return a.text_queue.Remove(a.text_queue.Front()).(string);
	}
	return "";
}
