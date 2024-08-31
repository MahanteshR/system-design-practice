package model

import "fmt"

type Game struct {
	Board   *Board
	Players []*Player
	Dice    *Dice
}

func (g *Game) MovePlayer(player *Player) {
	roll := g.Dice.RollDice()

	newPos := player.Position + roll

	if newPos > g.Board.Size {
		newPos = g.Board.Size
	}

	if finalPos, ok := g.Board.Snakes[newPos]; ok {
		fmt.Printf("player %s bit by snake from pos:%d to pos:%d\n", player.Name,
			player.Position, finalPos)

		newPos = finalPos
	} else if finalPos, ok := g.Board.Ladders[newPos]; ok {
		fmt.Printf("player %s climbed the ladder from pos:%d to pos:%d\n", player.Name,
			player.Position, finalPos)

		newPos = finalPos
	}

	player.Position = newPos
	fmt.Printf("player %s currently at pos:%d\n", player.Name, player.Position)
}

func (g *Game) HasPlayerWon(player *Player) bool {
	return player.Position == g.Board.Size
}

func (g *Game) Play() {
	for {
		for _, player := range g.Players {
			g.MovePlayer(player)

			if g.HasPlayerWon(player) {
				fmt.Printf("Player %s won the game\n", player.Name)

				return
			}
		}
	}
}
