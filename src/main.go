package main

import (
	"fmt"
	"image/color"
	"log"

	"github.com/Valdotorium/gobird/pkg/button"
	"github.com/hajimehoshi/ebiten/v2"
)

//fetching the paths of the games images from constants.go
var imagePaths []string = fetchGameImagePaths()

type Game struct{
	//go dict
	Textures map[string]*ebiten.Image
	IsDebuggingMode bool
	Score int
	Mouse Mouse
	Stage string
	Zoom float32
	Button button.Button
}
func NewGame() *Game {
	return &Game{
        Textures: LoadImages(imagePaths),
		IsDebuggingMode: true,
		Score: 0,
		Mouse: Mouse{},
		Stage: "game",
        Zoom: 1.0,
		Button: button.Button{
			XPos: 100,
			YPos: 100,
			Width: 20,
			Height: 20,
			Text: "Hi",
			TextColor: color.RGBA{200,200,200,255}, 
			ButtonColor: color.RGBA{20,20,20,255}, 
			HoveredColor: color.RGBA{50,50,50,255},
		},
    }
}
func (g *Game) Update() error {
	//detecting mouse clicks, touch and getting the cursor position
	g = UpdateMouse(g)
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	//clearing the screen
	screen.Fill(color.RGBA{100,120,180,255})
	if DEBUG_OVERLAY {
		debugOverlay(screen, g)
		g.Button.Update(screen)
	}

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return WIN_WIDTH, WIN_HEIGHT
}

func main() {
	fmt.Println("fetched image paths: ", imagePaths)
	g := NewGame()
	g.Button.Init()
	ebiten.SetWindowSize(WIN_WIDTH, WIN_HEIGHT)
	ebiten.SetWindowTitle("Hello, World!")
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeOnlyFullscreenEnabled)
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}