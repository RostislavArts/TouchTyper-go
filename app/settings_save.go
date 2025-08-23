package app

import "TouchTyper/config"

type saveData struct {
	SelectedTheme    int
	SelectedWordList int
	TestSettings     config.TestSettings
	CursorStyle      config.CursorStyle
	SoundOn          bool
}
