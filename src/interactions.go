package main

import (
	"github.com/Valdotorium/gobird/pkg/touch"
	"github.com/hajimehoshi/ebiten/v2"
)
type Touch struct{
	position Vector2i
	release bool
	press bool
}

type Mouse struct {
	isDown bool
	position Vector2i
}

func GetTouches()*Touch{
	//TODO: #1 ,implement this to the template
	touch.UpdateTouchIDs()
	touches := touch.GetTouchIDs()
	for i := range touches{
		touchposx, touchposy := ebiten.TouchPosition(touches[i])
		//if a touch is happening, the function returns the first touch in touches
		return &Touch{
			position :  Vector2i{x:touchposx, y:touchposy},
			press : touch.IsTouchJustPressed(touches[i]),
			release : touch.IsTouchJustReleased(touches[i])}
	}
	return nil
}

func UpdateMouse(g *Game)*Game{
	touched := GetTouches()
	//touches or left mouse button presses set this to true
    if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) || touched != nil {
        g.Mouse.isDown = true
    } else {
		g.Mouse.isDown = false
	}
	mouseposx, mouseposy := ebiten.CursorPosition()
	if touched != nil{
		g.Mouse.position = touched.position
	} else {
		g.Mouse.position = Vector2i{x:mouseposx, y:mouseposy}
	}

    return g
}