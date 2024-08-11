package button

import (
	"bytes"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
)

type Button struct {
    XPos int
	YPos int
	Width int
	Height int
	Text string
	TextColor color.Color
	ButtonColor color.Color
	HoveredColor color.Color
	//written in init(), must be called
	textFace text.GoTextFace
	//text draw options
	options text.DrawOptions
	//the rendered text (rendered in init()) on a image and its options
	textImage *ebiten.Image
	textImageOptions *ebiten.DrawImageOptions
	//the current state of the button (pressed 2, released 0, hovered 1)
	State int
}
func (b *Button) PreRenderText(){
	var textImage *ebiten.Image= ebiten.NewImage(b.Width, b.Height)
	text.Draw(textImage, b.Text, &b.textFace, &b.options)
	//the prefab image with the rendered text for the button
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(b.XPos), float64(b.YPos))
	b.textImageOptions = op
	b.textImage = textImage
}
func (b *Button) Init(){
	//initializes the texts draw options and text face, will panic if not run
	s, err := text.NewGoTextFaceSource(bytes.NewReader(fonts.MPlus1pRegular_ttf))
	if err != nil {
		panic(err)
	}
	mplusFaceSource := s
	b.textFace = text.GoTextFace{
		Size: float64(b.Height - 8),
		//loaded ttf file from github.com/hajimehoshi/ebiten/v2/examples/resources/fonts
		Source: mplusFaceSource,
	}
	b.options = text.DrawOptions{}
	//setting the text color through the ebiten.DrawImageOptions of b.options
	b.options.ColorScale.ScaleWithColor(b.TextColor)
	//get prerendered text image to put on screen
	b.PreRenderText()
}

func (b *Button) GetState(screen *ebiten.Image) (clickState, hoverState bool){
	mx, my := ebiten.CursorPosition()

    // Check if the mouse is hovering over the button
    isHovered := mx >= b.XPos && mx < b.XPos+b.Width && my >= b.YPos && my < b.YPos+b.Height

    // Update the button color based on hover state
    isClicked := isHovered && ebiten.IsMouseButtonPressed(ebiten.MouseButton0)

	return isClicked, isHovered
}

func (b *Button) Update(screen *ebiten.Image){
	var isClicked, isHovered bool = b.GetState(screen)
	// Drawing the button background
	if !isHovered{
		vector.DrawFilledRect(screen,
			float32(b.XPos),
			float32(b.YPos),
			float32(b.Width),
			float32(b.Height),
			b.ButtonColor, false)
	} else {
		vector.DrawFilledRect(screen,
            float32(b.XPos),
            float32(b.YPos),
            float32(b.Width),
            float32(b.Height),
            b.HoveredColor, false)
	}

    // Drawing the button text
	screen.DrawImage(b.textImage, b.textImageOptions)

	if isClicked {
		b.State = 2
	} else if isHovered {
		b.State = 1
	} else {
		b.State = 0
	}


}
