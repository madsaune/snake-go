package main

import (
	"math/rand"

	"github.com/veandco/go-sdl2/sdl"
)

type fruit struct {
	board_w, board_h int
	board_scale      int
	pos              point
	color            color
}

func NewFruit(board_w, board_h, board_scale int32) *fruit {
	f := &fruit{}
	f.board_w = int(board_w)
	f.board_h = int(board_h)
	f.board_scale = int(board_scale)
	f.pos = point{
		x: int32(rand.Intn(f.board_w / f.board_scale)),
		y: int32(rand.Intn(f.board_h / f.board_scale)),
	}
	f.color = color{0, 255, 0, 255}

	return f

}

func (f *fruit) Update() {
	f.pos = point{
		x: int32(rand.Intn(f.board_w / f.board_scale)),
		y: int32(rand.Intn(f.board_h / f.board_scale)),
	}
}

func (f *fruit) Draw(r *sdl.Renderer) {
	var rect sdl.Rect
	r.SetDrawColor(f.color.r, f.color.g, f.color.b, f.color.a)
	rect = sdl.Rect{
		X: f.pos.x * BOARD_SCALE,
		Y: f.pos.y * BOARD_SCALE,
		W: BOARD_SCALE,
		H: BOARD_SCALE,
	}
	r.FillRect(&rect)
}
