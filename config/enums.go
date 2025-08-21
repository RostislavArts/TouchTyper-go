package config

type CursorStyle int

const (
	CURSOR_BLOCK CursorStyle = iota
	CURSOR_LINE
	CURSOR_UNDERLINE
)

type TestMode int

const (
	TEST_MODE_TIME TestMode = iota
	TEST_MODE_WORDS
)

type Screen int

const (
	SCREEN_TEST Screen = iota
	SCREEN_RESULT
)
