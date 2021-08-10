package main

import (
	"fmt"
	"math/rand"

	"github.com/veandco/go-sdl2/sdl"
)

type fruit struct {
	board *board
	pos   point
	color color
}

func NewFruit(b *board) *fruit {
	f := &fruit{
		board: b,
		color: color{0, 255, 0, 255},
	}

	f.Update()

	return f
}

func (f *fruit) Update() {
	f.pos = point{
		x: int32(rand.Intn((f.board.w+f.board.x)/f.board.scale) + 1),
		y: int32(rand.Intn((f.board.h+f.board.y-20)/f.board.scale) + 2),
	}

	fmt.Printf("Fruit pos: %v\n", f.pos)
}

func (f *fruit) Draw(r *sdl.Renderer) {
	var rect sdl.Rect
	r.SetDrawColor(f.color.r, f.color.g, f.color.b, f.color.a)
	rect = sdl.Rect{
		X: f.pos.x * int32(f.board.scale),
		Y: f.pos.y * int32(f.board.scale),
		W: int32(f.board.scale),
		H: int32(f.board.scale),
	}
	r.FillRect(&rect)
}
