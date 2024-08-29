package main

import (
	"math"

	"github.com/Valdotorium/gobird/pkg/time"
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
    Date time.Date
    ReplayTime float64
    CommuteInterest float64
}

func generateWorld(Textures map[string]*ebiten.Image) World{
	//make an empty ocean world, 100x100 tiles
	world := World{
        Size: 100,
        Tiles: make([][]Tile, 100),
        Date: time.Date{Year: 2000, Month: 1, Day: 1, Hour:7, Minute: 0, Second: 0},
        ReplayTime: 1.0,
        CommuteInterest: 0.3,
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


func (w *World) SimulateWorld(){
    w.UpdateTime()
    w.CalculateCommuteInterest()
}
func (w *World) UpdateTime(){
    w.Date.AddSeconds(int(20 * w.ReplayTime))
}
func (w *World) CalculateCommuteInterest(){
    w.CommuteInterest = 0.3
    //increase commute interest around the times 8am and 5 pm
    if math.Abs(float64(8 - w.Date.Hour)) < 3 {
        w.CommuteInterest = 0.3 +(3 - math.Abs(float64(8 - w.Date.Hour))) / 4
    }
    if math.Abs(float64(17 - w.Date.Hour)) < 3 {
        w.CommuteInterest += 0.3 +(3 - math.Abs(float64(17 - w.Date.Hour))) / 6
    }
    //increase commute interest at noon
    if math.Abs(float64(12 - w.Date.Hour)) < 1 {
        w.CommuteInterest += 0.3+(1 - math.Abs(float64(12 - w.Date.Hour))) / 6
    }
    //decrease commute interest at night
    if w.Date.Hour >= 22 {
        w.CommuteInterest = 0.3-float64(w.Date.Hour - 22) / 8
    }
    if w.Date.Hour <= 6{
        w.CommuteInterest = 0.3-float64(6-w.Date.Hour) / 24
    }
}