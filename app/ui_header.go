package app

import (
	"TouchTyper/config"
	"TouchTyper/utils"
	"fmt"
	"strconv"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func header(context *Context) {
	width := min(context.ScreenWidth-config.PADDING*2, config.MAX_WIDTH)
	theme := context.Themes[context.SelectedTheme]
	center := utils.GetCenter(context.ScreenWidth, context.ScreenHeight)

	topLeftPosition := rl.Vector2{
		X: center.X - float32(width)/2.0,
		Y: config.PADDING,
	}

	topRightPosition := rl.Vector2{
		X: center.X + float32(width)/2.0,
		Y: config.PADDING,
	}

	// Draw title
	var color rl.Color
	var text string

	if !context.TestRunning {
		color = theme.Correct
		switch context.CurrentScreen {
		case config.SCREEN_TEST:
			text = "start typing"
		case config.SCREEN_RESULT:
			text = "result"
		}
	} else {
		color = theme.Text
		if context.TestSettings.TestMode == config.TEST_MODE_WORDS {
			text = fmt.Sprintf("%d/%d", len(context.Input), len(context.Sentence))
		} else {
			text = fmt.Sprintf("%ds", int(rl.GetTime()-context.TestStartTime))
		}
	}

	if !context.TestRunning {
		rl.DrawTextEx(context.Fonts.TitleFont.Font, text, topLeftPosition, context.Fonts.TitleFont.Size, 1, color)
	} else {
		drawMonospaceText(context.Fonts.TitleFont.Font, text, topLeftPosition, context.Fonts.TitleFont.Size, color)
	}

	// Draw options
	startingPosition := topRightPosition
	sizeOfCharacter := rl.MeasureTextEx(context.Fonts.TinyFont.Font, "a", context.Fonts.TinyFont.Size, 1)

	firstOptions := []string{"numbers", "punctuation"}
	drawOptions(context, firstOptions, startingPosition, sizeOfCharacter, false)

	startingPosition = topRightPosition
	startingPosition.Y += sizeOfCharacter.Y
	secondOptions := []string{"words", "time"}
	drawOptions(context, secondOptions, startingPosition, sizeOfCharacter, false)

	startingPosition.X = topRightPosition.X
	startingPosition.Y += sizeOfCharacter.Y

	var thirdOptions []string
	for _, amount := range context.TestSettings.TestModeAmounts {
		thirdOptions = append(thirdOptions, strconv.Itoa(amount))
	}
	drawOptions(context, thirdOptions, startingPosition, sizeOfCharacter, true)

	// Draw progress bar
	config.BarHeight = utils.Lerp(config.BarHeight, config.TargetBarHeight, rl.GetFrameTime()*4)
	percentage := float32(0)
	amount := context.TestSettings.TestModeAmounts[context.TestSettings.SelectedAmount]

	if context.TestRunning {
		config.TargetBarHeight = 5
	} else {
		config.TargetBarHeight = 0
	}

	switch context.TestSettings.TestMode {
	case config.TEST_MODE_TIME:
		percentage = float32(rl.GetTime()-context.TestStartTime) / float32(amount)
	case config.TEST_MODE_WORDS:
		if len(context.Sentence) > 0 {
			percentage = float32(context.FurthestVisitedIndex+1) / float32(len(context.Sentence))
		}
	}

	rl.DrawRectangle(0, 0, int32(float32(context.ScreenWidth)*percentage), int32(config.BarHeight), theme.Correct)
}
