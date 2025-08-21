package config

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type FontData struct {
	Font rl.Font
	Size float32
}

type Sounds struct {
	ClickSound1 rl.Sound
}

type Fonts struct {
	TitleFont      FontData
	TypingTestFont FontData
	TinyFont       FontData
	BigFont        FontData
}

type Theme struct {
	Name       string
	Background rl.Color
	Cursor     rl.Color
	Text       rl.Color
	Wrong      rl.Color
	Correct    rl.Color
	Highlight  rl.Color
}

type WordList struct {
	Name  string
	Words []string
}

type TestSettings struct {
	UsePunctuation  bool
	UseNumbers      bool
	TestMode        TestMode
	TestModeAmounts []int
	SelectedAmount  int
}
