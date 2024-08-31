package model

import (
	"math/rand"
	"time"
)

type Dice struct {
	Sides int
}

func (d *Dice) RollDice() int {
	side := rand.New(rand.NewSource(time.Now().UnixNano()))

	return side.Intn(d.Sides) + 1
}
