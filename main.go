package main

import (
	"log"
	"math/rand"

	"github.com/gdamore/tcell/v2"
)

const (
	appWidth  = gameWidth + characterWindowWidth
	appHeight = gameHeight

	gameWidth  = 80
	gameHeight = 30

	characterWindowWidth  = 17
	characterWindowHeigth = 5
)

// создание монеток
func setupCoins(level int) []*Sprite {
	coins := make([]*Sprite, level+2)
	for index := 0; index < level+2; index++ {
		coins[index] = NewSprite(
			'0',
			rand.Intn(gameWidth-2)+1,
			rand.Intn(gameHeight-2)+1,
		)
	}
	return coins
}

func main() {
	screen, err := tcell.NewScreen()
	if err != nil {
		log.Fatal(err)
	}

	defer screen.Fini()
	screen.Init()

	// game init section
	player := NewSprite('@', gameWidth/2, gameHeight/2)
	coins := setupCoins(1)

	score := 0
	level := 1

	// game loop
	run := true
	for run {
		sw, sh := screen.Size()
		offsetX := (sw - appWidth) / 2
		offsetY := (sh - appHeight) / 2
		charOffsetX := offsetX + gameWidth + 1
		charOffsetY := offsetY

		screen.Clear()
		drawFrame(screen)

		player.Draw(screen)
		for _, coin := range coins {
			coin.Draw(screen)
		}

		// ui
		// drawString(screen, 1, 1, fmt.Sprintf("Score: %d", score))
		// drawString(screen, 1, 2, fmt.Sprintf("Level: %d", level))

		// окно с игрой
		titleWindow := "Depth's"
		drawString(screen, offsetX+(gameWidth/2)-(len(titleWindow)/2), offsetY, titleWindow)

		// окно с игроком
		CharacterWindow := "Character"
		charName := "Bober"
		scoreStr := "Score:"
		levelStr := "Level:"
		drawString(screen, charOffsetX+(len(CharacterWindow)/2), charOffsetY, CharacterWindow)
		drawString(screen, 1+charOffsetX+1, charOffsetY+1, "Name: "+charName)
		drawString(screen, 1+charOffsetX+1, charOffsetY+2, scoreStr)
		drawCharacterInfo(screen, charOffsetX+(len(scoreStr))+2, charOffsetY+2, score)
		drawString(screen, 1+charOffsetX+1, charOffsetY+3, levelStr)
		drawCharacterInfo(screen, charOffsetX+(len(levelStr))+2, charOffsetY+3, level)

		screen.Show()

		// Управление
		playerMoved := false
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
				playerMoved = true
			case 'a':
				if player.X > 1 {
					player.X--
					playerMoved = true
				}
			case 's':
				if player.Y < gameHeight-2 {
					player.Y++
					playerMoved = true
				}
			case 'd':
				if player.X < gameWidth-2 {
					player.X++
					playerMoved = true
				}
			}
		}
		// check coin collision
		if playerMoved {
			coinCollectedIndex := -1
			for index, coin := range coins {
				if coin.X == player.X && coin.Y == player.Y {
					// collect coin
					coinCollectedIndex = index
					score++
				}
			}
			// handle coin collision
			if coinCollectedIndex > -1 {
				coins[coinCollectedIndex] = coins[len(coins)-1]
				coins = coins[0 : len(coins)-1]

				if len(coins) == 0 {
					level++
					coins = setupCoins(level)
				}
			}
		}
	}
}
