package main

import (
	"github.com/hajimehoshi/ebiten/v2"
    "github.com/Valdotorium/gobird/pkg/time"
)
type Tile struct {
	Type string
	Population int
	Texture *ebiten.Image
}
type World struct {
	Size int
	Tiles [][]Tile
    Date time.Date;
}

func generateWorld(Textures map[string]*ebiten.Image) World{
	//make an empty ocean world, 100x100 tiles
	world := World{
        Size: 100,
        Tiles: make([][]Tile, 100),
        Date: time.Date{Year: 2000, Month: 1, Day: 1, Daytime: time.Time{Hour: 7, Minute: 0, Second: 0}},
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