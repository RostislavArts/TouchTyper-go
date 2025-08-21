package config

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

var (
	CursorPosition         rl.Vector2
	NewCursorPosition      rl.Vector2
	CursorSpeed            float32 = 20
	YOffset                float32 = 0
	NewYOffset             float32 = 0
	CursorOpacity          float32 = 1
	CursorStayVisibleTimer float32 = 0

	// Options visibility
	ShowThemesOptions   bool = false
	ShowWordListOptions bool = false
	ShowCursorOptions   bool = false

	// Animation variables
	TargetBarHeight float32 = 0
	BarHeight       float32 = 0

	// Global variables for punctuation generation
	UseCapitalNext  bool = false
	PreviousWasDash bool = false
)

var Keyboard = [][]rune{
	{'q', 'w', 'e', 'r', 't', 'y', 'u', 'i', 'o', 'p', '[', ']'},
	{'a', 's', 'd', 'f', 'g', 'h', 'j', 'k', 'l', ';', '\''},
	{'z', 'x', 'c', 'v', 'b', 'n', 'm', ',', '.', '/'},
	{' '},
}

var Quotes = [][2]rune{
	{'\'', '\''},
	{'"', '"'},
	{'(', ')'},
}

var Punctuations = map[string]string{
	"are":    "aren't",
	"can":    "can't",
	"could":  "couldn't",
	"did":    "didn't",
	"does":   "doesn't",
	"do":     "don't",
	"had":    "hadn't",
	"has":    "hasn't",
	"have":   "haven't",
	"is":     "isn't",
	"must":   "mustn't",
	"should": "shouldn't",
	"was":    "wasn't",
	"were":   "weren't",
	"will":   "won't",
	"would":  "wouldn't",
}
