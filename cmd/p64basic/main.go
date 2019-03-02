package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/davecgh/go-spew/spew"

	"github.com/skx/gobasic/eval"
	"github.com/skx/gobasic/object"
	"github.com/skx/gobasic/token"
	"github.com/skx/gobasic/tokenizer"
)

// This version-string will be updated via travis for generated binaries.
var version = "master/unreleased"

func main() {

	//
	// Setup some command-line flags
	//
	lex := flag.Bool("lex", false, "Show the output of the lexer.")
	trace := flag.Bool("trace", false, "Trace execution.")
	vers := flag.Bool("version", false, "Show our version and exit.")

	//
	// Parse the flags
	//
	flag.Parse()

	//
	// Showing the version?
	//
	if *vers {
		fmt.Printf("gobasic %s\n", version)
		os.Exit(1)
	}

	//
	// Test we have a file to interpret
	//
	if len(flag.Args()) != 1 {
		fmt.Printf("Usage: gobasic /path/to/input/script.bas\n")
		os.Exit(2)
	}

	//
	// Load the file.
	//
	data, err := ioutil.ReadFile(flag.Args()[0])
	if err != nil {
		fmt.Printf("Error reading %s - %s\n", flag.Args()[0], err.Error())
		os.Exit(3)
	}

	//
	// Tokenize
	//
	t := tokenizer.New(string(data))

	//
	// Are we dumping tokens?
	//
	if *lex {
		for {
			tok := t.NextToken()
			fmt.Printf("%v\n", tok)
			if tok.Type == token.EOF {
				break
			}
		}
		os.Exit(0)
	}

	//
	// Create a new evaluator, to run the BASIC program.
	//
	e, err := eval.New(t)
	if err != nil {
		fmt.Printf("Error constructing interpreter:\n\t%s\n", err.Error())
		os.Exit(0)
	}

	ram := make(map[int]object.Object)
	blankObj := &object.NumberObject{Value: 0.0}

	// Register some builtins
	e.RegisterBuiltin("PEEK", 1, func(env interface{}, args []object.Object) object.Object {
		key := int(args[0].(*object.NumberObject).Value)
		fmt.Printf("PEEK <- %v ", key)
		value, ok := ram[key]
		if !ok {
			fmt.Println("(NOT FOUND)")
			return blankObj
		}
		fmt.Println("= ", value)
		return value
	})

	e.RegisterBuiltin("POKE", 2, func(env interface{}, args []object.Object) object.Object {
		key := int(args[0].(*object.NumberObject).Value)
		fmt.Printf("POKE %v -> %v\n", args[1], key)
		ram[key] = args[1]
		return blankObj
	})

	e.RegisterBuiltin("DEBUG", 0, func(env interface{}, args []object.Object) object.Object {
		fmt.Println("DEBUG Memory Banks")
		spew.Dump(ram)
		return blankObj
	})

	e.RegisterBuiltin(".INTR", 1, func(env interface{}, args []object.Object) object.Object {
		fmt.Println("interrupt:", args[0].(*object.StringObject).Value)
		return blankObj
	})

	//
	// Enable debugging if we should.
	//
	e.SetTrace(*trace)

	//
	// Run the code, and report on any error.
	//
	err = e.Run()
	if err != nil {
		fmt.Printf("Error running program:\n\t%s\n", err.Error())
	}
}

func registerFunctions(e eval.Interpreter) {
}

// Extra BASIC functions

func peekFunction(env interface{}, args []object.Object) object.Object {
	fmt.Printf("PEEK called with %v\n", args[0])
	return &object.NumberObject{Value: 0.0}
}

func pokeFunction(env interface{}, args []object.Object) object.Object {
	if len(args) == 2 {
		varname := args[1]
		value := args[2]
		fmt.Println("Storing", value, "into", varname)
	}
	return &object.NumberObject{Value: 0.0}
}
