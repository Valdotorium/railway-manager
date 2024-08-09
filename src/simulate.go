package main

import (
)
type Tile struct {
	Type string
	Population int
}


type World struct {
	Size int
	Tiles [][]Tile
}