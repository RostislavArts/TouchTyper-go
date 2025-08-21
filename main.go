package main

import (
	"TouchTyper/app"
	"TouchTyper/config"

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
	defer context.Unload()

	app.RestartTest(context, false)

	for !rl.WindowShouldClose() {
		app.Loop(context)
	}
}
