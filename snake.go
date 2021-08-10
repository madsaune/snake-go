package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

type snake struct {
	body      []point
	direction point
	points    int
	color     color
}

func NewSnake() *snake {
	return &snake{
		body: []point{
			{x: 10, y: 10},
			{x: 9, y: 10},
			{x: 8, y: 10},
		},
		direction: point{x: 0, y: 0},
		points:    0,
		color:     color{255, 0, 0, 255},
	}
}

func (s *snake) Update() {
	if len(s.body) > 1 {
		// Remove last element of s.Body and assign to tail. (pop)
		var tail point
		tail, s.body = s.body[len(s.body)-1], s.body[:len(s.body)-1]

		// Set coordinates of tail based on current direction
		if s.direction.x == 1 {
			tail.x = s.body[0].x + 1
			tail.y = s.body[0].y
		} else if s.direction.y == 1 {
			tail.x = s.body[0].x
			tail.y = s.body[0].y + 1
		} else if s.direction.x == -1 {
			tail.x = s.body[0].x - 1
			tail.y = s.body[0].y
		} else if s.direction.y == -1 {
			tail.x = s.body[0].x
			tail.y = s.body[0].y - 1
		}

		// Insert tail as the first element in s.Body (unshift)
		s.body = append([]point{tail}, s.body...)
	} else {
		s.body[0] = point{
			x: s.direction.x,
			y: s.direction.y,
		}
	}
}

func (s *snake) getDirection() int {
	if s.direction.x == 0 && s.direction.y == -1 { // up
		return 1
	} else if s.direction.x == 1 && s.direction.y == 0 { // right
		return 2
	} else if s.direction.x == 0 && s.direction.y == 1 { // down
		return 3
	} else if s.direction.x == -1 && s.direction.y == 0 { // left
		return 4
	} else { // stop
		return 0
	}
}

func (s *snake) ChangeDirection(direction int) {
	if direction == 1 {
		if s.getDirection() != 3 {
			s.direction = point{0, -1}
		}
	} else if direction == 2 {
		if s.getDirection() != 4 {
			s.direction = point{1, 0}
		}
	} else if direction == 3 {
		if s.getDirection() != 1 {
			s.direction = point{0, 1}
		}
	} else if direction == 4 {
		if s.getDirection() != 2 {
			s.direction = point{-1, 0}
		}
	} else if direction == 0 {
		s.direction = point{0, 0}
	}
}

func (s *snake) Reset() {
	s.body = []point{
		{x: 10, y: 10},
		{x: 9, y: 10},
		{x: 8, y: 10},
	}
	s.direction = point{x: 0, y: 0}
	s.points = 0
	s.color = color{255, 0, 0, 255}
}

func (s *snake) Eat(f *fruit) {

	// If snake head is on the same position as fruit
	if s.body[0].x == f.pos.x && s.body[0].y == f.pos.y {
		var new_x int32
		var new_y int32

		// Calculate where to add the new bodypart
		if s.direction.x == 1 {
			new_x = s.body[len(s.body)-1].x + 1
			new_y = s.body[len(s.body)-1].y
		} else if s.direction.x == -1 {
			new_x = s.body[len(s.body)-1].x - 1
			new_y = s.body[len(s.body)-1].y
		} else if s.direction.y == 1 {
			new_x = s.body[len(s.body)-1].x
			new_y = s.body[len(s.body)-1].y + 1
		} else if s.direction.y == -1 {
			new_x = s.body[len(s.body)-1].x
			new_y = s.body[len(s.body)-1].y - 1
		}

		// Append new bodypart to the end of snake
		s.body = append(s.body, point{x: new_x, y: new_y})

		// Increase score
		s.points += 1

		// Place fruit on another random spot
		f.Update()
	}
}

func (s *snake) DetectCollision() bool {
	if s.body[0].x >= BOARD_W/BOARD_SCALE ||
		s.body[0].x < 0 ||
		s.body[0].y >= BOARD_H/BOARD_SCALE ||
		s.body[0].y < 0 {
		return true
	}

	return false
}

func (s *snake) Draw(r *sdl.Renderer) {
	var rects []sdl.Rect

	for _, v := range s.body {
		rects = append(rects, sdl.Rect{
			X: v.x * BOARD_SCALE,
			Y: v.y * BOARD_SCALE,
			W: BOARD_SCALE,
			H: BOARD_SCALE,
		})

	}

	r.SetDrawColor(s.color.r, s.color.g, s.color.b, s.color.a)
	r.FillRects(rects)
}
