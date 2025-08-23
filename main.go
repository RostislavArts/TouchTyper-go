package main

import (
	"TouchTyper/app"
	"TouchTyper/config"
	"log"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	rl.SetConfigFlags(rl.FlagMsaa4xHint | rl.FlagVsyncHint)
	rl.InitWindow(800, 500, config.PROJECT_NAME)
	defer rl.CloseWindow()

	rl.SetWindowState(rl.FlagWindowResizable | rl.FlagWindowMaximized)
	rl.InitAudioDevice()
	defer rl.CloseAudioDevice()

	context := app.NewContext()
	context.Load()
	defer func() {
		err := context.Unload()
		if err != nil {
			log.Printf("Error saving user data! Details: %s", err)
		}
	}()

	app.RestartTest(context, false)

	for !rl.WindowShouldClose() {
		app.Loop(context)
	}
}
