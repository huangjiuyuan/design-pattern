package game

import (
	"fmt"
)

type Game struct {
	Initialize func()
	StartPlay  func()
	EndPlay    func()
}

func (g *Game) Play() {
	g.Initialize()
	g.StartPlay()
	g.EndPlay()
}

type Cricket struct {
	Game
}

func NewCricket() *Cricket {
	return &Cricket{
		Game: Game{
			Initialize: func() {
				fmt.Println("Cricket Game Initialized! Start playing.")
			},
			StartPlay: func() {
				fmt.Println("Cricket Game Started. Enjoy the game!")
			},
			EndPlay: func() {
				fmt.Println("Cricket Game Finished!")
			},
		},
	}
}

type Football struct {
	Game
}

func NewFootball() *Football {
	return &Football{
		Game: Game{
			Initialize: func() {
				fmt.Println("Football Game Initialized! Start playing.")
			},
			StartPlay: func() {
				fmt.Println("Football Game Started. Enjoy the game!")
			},
			EndPlay: func() {
				fmt.Println("Football Game Finished!")
			},
		},
	}
}
