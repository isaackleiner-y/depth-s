package main

import "github.com/gdamore/tcell/v2"

type Sprite struct {
	Char  rune
	X, Y  int
	Color tcell.Color
}

func NewSprite(char rune, x, y int) *Sprite {
	return &Sprite{
		Char:  char,
		X:     x,
		Y:     y,
		Color: tcell.ColorWhite,
	}
}

func (s *Sprite) Draw(screen tcell.Screen) {
	style := tcell.StyleDefault.Foreground(s.Color)
	cx, cy := toGameCoordinate(screen, s.X, s.Y)
	screen.SetContent(cx, cy, s.Char, nil, style)
}
