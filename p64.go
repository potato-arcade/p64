package p64

import (
	"fmt"
	"math/rand"
	"time"

	"fyne.io/fyne"
	"fyne.io/fyne/widget"
)

const hz = 24

func New(app fyne.App) *P64 {
	app.Settings().SetTheme(&p64Theme{})
	window := app.NewWindow("Potato 64")
	p := &P64{
		Hz:          hz,
		frameBuffer: newFramebuffer(64, 64),
	}
	p.run()
	p.On()
	window.SetContent(p)
	window.Show()
	return p
}

type P64 struct {
	Power       bool
	Booting     int // ms left to boot
	Hz          int
	size        fyne.Size
	position    fyne.Position
	hidden      bool
	frameBuffer *frameBuffer
}

func (p *P64) On() {
	p.Power = true
	p.Reboot()
}

func (p *P64) Off() {
	p.Power = false
}

func (p *P64) Reboot() {
	p.Booting = rand.Intn(hz) + hz
}

func (p *P64) run() {
	go func() {
		tick := time.NewTicker(time.Second / hz)

		for {
			select {
			case <-tick.C:
				fmt.Print(".")
				// Powered Off
				if !p.Power {
					p.frameBuffer.Clear()
				}
				if p.Power && p.Booting > 0 {
					p.frameBuffer.Rand()
					p.Booting--
					if p.Booting < 1 {
						p.frameBuffer.Clear()
						// we are now powered on and ready to run
					}
					widget.Refresh(p)
				}
			}
		}
	}()
}
