package app

import (
	"TouchTyper/config"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func Loop(context *Context) {
	context.ScreenHeight = int32(rl.GetScreenHeight())
	context.ScreenWidth = int32(rl.GetScreenWidth())

	theme := context.Themes[context.SelectedTheme]

	// Handle F11 for fullscreen
	if rl.IsKeyPressed(rl.KeyF11) {
		if rl.IsWindowFullscreen() {
			rl.ToggleFullscreen()
		} else {
			rl.ToggleFullscreen()
		}
	}

	key := rl.GetCharPressed()

	if context.MouseOnClickable {
		rl.SetMouseCursor(rl.MouseCursorPointingHand)
	} else {
		rl.SetMouseCursor(rl.MouseCursorDefault)
	}
	context.MouseOnClickable = false

	if rl.IsKeyPressed(rl.KeyEnter) {
		RestartTest(context, rl.IsKeyDown(rl.KeyLeftShift) || rl.IsKeyDown(rl.KeyRightShift))
	}

	if context.TestRunning {
		elapsed := rl.GetTime() - context.TestStartTime
		if elapsed >= float64(context.TestSettings.TestModeAmounts[context.TestSettings.SelectedAmount]) &&
			context.TestSettings.TestMode == config.TEST_MODE_TIME {
			endTest(context)
		}
	}

	rl.BeginDrawing()
	rl.ClearBackground(theme.Background)

	header(context)

	switch context.CurrentScreen {
	case config.SCREEN_TEST:
		handleTestInput(context, key)
		typingTest(context)
	case config.SCREEN_RESULT:
		drawResult(context)
	}

	footer(context)

	if rl.IsMouseButtonPressed(rl.MouseButtonLeft) {
		context.SaveSettings()
	}

	rl.EndDrawing()
}
