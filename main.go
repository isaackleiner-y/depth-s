package main

import (
	"log"

	"github.com/gdamore/tcell/v2"
)

const (
	appWidth  = gameWidth + characterWindowWidth + barsWindowWidth
	appHeight = gameHeight

	gameWidth  = 80
	gameHeight = 30

	characterWindowWidth  = 17
	characterWindowHeigth = 5

	barsWindowWidth  = 14
	barsWindowHeight = 5
)

func main() {
	screen, err := tcell.NewScreen()
	if err != nil {
		log.Fatal(err)
	}

	defer screen.Fini()
	screen.Init()

	// game init section
	player := NewSprite('@', gameWidth/2, gameHeight/2)
	player.Color = tcell.Color(tcell.ColorGold)

	score := 0
	level := 1
	hp := 100
	stamina := 100
	exp := 0
	playerName := "Isaac"

	// game loop
	run := true
	for run {
		sw, sh := screen.Size()
		offsetX := (sw - appWidth) / 2
		offsetY := (sh - appHeight) / 2
		charOffsetX := offsetX + gameWidth + 1
		charOffsetY := offsetY
		barOffsetX := offsetX + gameWidth + characterWindowWidth
		barOffsetY := offsetY

		screen.Clear()
		drawFrame(screen)
		player.Draw(screen)

		// ui
		// drawString(screen, 1, 1, fmt.Sprintf("Score: %d", score))
		// drawString(screen, 1, 2, fmt.Sprintf("Level: %d", level))

		// окно с игрой
		titleWindow := "Depth's"
		drawString(screen, offsetX+(gameWidth/2)-(len(titleWindow)/2), offsetY, titleWindow)

		// окно с игроком
		CharacterWindow := "Character"
		drawString(screen, charOffsetX+3, charOffsetY, CharacterWindow)
		//вывод значений
		drawString(screen, charOffsetX+1, charOffsetY+1, "Name: ")
		drawString(screen, charOffsetX+8, charOffsetY+1, playerName)

		drawString(screen, charOffsetX+1, charOffsetY+2, "Score: ")
		drawCharacterInfo(screen, charOffsetX+8, charOffsetY+2, score)

		drawString(screen, charOffsetX+1, charOffsetY+3, "Level: ")
		drawCharacterInfo(screen, charOffsetX+8, charOffsetY+3, level)

		//окно с барами
		BarsWindow := "Bars"
		drawString(screen, barOffsetX+5, barOffsetY, BarsWindow)
		//вывод значений hp, stn, exp
		drawString(screen, barOffsetX+2, barOffsetY+1, "HP: ")
		drawCharacterInfo(screen, barOffsetX+7, barOffsetY+1, hp)

		drawString(screen, barOffsetX+2, barOffsetY+2, "STM: ")
		drawCharacterInfo(screen, barOffsetX+7, barOffsetY+2, stamina)

		drawString(screen, barOffsetX+2, barOffsetY+3, "EXP: ")
		drawCharacterInfo(screen, barOffsetX+7, barOffsetY+3, exp)

		screen.Show()

		// Управление
		ev := screen.PollEvent()
		// проверили тип события
		switch ev := ev.(type) {
		case *tcell.EventKey:
			// проверили клавишу события
			switch ev.Rune() {
			case 'q':
				run = false
			case 'w':
				if player.Y > 1 {
					player.Y--
				}
			case 'a':
				if player.X > 1 {
					player.X--
				}
			case 's':
				if player.Y < gameHeight-2 {
					player.Y++
				}
			case 'd':
				if player.X < gameWidth-2 {
					player.X++
				}
			}
		}
	}
}
