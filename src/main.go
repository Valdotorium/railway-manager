package main

import (
	"fmt"
	"image/color"
	"log"

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
}
func NewGame() *Game {
	return &Game{
        Textures: LoadImages(imagePaths),
		IsDebuggingMode: true,
		Score: 0,
		Mouse: Mouse{},
		Stage: "game",
        Zoom: 1.0,

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
	}

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return WIN_WIDTH, WIN_HEIGHT
}

func main() {
	fmt.Println("fetched image paths: ", imagePaths)
	g := NewGame()
	ebiten.SetWindowSize(WIN_WIDTH, WIN_HEIGHT)
	ebiten.SetWindowTitle("Hello, World!")
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeOnlyFullscreenEnabled)
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}