package app

import (
	"TouchTyper/utils"
	"strconv"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func drawMonospaceText(font rl.Font, text string, position rl.Vector2, fontSize float32, color rl.Color) {
	sizeOfCharacter := rl.MeasureTextEx(font, "a", fontSize, 1)

	for _, char := range text {
		rl.DrawTextEx(font, string(char), position, fontSize, 1, color)
		position.X += sizeOfCharacter.X
	}
}

func textButtonClicked(context *Context, position rl.Vector2, text string) bool {
	sizeOfCharacter := rl.MeasureTextEx(context.Fonts.TinyFont.Font, "a", context.Fonts.TinyFont.Size, 1)
	theme := context.Themes[context.SelectedTheme]

	rect := rl.Rectangle{
		X:      position.X,
		Y:      position.Y,
		Width:  sizeOfCharacter.X * float32(len(text)),
		Height: sizeOfCharacter.Y,
	}

	color := theme.Text

	if rl.CheckCollisionPointRec(rl.GetMousePosition(), rect) {
		context.MouseOnClickable = true
		color = theme.Highlight
		if rl.IsMouseButtonPressed(rl.MouseButtonLeft) {
			return true
		}
	}

	drawMonospaceText(context.Fonts.TinyFont.Font, text, position, context.Fonts.TinyFont.Size, color)
	return false
}

func drawResult(context *Context) {
	sizeOfCharacter := rl.MeasureTextEx(context.Fonts.BigFont.Font, "a", context.Fonts.BigFont.Size, 1)
	theme := context.Themes[context.SelectedTheme]

	center := utils.GetCenter(context.ScreenWidth, context.ScreenHeight)

	scores := [][]string{
		{"wpm", strconv.Itoa(context.WPM)},
		{"raw", strconv.Itoa(context.Raw)},
		{"acc", strconv.Itoa(context.Accuracy) + "%"},
		{"time", strconv.Itoa(int(context.TestEndTime-context.TestStartTime)) + "s"},
	}

	totalWidth := float32(0)
	for _, score := range scores {
		totalWidth += float32(len(score[1])) * sizeOfCharacter.X
	}
	totalWidth += float32(len(scores)-1) * sizeOfCharacter.X

	position := rl.Vector2{
		X: center.X - totalWidth/2.0,
		Y: center.Y - sizeOfCharacter.Y/2.0,
	}

	for _, score := range scores {
		color := theme.Correct
		drawMonospaceText(context.Fonts.BigFont.Font, score[1], position, context.Fonts.BigFont.Size, color)

		nPosition := position
		nPosition.Y -= context.Fonts.TypingTestFont.Size / 2.0
		color = theme.Text
		drawMonospaceText(context.Fonts.TypingTestFont.Font, score[0], nPosition, context.Fonts.TypingTestFont.Size, color)

		position.X += float32(len(score[1]))*sizeOfCharacter.X + sizeOfCharacter.X
	}
}
