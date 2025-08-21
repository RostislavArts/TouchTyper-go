package app

import (
	"TouchTyper/config"
	"TouchTyper/utils"
	"math/rand"
	"strconv"
	"strings"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func handleTestInput(context *Context, key int32) {
	if rl.IsKeyPressed(rl.KeyBackspace) {
		if context.SoundOn {
			rl.PlaySound(context.Sounds.ClickSound1)
		}

		if len(context.Input) > 0 {
			// CTRL + Backspace
			if rl.IsKeyDown(rl.KeyLeftControl) || rl.IsKeyDown(rl.KeyRightControl) {
				// Remove trailing spaces
				for len(context.Input) > 0 && context.Input[len(context.Input)-1] == ' ' {
					context.Input = context.Input[:len(context.Input)-1]
				}
				// Remove word
				for len(context.Input) > 0 && context.Input[len(context.Input)-1] != ' ' {
					context.Input = context.Input[:len(context.Input)-1]
				}
			} else {
				// Normal backspace
				context.Input = context.Input[:len(context.Input)-1]
			}
		}
	}

	if key != 0 && len(context.Input) < len(context.Sentence) {
		context.Input += string(rune(key))

		if context.SoundOn {
			rl.PlaySound(context.Sounds.ClickSound1)
		}

		if len(context.Input) == 1 && !context.TestRunning {
			context.TestRunning = true
			context.TestStartTime = rl.GetTime()
		}

		// End test when sentence is complete
		if context.TestRunning && len(context.Input) == len(context.Sentence) {
			endTest(context)
		}

		// Calculate correct and incorrect letters
		if len(context.Input) > context.FurthestVisitedIndex {
			inputRunes := []rune(context.Input)
			sentenceRunes := []rune(context.Sentence)

			if inputRunes[len(inputRunes)-1] != sentenceRunes[len(inputRunes)-1] {
				context.IncorrectLetters++
			} else {
				context.CorrectLetters++
			}

			// Add more words for time mode
			if context.TestSettings.TestMode == config.TEST_MODE_TIME {
				if sentenceRunes[len(inputRunes)-1] == ' ' {
					context.Sentence += " " + generateSentence(context, 1)
				}
			}
		}

		context.FurthestVisitedIndex = max(context.FurthestVisitedIndex, len(context.Input))
	}

	// Calculate score
	if context.TestRunning && (rl.GetTime()-context.TestStartTime) > 3 {
		elapsed := rl.GetTime() - context.TestStartTime
		wpm := float64(context.CorrectLetters) * (60 / elapsed) / 5.0
		raw := float64(context.CorrectLetters+context.IncorrectLetters) * (60 / elapsed) / 5.0
		context.WPM = int(wpm)
		context.Raw = int(raw)
		if context.CorrectLetters+context.IncorrectLetters > 0 {
			context.Accuracy = int(float64(context.CorrectLetters) / float64(context.CorrectLetters+context.IncorrectLetters) * 100)
		}
	}
}

func generateSentence(context *Context, numberOfWords int) string {
	if len(context.WordLists) == 0 || context.SelectedWordList >= len(context.WordLists) {
		return "no words available"
	}

	words := make([]string, len(context.WordLists[context.SelectedWordList].Words))
	copy(words, context.WordLists[context.SelectedWordList].Words)

	// Shuffle words
	for i := range words {
		j := rand.Intn(i + 1)
		words[i], words[j] = words[j], words[i]
	}

	if numberOfWords > len(words) {
		numberOfWords = len(words)
	}

	var output strings.Builder

	for i := 0; i < numberOfWords; i++ {
		word := words[i]

		if context.TestSettings.UsePunctuation {
			if replacement, exists := config.Punctuations[word]; exists {
				word = replacement
			}
		}

		inQuotes := context.TestSettings.UsePunctuation && (rand.Intn(11) == 10)
		itsDashTime := context.TestSettings.UsePunctuation && (rand.Intn(11) == 10) && !config.PreviousWasDash && !config.UseCapitalNext
		itsNumber := context.TestSettings.UseNumbers && (rand.Intn(11) == 10)

		if config.UseCapitalNext && len(word) > 0 {
			word = strings.ToUpper(string(word[0])) + word[1:]
		}

		if inQuotes {
			quote := config.Quotes[rand.Intn(len(config.Quotes))]
			word = string(quote[0]) + word + string(quote[1])
		} else if itsDashTime {
			word = "-"
		}

		if itsNumber {
			word = strconv.Itoa(rand.Intn(1001))
		}

		output.WriteString(word)

		config.UseCapitalNext = false

		// Add punctuation randomly
		if rand.Intn(11) > 8 && context.TestSettings.UsePunctuation && !config.PreviousWasDash && !itsDashTime && !inQuotes {
			switch rand.Intn(4) {
			case 0:
				output.WriteRune(',')
			case 1:
				output.WriteRune('.')
				config.UseCapitalNext = true
			case 2:
				output.WriteRune('!')
				config.UseCapitalNext = true
			case 3:
				output.WriteRune('?')
				config.UseCapitalNext = true
			}
		}

		config.PreviousWasDash = itsDashTime

		if i < numberOfWords-1 {
			output.WriteRune(' ')
		}
	}

	return output.String()
}

func RestartTest(context *Context, repeat bool) {
	if !repeat {
		amount := context.TestSettings.TestModeAmounts[context.TestSettings.SelectedAmount]

		// For time mode, start with more words
		if context.TestSettings.TestMode == config.TEST_MODE_TIME {
			amount = 120
		}

		context.Sentence = generateSentence(context, amount)

		if context.TestSettings.UsePunctuation && len(context.Sentence) > 0 {
			context.Sentence = strings.ToUpper(string(context.Sentence[0])) + context.Sentence[1:]
			if context.TestSettings.TestMode == config.TEST_MODE_WORDS && !config.UseCapitalNext {
				context.Sentence += "."
			}
		}
	}

	context.Input = ""
	context.CurrentScreen = config.SCREEN_TEST
	context.WPM = 0
	context.CPM = 0
	context.Accuracy = 0
	context.Raw = 0
	context.TestRunning = false
	context.CorrectLetters = 0
	context.IncorrectLetters = 0
	context.FurthestVisitedIndex = -1
}

func endTest(context *Context) {
	context.TestRunning = false
	context.CurrentScreen = config.SCREEN_RESULT
	context.TestEndTime = rl.GetTime()
}

func typingTest(context *Context) {
	// Implementation of the typing test display
	sizeOfCharacter := rl.MeasureTextEx(context.Fonts.TypingTestFont.Font, "a", context.Fonts.TypingTestFont.Size, 1)
	theme := context.Themes[context.SelectedTheme]

	width := float32(min(context.ScreenWidth-config.PADDING*2, config.MAX_WIDTH))
	height := sizeOfCharacter.Y * 3

	center := utils.GetCenter(context.ScreenWidth, context.ScreenHeight)

	// Animate scroll and cursor
	speed := config.CursorSpeed * rl.GetFrameTime()
	if speed <= 0 || speed > 1 {
		speed = 1
	}
	config.YOffset = utils.Lerp(config.YOffset, config.NewYOffset, speed)
	config.CursorPosition.X = utils.Lerp(config.CursorPosition.X, config.NewCursorPosition.X, speed)
	config.CursorPosition.Y = utils.Lerp(config.CursorPosition.Y, config.NewCursorPosition.Y, speed)

	// Cursor blink timer
	if config.CursorStayVisibleTimer > 0 {
		config.CursorStayVisibleTimer -= rl.GetFrameTime()
	} else {
		config.CursorStayVisibleTimer = 0
	}

	config.CursorOpacity = utils.SinPulse(1.5)

	// Break sentence into lines
	var lines []string
	currentLine := ""
	currentWord := ""

	sentenceRunes := []rune(context.Sentence)

	for i, char := range sentenceRunes {
		if char == ' ' || i == len(sentenceRunes)-1 {
			currentWord += string(char)

			widthOfWord := float32(len(currentWord)) * sizeOfCharacter.X
			widthOfNewLine := widthOfWord + float32(len(currentLine))*sizeOfCharacter.X

			if widthOfNewLine > width-config.PADDING*2 {
				lines = append(lines, currentLine)
				currentLine = ""
			}

			currentLine += currentWord
			currentWord = ""
		} else {
			currentWord += string(char)
		}
	}

	if currentLine != "" {
		lines = append(lines, currentLine)
	}

	textBox := rl.Rectangle{
		X:      center.X - width/2.0,
		Y:      center.Y - height/2.0 - 100,
		Width:  width,
		Height: height,
	}

	if textBox.Y < 95 {
		textBox.Y = 95
	}

	// Begin scissor mode for text scrolling
	rl.BeginScissorMode(int32(textBox.X), int32(textBox.Y), int32(textBox.Width), int32(textBox.Height))

	currentLineY := textBox.Y - config.YOffset
	characterIndex := 0
	lineNumber := 1

	inputRunes := []rune(context.Input)

	for _, line := range lines {
		lineRunes := []rune(line)
		widthOfLine := sizeOfCharacter.X * float32(len(lineRunes))
		currentLetterX := center.X - widthOfLine/2

		for _, letter := range lineRunes {
			color := theme.Text

			if len(inputRunes) > characterIndex {
				if letter == inputRunes[characterIndex] {
					color = theme.Correct
				} else {
					color = theme.Wrong

					// Draw underline for wrong spaces
					if letter == ' ' {
						rl.DrawTextEx(context.Fonts.TypingTestFont.Font, "_",
							rl.Vector2{X: currentLetterX, Y: currentLineY},
							context.Fonts.TypingTestFont.Size, 1, color)
					}
				}
			}

			// Draw character
			rl.DrawTextEx(context.Fonts.TypingTestFont.Font, string(letter),
				rl.Vector2{X: currentLetterX, Y: currentLineY},
				context.Fonts.TypingTestFont.Size, 1, color)

			// Handle cursor
			if characterIndex == len(inputRunes) {
				if lineNumber > 2 {
					config.NewYOffset = float32(lineNumber-1)*sizeOfCharacter.Y - sizeOfCharacter.Y
				} else {
					config.NewYOffset = 1
				}
				config.NewCursorPosition = rl.Vector2{X: currentLetterX, Y: currentLineY}

				cursorColor := theme.Cursor
				blink := config.CursorOpacity
				if config.CursorStayVisibleTimer != 0 {
					blink = 1
				}

				// Draw cursor
				if blink > 0.5 {
					switch context.CursorStyle {
					case config.CURSOR_BLOCK:
						rl.DrawRectangle(int32(config.CursorPosition.X+1), int32(config.CursorPosition.Y),
							int32(sizeOfCharacter.X), int32(sizeOfCharacter.Y), cursorColor)
					case config.CURSOR_LINE:
						rl.DrawRectangle(int32(config.CursorPosition.X), int32(config.CursorPosition.Y),
							2, int32(sizeOfCharacter.Y), cursorColor)
					case config.CURSOR_UNDERLINE:
						rl.DrawRectangle(int32(config.CursorPosition.X+1), int32(config.CursorPosition.Y+sizeOfCharacter.Y),
							int32(sizeOfCharacter.X), 3, cursorColor)
					}
				}
			}

			currentLetterX += sizeOfCharacter.X
			characterIndex++
		}

		currentLineY += sizeOfCharacter.Y
		lineNumber++
	}

	rl.EndScissorMode()

	// Draw keyboard
	drawKeyboard(context, textBox, center)
}
