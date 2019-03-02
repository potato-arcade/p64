package p64

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/davecgh/go-spew/spew"
	"github.com/skx/gobasic/eval"
	"github.com/skx/gobasic/object"
	"github.com/skx/gobasic/tokenizer"
)

func (p *P64) registerFunctions(e *eval.Interpreter) {

	blankObj := &object.NumberObject{Value: 0.0}

	e.RegisterBuiltin("PEEK", 1, func(env interface{}, args []object.Object) object.Object {
		key := int(args[0].(*object.NumberObject).Value)
		value, ok := p.ram[key]
		if !ok {
			return blankObj
		}
		return value
	})

	e.RegisterBuiltin("POKE", 2, func(env interface{}, args []object.Object) object.Object {
		key := int(args[0].(*object.NumberObject).Value)
		p.ram[key] = args[1]
		return blankObj
	})

	e.RegisterBuiltin("CLEAR", 0, func(env interface{}, args []object.Object) object.Object {
		p.frameBuffer.Clear()
		return blankObj
	})

	e.RegisterBuiltin("SET", 3, func(env interface{}, args []object.Object) object.Object {
		x := int(args[0].(*object.NumberObject).Value)
		y := int(args[1].(*object.NumberObject).Value)
		v := int(args[2].(*object.NumberObject).Value)
		p.frameBuffer.Set(x, y, v)
		return blankObj
	})

	e.RegisterBuiltin("DEBUG", 0, func(env interface{}, args []object.Object) object.Object {
		fmt.Println("DEBUG Memory Banks")
		spew.Dump(env)
		spew.Dump(p.ram)
		for k, v := range p.code {
			fmt.Println("Interrupt Handler:",
				k,
				"\n-------------------------\n",
				v,
				"\n-------------------------\n",
			)
		}
		return blankObj
	})
}

func (p *P64) LoadROM() {
	if p.romFile == "" {
		return
	}
	b, err := ioutil.ReadFile(p.romFile) // just pass the file name
	if err != nil {
		fmt.Println("ERROR ROM LOAD:", err)
		return
	}
	p.src = string(b)

	t := tokenizer.New(p.src)
	e, err := eval.New(t)
	if err != nil {
		fmt.Println("ERROR ROM IMPORT:", err)
		return
	}
	p.registerFunctions(e)

	// Now we need to split the code into sections
	p.compileInterrupt("KEYDOWN")
	p.compileInterrupt("KEYUP")
	p.compileInterrupt("VSYNC")

	// Run Init Code
	if err := e.Run(); err != nil {
		fmt.Println("ERROR ROM EXEC INIT BOOT:", err)
		return
	}
}

func (p *P64) compileInterrupt(intr string) {
	if i := strings.Index(p.src, ".INTR "+intr); i != -1 {
		code := p.src[i+6:]
		ii := strings.Index(code, "\n")
		if ii != -1 {
			if code[:ii] == intr {
				code = code[ii:]
				endcode := strings.Index(code, "END\n")
				if endcode != -1 {
					code = code[:endcode]
				}
				if false {
					t := tokenizer.New(code)
					e, err := eval.New(t)
					if err != nil {
						fmt.Println("ERROR ROM IMPORT:", err)
						return
					}
					p.registerFunctions(e)
				}
				p.code[intr] = code
				return
			}
		}
	}
}

func (p *P64) interrupt(what string, data int) {
	handler, ok := p.code[what]
	if !ok {
		fmt.Println("No interrupt for", what, data)
		spew.Dump(p.code)
		return
	}
	t := tokenizer.New(handler)
	e, err := eval.New(t)
	if err != nil {
		fmt.Println("ERROR ROM INTR:", err)
		//delete(p.code, what)
	}
	p.registerFunctions(e)

	// Set the KEY variable
	e.SetVariable("KEY", &object.NumberObject{Value: float64(data)})

	// Run the interrupt code
	if err := e.Run(); err != nil {
		fmt.Println("ERROR ROM EXEC INTR:", err)
		//delete(p.code, what)
		return
	}
}
