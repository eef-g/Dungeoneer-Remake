package main

import (
	"context"
	"eef.dungeoneer/backend/dungeons"
)

// App struct
type App struct {
	ctx context.Context
	num_rooms int
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
}

// GenerateDungeon returns a string value of the generated dungeon [This is just for testing purposes]
func (a *App) GenerateDungeon() string{
	dungeon := dungeons.NewDungeon();
	dungeon.GenerateDungeon();
	return dungeon.ToString();
}
