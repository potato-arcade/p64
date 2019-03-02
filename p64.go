package p64

import (
	"math/rand"
	"sync"
	"time"

	"fyne.io/fyne"
	"fyne.io/fyne/driver/desktop"
	"fyne.io/fyne/widget"
	"github.com/skx/gobasic/object"
)

const hz = 24

func New(app fyne.App) *P64 {
	app.Settings().SetTheme(&p64Theme{})
	window := app.NewWindow("Potato 64")
	p := &P64{
		Hz:          hz,
		frameBuffer: newFramebuffer(64, 64),
		ram:         make(map[int]object.Object),
		code:        make(map[string]string),
	}
	p.run()
	p.On()
	window.SetContent(p)
	window.Show()
	window.Canvas().(desktop.Canvas).SetOnKeyDown(func(ev *fyne.KeyEvent) {
		switch string(ev.Name) {
		case "Escape":
			p.Reboot()
		default:
			p.interrupt("KEYDOWN", string(ev.Name))
		}
	})
	window.Canvas().(desktop.Canvas).SetOnKeyUp(func(ev *fyne.KeyEvent) {
		switch string(ev.Name) {
		case "Escape":
		default:
			p.interrupt("KEYUP", string(ev.Name))
		}
	})

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
	ramMutex    sync.RWMutex
	ram         map[int]object.Object
	romFile     string
	src         string
	code        map[string]string
}

func (p *P64) On() {
	p.Power = true
	p.Reboot()
}

func (p *P64) Off() {
	p.Power = false
}

func (p *P64) Reboot() {
	p.Booting = 3*rand.Intn(hz) + hz
	p.ram = make(map[int]object.Object)
	p.code = make(map[string]string)
}

func (p *P64) InsertROM(f string) {
	p.romFile = f
}

func (p *P64) run() {
	go func() {
		tick := time.NewTicker(time.Second / hz)
		t1 := 0

		for {
			select {
			case <-tick.C:
				// Powered Off
				if !p.Power {
					p.frameBuffer.Clear()
					continue
				}

				// Booting up and loading ROM cartrige
				if p.Power && p.Booting > 0 {
					if rand.Intn(10) == 1 {
						p.frameBuffer.Rand()
					}
					p.Booting--
					if p.Booting < 1 {
						p.frameBuffer.Clear()
						p.LoadROM()
					}
					widget.Refresh(p)
					continue
				}

				// It lives !
				p.interrupt("VSYNC", t1)
				t1++
				widget.Refresh(p)
			}
		}
	}()
}
