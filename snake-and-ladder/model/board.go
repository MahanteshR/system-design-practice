package model

type Board struct {
	Size    int
	Snakes  map[int]int
	Ladders map[int]int
}
