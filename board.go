package main

import "github.com/veandco/go-sdl2/sdl"

type board struct {
	w, h  int
	x, y  int
	scale int
}

func NewBoard(w, h, x, y, scale int) *board {
	return &board{
		w:     w,
		h:     h,
		x:     x,
		y:     y,
		scale: scale,
	}
}

func (b *board) Draw(r *sdl.Renderer) {
	r.SetDrawColor(255, 255, 255, 255)
	r.DrawRect(&sdl.Rect{
		X: int32(b.x),
		Y: int32(b.y),
		W: int32(b.w),
		H: int32(b.h),
	})
}
