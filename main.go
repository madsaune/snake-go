package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

const (
	WIN_W, WIN_H int32 = 420, 240
	BOARD_SCALE  int32 = 10
)

var (
	MAX_FPS uint32 = 15
)

type point struct {
	x int32
	y int32
}

type color struct {
	r, g, b, a byte
}

func run() (err error) {

	var window *sdl.Window
	var renderer *sdl.Renderer
	var font *ttf.Font

	if err = ttf.Init(); err != nil {
		return err
	}
	defer ttf.Quit()

	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		return err
	}
	defer sdl.Quit()

	window, err = sdl.CreateWindow(
		"Snake-GO",
		sdl.WINDOWPOS_CENTERED,
		sdl.WINDOWPOS_CENTERED,
		WIN_W,
		WIN_H,
		sdl.WINDOW_SHOWN,
	)
	if err != nil {
		return err
	}
	defer window.Destroy()

	if font, err = ttf.OpenFont("/Users/mm/Library/Fonts/CascadiaMono.ttf", 14); err != nil {
		return err
	}
	defer font.Close()

	renderer, err = sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		return err
	}
	defer renderer.Destroy()

	// TODO: Need label for points
	// TODO: Need label for highscore

	board := NewBoard(int(WIN_W-20), int(WIN_H-40), 10, 30, int(BOARD_SCALE))
	snake := NewSnake(board)
	fruit := NewFruit(board)

	// Game Loop
	running := true

	for running {
		// 1. Get Keyboard input
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch t := event.(type) {
			case *sdl.QuitEvent:
				running = false
			case *sdl.KeyboardEvent:
				keyCode := t.Keysym.Sym
				keys := "" // for debugging

				if keyCode < 10000 {
					if t.State == sdl.PRESSED {

						// key: w
						if keyCode == 119 {
							// Change direction to up
							snake.ChangeDirection(1)
						}

						// key: d
						if keyCode == 100 {
							// Change direction to right
							snake.ChangeDirection(2)
						}

						// key: s
						if keyCode == 115 {
							// Change direction to down
							snake.ChangeDirection(3)
						}

						// key: a
						if keyCode == 97 {
							// change direction to left
							snake.ChangeDirection(4)
						}

						// key: ESC
						if keyCode == 27 {
							// change direction to stop
							snake.ChangeDirection(0)
						}

						if keyCode == 32 {
							snake.Reset()
							fruit.Update()
						}
					}

					// if t.Repeat > 0 {
					// 	keys += string(keyCode) + " repeating"
					// } else {
					// 	if t.State == sdl.RELEASED {
					// 		keys += string(keyCode) + " released"
					// 	} else if t.State == sdl.PRESSED {
					// 		keys += string(keyCode) + " pressed"
					// 	}
					// }
				}

				// For debug purposes
				if keys != "" {
					fmt.Printf("KeyCode: %v\n", keyCode)
					fmt.Println("Keys: " + keys)
				}
			}
		} // end for event

		// 2. Update objects
		update(snake, fruit)

		// 3. Draw objects

		// clear screen
		renderer.SetDrawColor(0, 0, 0, 255)
		renderer.Clear()

		var scoreText *sdl.Surface
		if scoreText, err = font.RenderUTF8Blended(fmt.Sprintf("Score: %d", snake.points), sdl.Color{R: 255, G: 255, B: 255, A: 255}); err != nil {
			return err
		}
		defer scoreText.Free()

		txt, err := renderer.CreateTextureFromSurface(scoreText)
		if err != nil {
			return err
		}
		defer txt.Destroy()

		if err = renderer.Copy(txt, nil, &sdl.Rect{X: 10, Y: 7, W: scoreText.W, H: scoreText.H}); err != nil {
			return err
		}

		board.Draw(renderer)
		snake.Draw(renderer)
		fruit.Draw(renderer)

		renderer.Present()
		sdl.Delay(1000 / MAX_FPS)
	}

	return
}

func update(s *snake, f *fruit) {
	if s.points == 5 {
		// speed = 20
		MAX_FPS = 20
	} else if s.points == 10 {
		// speed = 25
		MAX_FPS = 25
	} else if s.points == 15 {
		// speed = 30
		MAX_FPS = 30
	} else if s.points == 20 {
		// speed = 35
		MAX_FPS = 35
	} else if s.points < 5 {
		// speed = 15
		MAX_FPS = 15
	}

	// Update SCORE LABEL

	if s.DetectCollision() {
		s.Reset()
	} else {
		s.Eat(f)
		s.Update()
	}
}

func main() {

	rand.Seed(time.Now().UnixNano())

	if err := run(); err != nil {
		os.Exit(1)
	}
}
