package main

import (
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

func debugOverlay(screen *ebiten.Image, g *Game) {
	// Showing whether a touch is currently happening or not.
	isTouched := GetTouches()
	if isTouched!= nil{
		ebitenutil.DebugPrint(screen, "TOUCH")
	} else {
		ebitenutil.DebugPrint(screen, "NOT TOUCHED")
	} 
	//printing the current mouse information.
	ebitenutil.DebugPrintAt(screen, strconv.FormatBool(g.Mouse.IsDown) + " "+strconv.FormatInt(int64(g.Mouse.XPosition), 10) + " "+strconv.FormatInt(int64(g.Mouse.YPosition), 10), 0, 20 )
	//printing the current zoom level.
	ebitenutil.DebugPrintAt(screen, "ZOOM: "+strconv.FormatFloat(g.Camera.zoom, 'f', 2, 64), 0, 40 )
    //printing the current stage.
    ebitenutil.DebugPrintAt(screen, "STAGE: "+g.Stage, 0, 60 )
	//printing the current date.
	ebitenutil.DebugPrintAt(screen, "DATE: "+strconv.FormatInt(int64(g.World.Date.Year), 10) + "-" + strconv.FormatInt(int64(g.World.Date.Month), 10) + "-" + strconv.FormatInt(int64(g.World.Date.Day), 10) + " " + strconv.FormatInt(int64(g.World.Date.Hour), 10)+ ":" + strconv.FormatInt(int64(g.World.Date.Minute), 10)+ ":" +strconv.FormatInt(int64(g.World.Date.Second), 10), 0, 80)
	//printing the current replay time.
	ebitenutil.DebugPrintAt(screen, "REPLAY TIME: "+strconv.FormatFloat(g.World.ReplayTime, 'f', 2, 64), 0, 100 )

	ebitenutil.DebugPrintAt(screen, "FPS: "+strconv.FormatFloat(ebiten.ActualFPS(), 'f', 2, 64), 0, 140 )
	//print g.IsMouseDragging
	ebitenutil.DebugPrintAt(screen, strconv.FormatBool(g.IsMouseDragging), 0, 160 )
    //print g.MouseClickPosition
    ebitenutil.DebugPrintAt(screen, "CLICK POSITION: "+strconv.FormatInt(int64(g.MouseClickPosition.x), 10) + " "+strconv.FormatInt(int64(g.MouseClickPosition.y), 10), 0, 180 )
	//print drawn tiles
	ebitenutil.DebugPrintAt(screen, "Drawn Tiles: "+strconv.FormatInt(int64(g.TilesDrawn), 10), 0, 200 )
}