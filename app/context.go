package app

import (
	"TouchTyper/config"
	"bufio"
	"os"
	"path/filepath"
	"strings"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Context struct {
	ScreenWidth          int32
	ScreenHeight         int32
	CurrentScreen        config.Screen
	Sounds               config.Sounds
	WordLists            []config.WordList
	SelectedWordList     int
	Sentence             string
	Input                string
	TestStartTime        float64
	TestRunning          bool
	TestEndTime          float64
	SoundOn              bool
	WPM                  int
	CPM                  int
	Raw                  int
	IncorrectLetters     int
	CorrectLetters       int
	FurthestVisitedIndex int
	Accuracy             int
	Fonts                config.Fonts
	Themes               []config.Theme
	SelectedTheme        int
	CursorStyle          config.CursorStyle
	TestSettings         config.TestSettings
	MouseOnClickable     bool
}

func NewContext() *Context {
	return &Context{}
}

func (c *Context) Load() {
	// Load themes
	c.Themes = []config.Theme{
		{
			Name:       "Arch",
			Background: rl.Color{R: 6, G: 7, B: 9, A: 255},
			Text:       rl.Color{R: 92, G: 96, B: 133, A: 255},
			Cursor:     rl.Color{R: 214, G: 227, B: 255, A: 255},
			Wrong:      rl.Red,
			Correct:    rl.Color{R: 214, G: 227, B: 255, A: 255},
			Highlight:  rl.Color{R: 214, G: 227, B: 255, A: 255},
		},
		{
			Name:       "Black",
			Background: rl.Color{R: 17, G: 17, B: 17, A: 255},
			Text:       rl.Color{R: 96, G: 96, B: 96, A: 255},
			Cursor:     rl.Color{R: 225, G: 225, B: 225, A: 255},
			Wrong:      rl.Color{R: 218, G: 51, B: 51, A: 255},
			Correct:    rl.Color{R: 225, G: 225, B: 225, A: 255},
			Highlight:  rl.Color{R: 225, G: 225, B: 225, A: 255},
		},
		{
			Name:       "White",
			Background: rl.Color{R: 238, G: 235, B: 226, A: 255},
			Text:       rl.Color{R: 153, G: 148, B: 127, A: 255},
			Cursor:     rl.White,
			Wrong:      rl.Color{R: 209, G: 97, B: 67, A: 255},
			Correct:    rl.Color{R: 17, G: 17, B: 17, A: 255},
			Highlight:  rl.Color{R: 17, G: 17, B: 17, A: 255},
		},
		{
			Name:       "Cyberpunk",
			Background: rl.Color{R: 13, G: 13, B: 13, A: 255},
			Text:       rl.Color{R: 84, G: 84, B: 84, A: 255},
			Cursor:     rl.Color{R: 208, G: 237, B: 87, A: 255},
			Wrong:      rl.Color{R: 248, G: 82, B: 74, A: 255},
			Correct:    rl.Color{R: 26, G: 214, B: 118, A: 255},
			Highlight:  rl.White,
		},
	}

	// Load fonts
	exePath, _ := os.Executable()
	basePath := filepath.Dir(exePath)

	fontPath := filepath.Join(basePath, "assets", "fonts", "JetBrainsMono-Regular.ttf")
	if _, err := os.Stat(fontPath); os.IsNotExist(err) {
		// Fallback to default font if custom font not found
		c.Fonts.TypingTestFont = config.FontData{Font: rl.GetFontDefault(), Size: 32}
		c.Fonts.TitleFont = config.FontData{Font: rl.GetFontDefault(), Size: 40}
		c.Fonts.TinyFont = config.FontData{Font: rl.GetFontDefault(), Size: 18}
		c.Fonts.BigFont = config.FontData{Font: rl.GetFontDefault(), Size: 90}
	} else {
		c.Fonts.TypingTestFont = config.FontData{Font: rl.LoadFontEx(fontPath, 32, nil), Size: 32}
		c.Fonts.TitleFont = config.FontData{Font: rl.LoadFontEx(fontPath, 40, nil), Size: 40}
		c.Fonts.TinyFont = config.FontData{Font: rl.LoadFontEx(fontPath, 18, nil), Size: 18}
		c.Fonts.BigFont = config.FontData{Font: rl.LoadFontEx(fontPath, 90, nil), Size: 90}
	}

	// Load word lists
	wordListsPath := filepath.Join(basePath, "assets", "word_lists")
	c.loadWordLists(wordListsPath)

	// Load sounds
	soundPath := filepath.Join(basePath, "assets", "audio", "otemu_browns.wav")
	if _, err := os.Stat(soundPath); !os.IsNotExist(err) {
		c.Sounds.ClickSound1 = rl.LoadSound(soundPath)
	}

	// Initialize settings with defaults
	c.SelectedTheme = 0
	c.SelectedWordList = 0
	c.TestSettings = config.TestSettings{
		TestModeAmounts: []int{15, 30, 60, 120},
		SelectedAmount:  1,
		TestMode:        config.TEST_MODE_TIME,
	}
	c.CursorStyle = config.CURSOR_BLOCK
	c.SoundOn = true
}

func (c *Context) loadWordLists(basePath string) {
	// Default word list if no files found
	defaultWords := []string{"the", "be", "to", "of", "and", "a", "in", "that",
		"have", "I", "it", "for", "not", "on", "with", "he", "as", "you", "do", "at"}
	c.WordLists = []config.WordList{{Name: "Default", Words: defaultWords}}

	if _, err := os.Stat(basePath); os.IsNotExist(err) {
		return
	}

	files, err := os.ReadDir(basePath)
	if err != nil {
		return
	}

	for _, file := range files {
		if !file.IsDir() && strings.HasSuffix(file.Name(), ".txt") {
			words := c.loadWordsFromFile(filepath.Join(basePath, file.Name()))
			if len(words) > 0 {
				name := strings.TrimSuffix(file.Name(), ".txt")
				name = strings.ReplaceAll(name, "_", " ")
				if len(name) > 0 {
					name = strings.ToUpper(string(name[0])) + name[1:]
				}
				c.WordLists = append(c.WordLists, config.WordList{Name: name, Words: words})
			}
		}
	}
}

func (c *Context) loadWordsFromFile(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		return nil
	}
	defer file.Close()

	var words []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line != "" {
			words = append(words, line)
		}
	}
	return words
}

func (c *Context) SaveSettings() {
	// TODO
}

func (c *Context) Unload() {
	rl.UnloadFont(c.Fonts.TypingTestFont.Font)
	rl.UnloadFont(c.Fonts.TitleFont.Font)
	rl.UnloadFont(c.Fonts.TinyFont.Font)
	rl.UnloadFont(c.Fonts.BigFont.Font)
	rl.UnloadSound(c.Sounds.ClickSound1)
}
