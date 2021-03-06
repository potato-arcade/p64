package p64

import (
	"fmt"
	"image"
	"image/color"
	"math/rand"

	"fyne.io/fyne"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/theme"
)

func (p *P64) CreateRenderer() fyne.WidgetRenderer {
	renderer := &renderer{frameBuffer: p.frameBuffer}

	render := canvas.NewRaster(renderer.frame)
	renderer.render = render
	renderer.objects = []fyne.CanvasObject{render}
	renderer.ApplyTheme()
	p.renderer = renderer

	return renderer
}

type renderer struct {
	render      *canvas.Raster
	objects     []fyne.CanvasObject
	color       color.Color
	backColor   color.Color
	frameBuffer *frameBuffer
	size        fyne.Size
	img         *image.RGBA
	static      bool
}

func (r *renderer) Static(s bool) {
	if r != nil {
		r.static = s
		if !s && r.img != nil {
			w := r.img.Bounds().Size().X
			h := r.img.Bounds().Size().Y
			for y := 0; y < h; y++ {
				for x := 0; x < w; x++ {
					r.img.Set(x, y, r.backColor)
				}
			}
		}
	}
}

func (r *renderer) getImage(w, h int) *image.RGBA {
	if r.img == nil || r.img.Bounds().Size().X != w || r.img.Bounds().Size().Y != h {
		fmt.Print(".")
		r.img = image.NewRGBA(image.Rect(0, 0, w, h))
	}
	return r.img
}

func (r *renderer) frame(w, h int) image.Image {
	img := r.getImage(w, h)
	if r.static {
		for y := 0; y < h; y++ {
			for x := 0; x < w; x++ {
				c := r.color
				if rand.Intn(2) == 1 {
					c = r.backColor
				}
				img.Set(x, y, c)
			}
		}
		return img
	}
	bx := int(float64(w) / float64(frameWidth))
	by := int(float64(h) / float64(frameHeight))
	for y := 0; y < frameHeight; y++ {
		for x := 0; x < frameWidth; x++ {
			c := r.backColor
			if r.frameBuffer.At(x, y) == 1 {
				c = r.color
			}
			for yy := 0; yy < by; yy++ {
				for xx := 0; xx < bx; xx++ {
					img.Set(x*bx+xx, y*by+yy, c)
				}
			}
		}
	}

	return img
}

func (r *renderer) MinSize() fyne.Size {
	return fyne.NewSize(400, 300)
}

func (r *renderer) Layout(size fyne.Size) {
	r.render.Resize(size)
	r.size = size
}

func (r *renderer) ApplyTheme() {
	r.color = theme.TextColor()
	r.backColor = theme.BackgroundColor()
}

func (r *renderer) BackgroundColor() color.Color {
	return theme.BackgroundColor()
}

func (r *renderer) Refresh() {
	canvas.Refresh(r.render)
}

func (r *renderer) Objects() []fyne.CanvasObject {
	return r.objects
}

func (r *renderer) Destroy() {

}
