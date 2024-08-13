package main

import (
	"github.com/Valdotorium/gobird/pkg/touch"
	"github.com/hajimehoshi/ebiten/v2"
)


func GetTouches()*touch.Touch{
	//TODO: #1 ,implement this to the template
	touch.UpdateTouchIDs()
	touches := touch.GetTouchIDs()
	for i := range touches{
		touchposx, touchposy := ebiten.TouchPosition(touches[i])
		//if a touch is happening, the function returns the first touch in touches
		return &touch.Touch{
			XPosition : touchposx,
			YPosition : touchposy,
			Press : touch.IsTouchJustPressed(touches[i]),
			Release : touch.IsTouchJustReleased(touches[i])}
	}
	return nil
}

func UpdateMouse(g *Game){
	touched := GetTouches()
	//touches or left mouse button presses set this to true
    if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) || touched != nil {
        g.Mouse.IsDown = true
    } else {
		g.Mouse.IsDown = false
	}
	mouseposx, mouseposy := ebiten.CursorPosition()
	if touched != nil{
		g.Mouse.XPosition = touched.XPosition
		g.Mouse.YPosition = touched.YPosition
	} else {
		g.Mouse.XPosition = mouseposx
        g.Mouse.YPosition = mouseposy 
	}

}

func (g *Game) MoveCamera() {
	//simple WASD movement
	if ebiten.IsKeyPressed(ebiten.KeyW) {
        g.Camera.position.y += 10
    }
    if ebiten.IsKeyPressed(ebiten.KeyS) {
        g.Camera.position.y -= 10
    }
    if ebiten.IsKeyPressed(ebiten.KeyA) {
        g.Camera.position.x += 10
    }
    if ebiten.IsKeyPressed(ebiten.KeyD) {
        g.Camera.position.x -= 10
    }
    //zooming in and out
    if ebiten.IsKeyPressed(ebiten.KeyQ) {
        g.Camera.zoom += 0.02
    }
    if ebiten.IsKeyPressed(ebiten.KeyE) {
        g.Camera.zoom -= 0.02
	}
	//clamping zoom to prevent zooming out too far
	if g.Camera.zoom < MIN_ZOOM {
        g.Camera.zoom = MIN_ZOOM
    }
    if g.Camera.zoom > MAX_ZOOM {
        g.Camera.zoom = MAX_ZOOM
    }
}