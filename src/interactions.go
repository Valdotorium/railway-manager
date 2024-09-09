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

    //set oldx and y to x and y pos before update
    g.Mouse.OldXPosition = g.Mouse.XPosition
    g.Mouse.OldYPosition = g.Mouse.YPosition
    
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
    if g.Mouse.IsDown && g.Mouse.Ticks == 1{
        g.MouseClickPosition = Vector2i{x: mouseposx,  y:mouseposy}
    }
    if g.Mouse.IsDown {
        g.Mouse.Ticks++
    } else {
        g.Mouse.Ticks = 0
    }


    //dragging
    if g.Mouse.Ticks >= 12 {
        g.IsMouseDragging = true
    } else {
        g.IsMouseDragging = false
    }

}

func (g *Game) MoveCamera() {

    //dragging to move the map, by changing the cameras position by the movement of the mouse in the last frame
    if g.IsMouseDragging {
        g.Camera.position.x -= int(float64(g.Mouse.OldXPosition - g.Mouse.XPosition) * g.Camera.zoom)
        g.Camera.position.y -= int(float64(g.Mouse.OldYPosition - g.Mouse.YPosition) * g.Camera.zoom)
    }
    //clamping camera position to prevent it from going off the screen
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
	//using up and down arrow, adjust ReplayTime
	if ebiten.IsKeyPressed(ebiten.KeyUp) {
        g.World.ReplayTime += 0.05
    }
    if ebiten.IsKeyPressed(ebiten.KeyDown) {
        g.World.ReplayTime -= 0.05
    }
	//clamping ReplayTime, so that it is not negative or too large
	if g.World.ReplayTime < 0.2 {
        g.World.ReplayTime = 0.2
    }
    if g.World.ReplayTime > 8 {
        g.World.ReplayTime = 8
    }
	//clamping zoom to prevent zooming out too far
	if g.Camera.zoom < MIN_ZOOM {
        g.Camera.zoom = MIN_ZOOM
    }
    if g.Camera.zoom > MAX_ZOOM {
        g.Camera.zoom = MAX_ZOOM
    }
}