package p64

import (
	"fmt"
	"math/rand"
)

const (
	frameWidth  = 64
	frameHeight = 64
	frameSize   = frameWidth * frameHeight
)

type frameBuffer struct {
	data [frameSize]int
}

func newFramebuffer(w, h int) *frameBuffer {
	return &frameBuffer{}
}

func (f *frameBuffer) Rand() {
	for i := 0; i < frameSize; i++ {
		f.data[i] = rand.Intn(2)
	}
}

func (f *frameBuffer) Clear() {
	for i := 0; i < frameSize; i++ {
		f.data[i] = 0
	}
}

func (f *frameBuffer) Set(x, y, value int) {
	if x < 0 || x > 63 || y < 0 || y > 63 {
		return
	}
	f.data[y*frameWidth+x] = value
}

func (f *frameBuffer) At(x, y int) int {
	return f.data[y*frameWidth+x]
}

func (f *frameBuffer) Line(x, y, x2, y2 int) {

	dx := 1
	if x2 < x {
		dx = -1
	}

	dy := 1
	if y2 < y {
		dy = -1
	}
	fmt.Println("xy", dx, dy)

}
