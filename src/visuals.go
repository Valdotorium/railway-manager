package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"fmt"
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
	//all tile images must be 32x32 pixels
	op := &ebiten.DrawImageOptions{}
	g.TilesDrawn = 0
	op.GeoM.Scale(float64(g.Camera.zoom), float64(g.Camera.zoom))  //scaling the tilemap by camera zoom level

	var x float64 = float64(g.Camera.position.x)
	var y float64 = float64(g.Camera.position.y)
	//offsetting the tilemap by camera position
	var zoom float64 = float64(g.Camera.zoom)
	op.GeoM.Translate(float64(x) * zoom, float64(y) * zoom)
	
	for i := 0; i < len(g.World.Tiles); i++{
		list := g.World.Tiles[i]
		//only draw tiles within the screen boundaries

		for j := 0; j < len(list); j++{
            tile := list[j]

			//only draw visible tiles
			CurrentTileX := (float64(j) * float64(g.TileSize) + x) * zoom
			CurrentTileY := (float64(i) * float64(g.TileSize) + y) * zoom 
			if -float64(g.TileSize) * zoom< CurrentTileX && CurrentTileX < float64(WIN_WIDTH){
				fmt.Println(CurrentTileX , CurrentTileX + float64(WIN_WIDTH) * zoom)
				if -float64(g.TileSize) * zoom < CurrentTileY && CurrentTileY < float64(WIN_HEIGHT){
                    screen.DrawImage(tile.Texture, op)
					g.TilesDrawn++
                }
			}
			op.GeoM.Translate(float64(g.TileSize) * zoom,0)
        }
		op.GeoM.Translate(-float64(g.TileSize * g.World.Size) * zoom,float64(g.TileSize) * zoom)
	}
}