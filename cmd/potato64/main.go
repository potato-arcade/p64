package main

import (
	"time"

	"fyne.io/fyne/app"
	"github.com/potato-arcade/p64"
	"github.com/steveoc64/memdebug"
)

func main() {
	app := app.New()
	p := p64.New(app)
	memdebug.Print(time.Now(), "got p", p)
	app.Run()
}
