package editor

import (
	"log"

	"github.com/mattn/go-runewidth"
	"github.com/nsf/termbox-go"
)

func printMessage(col, row int, fg, bg termbox.Attribute, message string) {
	for _, ch := range message {
		termbox.SetCell(col, row, ch, fg, bg)
		col += runewidth.RuneWidth(ch)
	}
}

func Run() {
	err := termbox.Init()
	if err != nil {
		log.Fatal(err)
	}
	defer termbox.Close()

	width, height := termbox.Size()
	message := "gim - a text editor in go"
	messageWidth := runewidth.StringWidth(message)
	col := (width - messageWidth) / 2
	row := height / 2

	printMessage(col, row, termbox.ColorDefault, termbox.ColorDefault, message)
	termbox.Flush()

	termbox.PollEvent()
}
