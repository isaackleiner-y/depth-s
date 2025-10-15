package main

import (
	"strconv"

	"github.com/gdamore/tcell/v2"
)

// отцентровка
func toGameCoordinate(screen tcell.Screen, x, y int) (int, int) {
	sw, sh := screen.Size()
	offsetX := (sw - appWidth) / 2
	offsetY := (sh - appHeight) / 2
	return x + offsetX, y + offsetY
}

// вывод текста
func drawString(screen tcell.Screen, x, y int, msg string) {
	for index, char := range msg {
		screen.SetContent(x+index, y, char, nil, tcell.StyleDefault)
	}
}

func drawCharacterInfo(screen tcell.Screen, x, y int, msg int) {
	str := strconv.Itoa(msg)
	for index, char := range str {
		screen.SetContent(x+index, y, char, nil, tcell.StyleDefault)
	}
}

func drawFrame(screen tcell.Screen) {
	grayBorder := tcell.StyleDefault.Foreground(tcell.ColorGray)

	sw, sh := screen.Size()
	offsetX := (sw - appWidth) / 2
	offsetY := (sh - appHeight) / 2

	// окно игры
	for x := 0; x < gameWidth; x++ {
		screen.SetContent(offsetX+x, offsetY, '─', nil, grayBorder)
		screen.SetContent(offsetX+x, offsetY+gameHeight-1, '─', nil, grayBorder)
	}

	for y := 0; y < gameHeight; y++ {
		screen.SetContent(offsetX, offsetY+y, '│', nil, grayBorder)
		screen.SetContent(offsetX+gameWidth-1, offsetY+y, '│', nil, grayBorder)
	}

	screen.SetContent(offsetX, offsetY, '┌', nil, grayBorder)
	screen.SetContent(offsetX+gameWidth-1, offsetY, '┐', nil, grayBorder)
	screen.SetContent(offsetX, offsetY+gameHeight-1, '└', nil, grayBorder)
	screen.SetContent(offsetX+gameWidth-1, offsetY+gameHeight-1, '┘', nil, grayBorder)

	// окна статистики
	charOffsetX := offsetX + gameWidth + 1
	charOffsetY := offsetY

	for x := 0; x < characterWindowWidth; x++ {
		screen.SetContent(charOffsetX+x, charOffsetY, '─', nil, grayBorder)
		screen.SetContent(charOffsetX+x, charOffsetY+characterWindowHeigth-1, '─', nil, grayBorder)
	}

	for y := 0; y < characterWindowHeigth; y++ {
		screen.SetContent(charOffsetX, charOffsetY+y, '│', nil, grayBorder)
		screen.SetContent(charOffsetX+characterWindowWidth-1, charOffsetY+y, '│', nil, grayBorder)
	}

	screen.SetContent(charOffsetX, charOffsetY, '┌', nil, grayBorder)
	screen.SetContent(charOffsetX+characterWindowWidth-1, charOffsetY, '┐', nil, grayBorder)
	screen.SetContent(charOffsetX, charOffsetY+characterWindowHeigth-1, '└', nil, grayBorder)
	screen.SetContent(charOffsetX+characterWindowWidth-1, charOffsetY+characterWindowHeigth-1, '┘', nil, grayBorder)
}
