package app

import (
	"TouchTyper/config"
	"TouchTyper/utils"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func drawKeyboard(context *Context, textBox rl.Rectangle, center rl.Vector2) {
	const sizeOfKey = 35
	const margin = 5

	theme := context.Themes[context.SelectedTheme]
	sizeOfCharacter := rl.MeasureTextEx(context.Fonts.TinyFont.Font, "a", context.Fonts.TinyFont.Size, 1)

	if rl.IsKeyPressed(rl.KeyBackspace) {
		config.CursorStayVisibleTimer = 1
	}

	for i, row := range config.Keyboard {
		totalWidth := sizeOfKey * len(row)
		if len(row) == 1 && row[0] == ' ' {
			totalWidth = 200
		} else {
			totalWidth += margin * (len(row) - 1)
		}

		position := rl.Vector2{
			X: center.X - float32(totalWidth)/2.0,
			Y: textBox.Y + sizeOfCharacter.Y*8 + float32(sizeOfKey*i) + float32(margin*i),
		}

		for _, key := range row {
			rect := rl.Rectangle{
				X:      position.X,
				Y:      position.Y,
				Width:  sizeOfKey,
				Height: sizeOfKey,
			}

			if len(row) == 1 && key == ' ' {
				rect.Width = 200
			}

			rl.DrawRectangleRoundedLines(rect, 0.1, 5, theme.Text)

			color := theme.Text
			keyPosition := rl.Vector2{
				X: rect.X + rect.Width/2.0 - sizeOfCharacter.X/2.0,
				Y: rect.Y + rect.Height/2.0 - sizeOfCharacter.Y/2.0,
			}

			// Check if key is pressed
			keyPressed := false
			if key == ' ' {
				keyPressed = rl.IsKeyDown(rl.KeySpace)
			} else {
				// Convert rune to raylib key code
				keyCode := utils.RuneToKeyCode(key)
				if keyCode != -1 {
					keyPressed = rl.IsKeyDown(int32(keyCode))
				}
			}

			if keyPressed {
				rl.DrawRectangleRounded(rect, 0.1, 5, theme.Cursor)
				color = theme.Background
				config.CursorStayVisibleTimer = 1
			}

			drawMonospaceText(context.Fonts.TinyFont.Font, string(key), keyPosition, context.Fonts.TinyFont.Size, color)
			position.X += rect.Width + margin
		}
	}
}
