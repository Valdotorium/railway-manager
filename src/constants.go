//this file stores constant values, such as file paths that need to be loaded
package main
//can be used to describe positions in 2d space with integers
type Vector2i struct {
	x int
	y int
}
//can be used to describe positions in 2d space with floats
type Vector2 struct {
	x float64
	y float64
}
//dimensions of the window
var WIN_WIDTH int = 720
var WIN_HEIGHT int = 480
//size of one tile with zoom = 1
var TILE_SIZE int = 32
//limits of the zoom variable
var MIN_ZOOM float64 = 0.5
var MAX_ZOOM float64 = 2.5
//settings
//debugging texts / information on screen
var DEBUG_OVERLAY bool = true
