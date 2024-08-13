package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)
type Tile struct {
	Type string
	Population int
	Texture *ebiten.Image
}
type World struct {
	Size int
	Tiles [][]Tile
}

func generateWorld(Textures map[string]*ebiten.Image) World{
	//make an empty ocean world, 100x100 tiles
	world := World{
        Size: 100,
        Tiles: make([][]Tile, 100),
    }
    for i := 0; i < world.Size; i++ {
        world.Tiles[i] = make([]Tile, 100)
        for j := 0; j < world.Size; j++ {
            world.Tiles[i][j] = Tile{
                Type: "water",
                Population: 0,
				Texture: Textures["assets/water.png"],
            }
        }
    }
    return world
}