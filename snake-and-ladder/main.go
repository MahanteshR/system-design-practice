package main

import "snake-and-ladder/model"

func main() {
	board := &model.Board{
		Size:    100,
		Snakes:  map[int]int{45: 12, 78: 53, 23: 10, 94: 64},
		Ladders: map[int]int{2: 21, 56: 76, 67: 99},
	}

	players := []*model.Player{
		{
			Name: "Mahantesh", Position: 0,
		},
		{
			Name: "Maji", Position: 0,
		},
	}

	dice := &model.Dice{Sides: 6}

	game := &model.Game{
		Board:   board,
		Players: players,
		Dice:    dice,
	}

	game.Play()
}
