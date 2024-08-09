package button

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Button struct {
    xPos int
	yPos int
	width int
	height int
	text string
	textColor color.Color
	buttonColor color.Color
	hoveredColor color.Color
}


func (b *Button) GetState(screen *ebiten.Image) (clickState, hoverState bool){
	mx, my := ebiten.CursorPosition()

    // Check if the mouse is hovering over the button
    isHovered := mx >= b.xPos && mx < b.xPos+b.width && my >= b.yPos && my < b.yPos+b.height

    // Update the button color based on hover state
    isClicked := isHovered && ebiten.IsMouseButtonPressed(ebiten.MouseButton0)

	return isClicked, isHovered
}

func (b *Button) Update(screen *ebiten.Image) (clickState, hoverState bool){
	var isClicked, isHovered bool = b.GetState(screen)
	// Drawing the button background
	if !isHovered{
		vector.DrawFilledRect(screen,
			float32(b.xPos),
			float32(b.yPos),
			float32(b.width),
			float32(b.height),
			b.buttonColor, false)
	}

    // Drawing the button text
    ebitenutil.DebugPrintAt(screen, b.text, b.xPos, b.yPos)

	return isClicked, isHovered
}
