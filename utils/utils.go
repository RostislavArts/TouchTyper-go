package utils

import (
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func GetCenter(width, height int32) rl.Vector2 {
	return rl.Vector2{X: float32(width) / 2.0, Y: float32(height) / 2.0}
}

func SinPulse(frequency float32) float32 {
	return 0.5 * (1 + float32(math.Sin(float64(2*math.Pi*frequency)*rl.GetTime())))
}

func Lerp(a, b, t float32) float32 {
	return a + (b-a)*t
}

func RuneToKeyCode(r rune) int {
	// Map common runes to raylib key codes
	switch r {
	case 'a':
		return int(rl.KeyA)
	case 'b':
		return int(rl.KeyB)
	case 'c':
		return int(rl.KeyC)
	case 'd':
		return int(rl.KeyD)
	case 'e':
		return int(rl.KeyE)
	case 'f':
		return int(rl.KeyF)
	case 'g':
		return int(rl.KeyG)
	case 'h':
		return int(rl.KeyH)
	case 'i':
		return int(rl.KeyI)
	case 'j':
		return int(rl.KeyJ)
	case 'k':
		return int(rl.KeyK)
	case 'l':
		return int(rl.KeyL)
	case 'm':
		return int(rl.KeyM)
	case 'n':
		return int(rl.KeyN)
	case 'o':
		return int(rl.KeyO)
	case 'p':
		return int(rl.KeyP)
	case 'q':
		return int(rl.KeyQ)
	case 'r':
		return int(rl.KeyR)
	case 's':
		return int(rl.KeyS)
	case 't':
		return int(rl.KeyT)
	case 'u':
		return int(rl.KeyU)
	case 'v':
		return int(rl.KeyV)
	case 'w':
		return int(rl.KeyW)
	case 'x':
		return int(rl.KeyX)
	case 'y':
		return int(rl.KeyY)
	case 'z':
		return int(rl.KeyZ)
	case '[':
		return int(rl.KeyLeftBracket)
	case ']':
		return int(rl.KeyRightBracket)
	case ';':
		return int(rl.KeySemicolon)
	case '\'':
		return int(rl.KeyApostrophe)
	case ',':
		return int(rl.KeyComma)
	case '.':
		return int(rl.KeyPeriod)
	case '/':
		return int(rl.KeySlash)
	default:
		return -1
	}
}
