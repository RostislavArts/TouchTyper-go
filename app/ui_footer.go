package app

import (
	"TouchTyper/config"
	"TouchTyper/utils"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func footer(context *Context) {
	sizeOfCharacter := rl.MeasureTextEx(context.Fonts.TinyFont.Font, "a", context.Fonts.TinyFont.Size, 1)
	theme := context.Themes[context.SelectedTheme]

	width := min(context.ScreenWidth-config.PADDING*2, config.MAX_WIDTH)
	center := utils.GetCenter(context.ScreenWidth, context.ScreenHeight)

	bottomLeftPosition := rl.Vector2{
		X: center.X - float32(width)/2.0,
	}

	bottomRightPosition := rl.Vector2{
		X: center.X + float32(width)/2.0,
		Y: float32(context.ScreenHeight) - config.PADDING,
	}

	// Draw version
	versionPosition := bottomRightPosition
	versionPosition.Y -= sizeOfCharacter.Y
	versionPosition.X -= float32(len(config.VERSION)) * sizeOfCharacter.X

	drawMonospaceText(context.Fonts.TinyFont.Font, config.VERSION, versionPosition, context.Fonts.TinyFont.Size, theme.Text)

	// Draw shortcuts
	shortcut := "shift  +  enter  - repeat test"
	position := center
	position.Y = float32(context.ScreenHeight) - (config.PADDING + sizeOfCharacter.Y)
	position.X -= float32(len(shortcut)) * sizeOfCharacter.X / 2.0

	drawMonospaceText(context.Fonts.TinyFont.Font, shortcut, position, context.Fonts.TinyFont.Size, theme.Text)

	// Draw key boxes
	rect := rl.Rectangle{
		X:      position.X - 4,
		Y:      position.Y - 2,
		Width:  sizeOfCharacter.X*5 + 8,
		Height: sizeOfCharacter.Y + 4,
	}
	rl.DrawRectangleRoundedLines(rect, 0.1, 5, theme.Text)

	rect.X = position.X + sizeOfCharacter.X*10 - 4
	rl.DrawRectangleRoundedLines(rect, 0.1, 5, theme.Text)

	shortcut = "enter  -  new test"
	position.X = center.X - float32(len(shortcut))*sizeOfCharacter.X/2.0
	position.Y -= sizeOfCharacter.Y + 10
	drawMonospaceText(context.Fonts.TinyFont.Font, shortcut, position, context.Fonts.TinyFont.Size, theme.Text)

	rect.X = position.X - 4
	rect.Y = position.Y - 2
	rect.Width = sizeOfCharacter.X*5 + 8
	rl.DrawRectangleRoundedLines(rect, 0.1, 5, theme.Text)

	// Draw footer options
	themePosition := rl.Vector2{
		X: bottomLeftPosition.X,
		Y: bottomRightPosition.Y - sizeOfCharacter.Y,
	}

	if config.ShowThemesOptions {
		var themeOptions []string
		for _, theme := range context.Themes {
			themeOptions = append(themeOptions, theme.Name)
		}

		selected := optionSelect(context, themeOptions, context.SelectedTheme)
		if selected != -1 {
			context.SelectedTheme = selected
			config.ShowThemesOptions = false
		}
	}

	if textButtonClicked(context, themePosition, "theme") {
		config.ShowThemesOptions = !config.ShowThemesOptions
	} else if rl.IsMouseButtonPressed(rl.MouseButtonLeft) {
		config.ShowThemesOptions = false
	}

	worldlistPosition := rl.Vector2{
		X: themePosition.X + sizeOfCharacter.X*6,
		Y: themePosition.Y,
	}

	if config.ShowWordListOptions {
		var wordListOptions []string
		for _, wordList := range context.WordLists {
			wordListOptions = append(wordListOptions, wordList.Name)
		}

		selected := optionSelect(context, wordListOptions, context.SelectedWordList)
		if selected != -1 {
			context.SelectedWordList = selected
			config.ShowWordListOptions = false
			RestartTest(context, false)
		}
	}

	if textButtonClicked(context, worldlistPosition, "word list") {
		config.ShowWordListOptions = !config.ShowWordListOptions
	} else if rl.IsMouseButtonPressed(rl.MouseButtonLeft) {
		config.ShowWordListOptions = false
	}

	cursorPosition := rl.Vector2{
		X: bottomRightPosition.X - sizeOfCharacter.X*float32(len(config.VERSION)+7),
		Y: bottomRightPosition.Y - sizeOfCharacter.Y,
	}

	if config.ShowCursorOptions {
		cursorOptions := []string{"Block", "Line", "Underline"}
		selected := optionSelect(context, cursorOptions, int(context.CursorStyle))
		if selected != -1 {
			context.CursorStyle = config.CursorStyle(selected)
			config.ShowCursorOptions = false
		}
	}

	if textButtonClicked(context, cursorPosition, "cursor") {
		config.ShowCursorOptions = !config.ShowCursorOptions
	} else if rl.IsMouseButtonPressed(rl.MouseButtonLeft) {
		config.ShowCursorOptions = false
	}

	soundText := "sound off"
	if context.SoundOn {
		soundText = "sound on"
	}

	soundPosition := rl.Vector2{
		X: cursorPosition.X - sizeOfCharacter.X*float32(len(soundText)+1),
		Y: cursorPosition.Y,
	}

	if textButtonClicked(context, soundPosition, soundText) {
		context.SoundOn = !context.SoundOn
	}
}
