package main

import "github.com/banthar/Go-SDL/sdl"


type Object struct {
	sdl.Rect

	Color	uint32
}

func (o *Object) Draw(surface *sdl.Surface, gc *GameContext) {
	baseX := int16((gc.PlayerPosition % 641) - 320)

	drawRect := &sdl.Rect{
		X: o.Rect.X - baseX,
		Y: o.Rect.Y,
		W: o.Rect.W,
		H: o.Rect.H,
	}

	surface.FillRect(drawRect, o.Color)
}


var levelObjects []*Object


func inRange(o *Object, ppos int) bool {
	return o.X >= int16(ppos-320) && o.X <= int16(ppos+320)
}


// Returns all level objects which are in range of the player, that is
// are in visibility interval [playerpos-320, playerpos+320].
func findLevelObjects(playerPosition int) []*Object {
	foundObjects := make([]*Object, 0, len(levelObjects))

	for _,e := range levelObjects {
		if inRange(e, playerPosition) {
			foundObjects = append(foundObjects, e)
		}
	}

	return foundObjects
}


func addLevelObject(o *Object) {
	levelObjects = append(levelObjects, o)
}


func loadLevel() {
	levelObjects = make([]*Object, 0, 10)

	addLevelObject(&Object{
		sdl.Rect: sdl.Rect{
			X: 100,
			Y: 70,
			W: 40,
			H: 40,
		},
		Color: 0x00ff00,
	})

	addLevelObject(&Object{
		sdl.Rect: sdl.Rect{
			X: 300,
			Y: 120,
			W: 40,
			H: 40,
		},
		Color: 0x00f000,
	})

	addLevelObject(&Object{
		sdl.Rect: sdl.Rect{
			X: 800,
			Y: 120,
			W: 40,
			H: 40,
		},
		Color: 0xff0000,
	})
}
