package graphics

import (
	"github.com/nsf/termbox-go"
	"github.com/mattn/go-runewidth"
	"fmt"
)

var Yellow termbox.Attribute = termbox.ColorYellow
var Default termbox.Attribute = termbox.ColorDefault
var Red termbox.Attribute = termbox.ColorRed

func Initialize() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	DrawLine(0, 50, 11)
}

type DrawAdapater func()
type PollEventHandler func(termbox.Event)

func DrawCall(drawAdapater DrawAdapater) {
	err := termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	if err != nil {
		panic(err)
	}
	drawAdapater()
	DrawLine(0, 50, 11)
	termbox.Flush()
}

func Close() {
	termbox.Close()
}

func DrawUI(score int, frame float64) {
	DrawCell(0, 0, termbox.ColorWhite, termbox.ColorDefault, fmt.Sprintf("Score: %v", score))
	DrawCell(0, 1, termbox.ColorWhite, termbox.ColorDefault, fmt.Sprintf("FrameRate: %v", frame))
}

func draw_bg(w, h int, bgColor termbox.Attribute) {
	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			termbox.SetBg(x, y, bgColor)
		}
	}
}

func DrawLine(startAt, w, h int) {
	for x := 0; x < w; x++ {
		termbox.SetCell(x, h, '─', termbox.ColorWhite, termbox.ColorDefault)
	}
}

// This function is often useful:
func DrawCell(x, y int, fg, bg termbox.Attribute, msg string) {
	for _, c := range msg {
		termbox.SetCell(x, y, c, fg, bg)
		x += runewidth.RuneWidth(c)
	}
}