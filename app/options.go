package app

import (
	"TouchTyper/config"
	"TouchTyper/utils"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func drawOptions(context *Context, options []string, startingPosition rl.Vector2, sizeOfCharacter rl.Vector2, isAmounts bool) {
	theme := context.Themes[context.SelectedTheme]

	for i := len(options) - 1; i >= 0; i-- {
		word := options[i]

		optionPosition := startingPosition
		optionPosition.X -= float32(len(word)) * sizeOfCharacter.X
		color := theme.Text

		optionRect := rl.Rectangle{
			X:      optionPosition.X,
			Y:      optionPosition.Y,
			Width:  float32(len(word))*sizeOfCharacter.X + sizeOfCharacter.X,
			Height: sizeOfCharacter.Y,
		}

		if rl.CheckCollisionPointRec(rl.GetMousePosition(), optionRect) {
			color = theme.Highlight
			context.MouseOnClickable = true

			if rl.IsMouseButtonPressed(rl.MouseButtonLeft) {
				switch word {
				case "numbers":
					context.TestSettings.UseNumbers = !context.TestSettings.UseNumbers
				case "punctuation":
					context.TestSettings.UsePunctuation = !context.TestSettings.UsePunctuation
				case "words":
					context.TestSettings.TestMode = config.TEST_MODE_WORDS
				case "time":
					context.TestSettings.TestMode = config.TEST_MODE_TIME
				}

				if isAmounts {
					context.TestSettings.SelectedAmount = i
				}

				if context.CurrentScreen == config.SCREEN_TEST {
					RestartTest(context, false)
				}
			}
		}

		// Check if option is selected
		if (word == "punctuation" && context.TestSettings.UsePunctuation) ||
			(word == "numbers" && context.TestSettings.UseNumbers) ||
			(word == "time" && context.TestSettings.TestMode == config.TEST_MODE_TIME) ||
			(word == "words" && context.TestSettings.TestMode == config.TEST_MODE_WORDS) ||
			(isAmounts && i == context.TestSettings.SelectedAmount) {
			color = theme.Correct
		}

		drawMonospaceText(context.Fonts.TinyFont.Font, word, optionPosition, context.Fonts.TinyFont.Size, color)
		startingPosition.X = optionPosition.X - sizeOfCharacter.X
	}
}

func optionSelect(context *Context, options []string, selected int) int {
	const optionSelectWidth = 200
	const optionPadding = 5
	const optionsContainerPadding = 10

	sizeOfCharacter := rl.MeasureTextEx(context.Fonts.TinyFont.Font, "a", context.Fonts.TinyFont.Size, 1)
	theme := context.Themes[context.SelectedTheme]

	height := float32(len(options)) * (sizeOfCharacter.Y + optionPadding*2)
	center := utils.GetCenter(context.ScreenWidth, context.ScreenHeight)

	rect := rl.Rectangle{
		X:      center.X - optionSelectWidth/2.0,
		Y:      center.Y - (height+optionsContainerPadding*2)/2.0,
		Width:  optionSelectWidth,
		Height: height + optionsContainerPadding*2,
	}

	// Draw overlay
	rl.DrawRectangle(0, 0, context.ScreenWidth, context.ScreenHeight, rl.Color{R: 0, G: 0, B: 0, A: 100})
	rl.DrawRectangleRec(rect, theme.Background)

	startingPosition := rl.Vector2{
		X: rect.X + optionsContainerPadding,
		Y: rect.Y + optionsContainerPadding,
	}

	for i, option := range options {
		optionRect := rl.Rectangle{
			X:      startingPosition.X - optionsContainerPadding,
			Y:      startingPosition.Y,
			Width:  rect.Width,
			Height: sizeOfCharacter.Y + optionPadding*2,
		}

		color := theme.Text
		mouseOnOption := rl.CheckCollisionPointRec(rl.GetMousePosition(), optionRect)

		if mouseOnOption {
			context.MouseOnClickable = true
			if rl.IsMouseButtonPressed(rl.MouseButtonLeft) {
				return i
			}
		}

		if i == selected || mouseOnOption {
			rl.DrawRectangleRec(optionRect, theme.Cursor)
			color = theme.Background
		}

		startingPosition.Y += optionPadding
		drawMonospaceText(context.Fonts.TinyFont.Font, option, startingPosition, context.Fonts.TinyFont.Size, color)
		startingPosition.Y += sizeOfCharacter.Y + optionPadding
	}

	rl.DrawRectangleLinesEx(rect, 1, theme.Correct)
	return -1
}
