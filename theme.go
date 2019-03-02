package p64

import (
	"image/color"

	"fyne.io/fyne"
	"fyne.io/fyne/theme"
)

type p64Theme struct{}

func (p64Theme) BackgroundColor() color.Color {
	return color.RGBA{0x10, 0x10, 0x10, 0xff}
}

func (p64Theme) ButtonColor() color.Color {
	return color.White
}

func (p64Theme) HyperlinkColor() color.Color {
	return color.White
}

func (p64Theme) TextColor() color.Color {
	return color.White
}

func (p64Theme) PlaceHolderColor() color.Color {
	return color.White
}

func (p64Theme) PrimaryColor() color.Color {
	return color.White
}

func (p64Theme) FocusColor() color.Color {
	return color.White
}

func (p64Theme) ScrollBarColor() color.Color {
	return color.Black
}

func (p64Theme) TextSize() int {
	return 18
}

func (p64Theme) TextFont() fyne.Resource {
	return theme.DefaultTextFont()
}

func (p64Theme) TextBoldFont() fyne.Resource {
	return theme.DefaultTextBoldFont()
}

func (p64Theme) TextItalicFont() fyne.Resource {
	return theme.DefaultTextItalicFont()
}

func (p64Theme) TextBoldItalicFont() fyne.Resource {
	return theme.DefaultTextBoldItalicFont()
}

func (p64Theme) TextMonospaceFont() fyne.Resource {
	return font
}

func (p64Theme) Padding() int {
	return 0
}

func (p64Theme) IconInlineSize() int {
	return 10
}

func (p64Theme) ScrollBarSize() int {
	return 10
}
