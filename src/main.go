package main

import (
	"fmt"
	"image/color"
	"log"

	"github.com/Valdotorium/gobird/pkg/button"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/Valdotorium/gobird/pkg/touch"
)

//fetching the paths of the games images from constants.go
var imagePaths []string = fetchGameImagePaths()

type Game struct{
	//go dict
	//image related
	Textures map[string]*ebiten.Image
	//mouse related
	Mouse touch.Mouse
	IsMouseDragging bool
	MouseClickPosition Vector2i
	//other
	Stage string
	Camera *Camera
	Button button.Button
	TilesDrawn int
	World World
	TileSize int
	IsDebuggingMode bool
	Score int
}
func NewGame() *Game {
	return &Game{
        Textures: LoadImages(imagePaths),
		IsDebuggingMode: true,
		Score: 0,
		Mouse: touch.Mouse{},
		Stage: "menu",
        Camera: NewCamera(),
		TileSize: 32,
		Button: button.Button{
			XPos: 100,
			YPos: 100,
			Width: 90,
			Height: 24,
			Text: "Start Game",
			TextColor: color.RGBA{200,200,200,255}, 
			ButtonColor: color.RGBA{20,20,20,255}, 
			HoveredColor: color.RGBA{50,50,50,255},
		},
		TilesDrawn: 0,
    }
}
func (g *Game) Update() error {
	//detecting mouse clicks, touch and getting the cursor position
	UpdateMouse(g)
	g.World.SimulateWorld()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	//clearing the screen
	screen.Fill(color.RGBA{100,120,180,255})
	if g.Stage == "game"{
		g.MoveCamera()
		g.drawTilemap(screen)
	} else if g.Stage == "menu" {
		g.Button.Update(screen, &g.Mouse)
		g.UpdateMenu(screen)
	}
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
	g.Button.Init()
	g.World = generateWorld(g.Textures)
	ebiten.SetWindowSize(WIN_WIDTH, WIN_HEIGHT)
	ebiten.SetWindowTitle("Hello, World!")
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeOnlyFullscreenEnabled)
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}