package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Camera struct{
	position Vector2i
	zoom float64
}

func NewCamera() *Camera {
	return &Camera{
        position: Vector2i{0, 0},
        zoom: 1.0,
    }
}
func (g *Game) UpdateMenu(screen *ebiten.Image) {
	if g.Button.State == 2{
		g.Stage = "game"
	}    
}
func (g *Game) drawTilemap(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}

	var x int = g.Camera.position.x
	var y int = g.Camera.position.y
	//offsetting the tilemap by camera position
	var zoom int = int(g.Camera.zoom)
	op.GeoM.Translate(float64(x * zoom), float64(y * zoom))
	
	for i := 0; i < len(g.World.Tiles); i++{
		list := g.World.Tiles[i]
		
		for j := 0; j < len(list); j++{
            tile := list[j]
            screen.DrawImage(tile.Texture, op)
			op.GeoM.Translate(float64(g.TileSize * zoom),0)
        }
		op.GeoM.Translate(-float64(g.TileSize * zoom * g.World.Size),float64(g.TileSize * zoom))
	}
}