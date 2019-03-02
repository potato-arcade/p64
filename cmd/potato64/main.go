package main

import (
	"fmt"
	"os"

	"fyne.io/fyne/app"
	"github.com/potato-arcade/p64"
)

func main() {
	app := app.New()
	p := p64.New(app)
	switch len(os.Args) {
	case 1:
		app.Run()
	case 2:
		p.InsertROM(os.Args[1])
		app.Run()
	default:
		fmt.Println("USAGE:", os.Args[0], "ROM-Filename")
	}
}
