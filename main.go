package main

import "fmt"
import "github.com/banthar/Go-SDL/sdl"

var floorTexture *sdl.Surface
var playerTexture *sdl.Surface

func main() {
	if sdl.Init(sdl.INIT_EVERYTHING) != 0 {
		panic(sdl.GetError())
	}

	defer sdl.Quit()

	screen := sdl.SetVideoMode(640, 480, 32, 0)

	if screen == nil {
		panic(sdl.GetError())
	}

	sdl.WM_SetCaption("Ohai","")

	fmt.Println("Launching mainloop")

	loadTextures()

	gameloop(screen)
}



func loadTextures() {
	floorTexture = sdl.Load("floor.png")
	playerTexture = sdl.Load("a.gif")
}




type GameContext struct {
	Screen	*sdl.Surface

	PlayerPosition	int	// Global position (advanced by each step)
	PagePosition	int	// Relative to the page (0-640)
	PlayerSpeed		int
	PageCount		int	// count for every passed page
}


func (gc *GameContext) drawFloor() {
	offset := 320 - (gc.PlayerPosition % 641)

	destRect := &sdl.Rect{
		int16(-320 + offset),
		int16(480 - floorTexture.H),
		0,
		0,
	}

	destRect2 := &sdl.Rect{
		int16(320 + offset),
		int16(480 - floorTexture.H),
		0,
		0,
	}

	gc.Screen.Blit(
		destRect,
		floorTexture,
		nil)

	gc.Screen.Blit(
		destRect2,
		floorTexture,
		nil)
}


func (gc *GameContext) drawPlayer() {
	gc.Screen.FillRect(&sdl.Rect{
			int16(gc.PagePosition),
			int16(floorTexture.H),
			uint16(playerTexture.W),
			uint16(playerTexture.H),
		}, 0xaaaaaa)

	gc.Screen.Blit(
		&sdl.Rect{
			int16(gc.PagePosition),
			int16(floorTexture.H),
			0,
			0,
		},
		playerTexture,
		nil)
}


// Return width of the player in pixel
func (gc *GameContext) PlayerWidth() int {
	return int(playerTexture.W)
}


func (gc *GameContext) moveLeft() {
	gc.PlayerPosition -= gc.PlayerSpeed
	gc.PagePosition -= gc.PlayerSpeed

	if(gc.PagePosition - gc.PlayerWidth() < 0) {
		gc.PagePosition = 0 + gc.PlayerWidth()
	}

	if(gc.PlayerPosition < 0) {
		gc.PlayerPosition = 0
	}
}


func (gc *GameContext) moveRight() {
	gc.PlayerPosition += gc.PlayerSpeed
	gc.PagePosition += gc.PlayerSpeed

	if(gc.PagePosition + gc.PlayerWidth() > 640) {
		gc.PagePosition = 640 - gc.PlayerWidth()
	}

	if(gc.PlayerPosition > 0 && gc.PlayerPosition % 640 == 0) {
		gc.PageCount++
	}
}


func (gc *GameContext) Dump() {
	fmt.Println("PlayerPos:", gc.PlayerPosition, "PagePos:", gc.PagePosition)
}



func gameloop(screen *sdl.Surface) {
	gc := &GameContext{screen, 320, 320, 16, 0}

	for {
		e := sdl.WaitEvent()

		screen.FillRect(nil, 0x0)

		gc.drawFloor()
		gc.drawPlayer()

		switch re := e.(type) {
			case *sdl.QuitEvent:
				return

			case *sdl.MouseMotionEvent:

				screen.FillRect(&sdl.Rect{
					int16(re.X),
					int16(re.Y),
					50, 50}, 0xffffff)

				screen.Blit(&sdl.Rect{
					int16(re.X),
					int16(re.Y),
					0,0}, playerTexture, nil)

				fmt.Println(re.X, re.Y)

			case *sdl.KeyboardEvent:
				if(re.Type == sdl.KEYUP) {
					keyname := sdl.GetKeyName(sdl.Key(re.Keysym.Sym))
					fmt.Println("pressed:", keyname)

					switch keyname {
						case "right": gc.moveRight()
						case "left": gc.moveLeft()
					}
				}
			default:
				//fmt.Println("What the heck?!")
		}

		gc.Dump()

		screen.Flip()
	}
}

