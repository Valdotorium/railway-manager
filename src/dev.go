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
	ebitenutil.DebugPrintAt(screen, strconv.FormatBool(g.Mouse.isDown) + " "+strconv.FormatInt(int64(g.Mouse.position.x), 10) + " "+strconv.FormatInt(int64(g.Mouse.position.y), 10), 0, 40 )

}