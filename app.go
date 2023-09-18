package main

import (
	"context"
	"fmt"
	"math/rand"
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

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	a.num_rooms = rand.Intn(10);
	return fmt.Sprintf("Hello %s, you'll be running a dungeon with %d rooms!", name, a.num_rooms);
}
