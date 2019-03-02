package p64

import (
	"fyne.io/fyne"
	"fyne.io/fyne/widget"
)

func (p *P64) Size() fyne.Size {
	return p.size
}

func (p *P64) Resize(size fyne.Size) {
	p.size = size
	widget.Renderer(p).Layout(size)
}

func (p *P64) Position() fyne.Position {
	return p.position
}

func (p *P64) Move(pos fyne.Position) {
	p.position = pos
	widget.Renderer(p).Layout(p.size)
}

func (p *P64) MinSize() fyne.Size {
	return widget.Renderer(p).MinSize()
}

func (p *P64) Visible() bool {
	return p.hidden
}

func (p *P64) Show() {
	p.hidden = false
}

func (p *P64) Hide() {
	p.hidden = true
}
