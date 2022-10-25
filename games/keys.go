package games

import (
	"fmt"
	"github.com/nsf/termbox-go"
)

const (
	ESC uint16 = uint16(termbox.KeyEsc)
	SPACE uint16 = uint16(termbox.KeySpace)
)

type KeyboardEventHandler func(uint16)
var keyboardEventHandler KeyboardEventHandler = nil

func RegisterKeyboardHandler(keyboardInHandler KeyboardEventHandler) {
	keyboardEventHandler = keyboardInHandler
}

func StartCaptureKeyboard() {
	if keyboardEventHandler == nil {
		fmt.Println("You should register keyboardEventHandler")
	}
	for {
		ev := termbox.PollEvent();
		if ev.Type == termbox.EventKey && keyboardEventHandler != nil {
			keyboardEventHandler(uint16(ev.Key))
		}
	}
}