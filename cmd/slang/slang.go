package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/c-bata/go-prompt"
	"github.com/zachorosz/slang"
	"github.com/zachorosz/slang/subroutines"
)

const (
	replExit = "(exit)"
)

var (
	expression = flag.String("e", "", "Evaluate expression and print")
)

var (
	program = ""
	env     = slang.MakeEnv(nil)
)

func usage() {
	fmt.Println("usage: slang [[-e expression | filename] [arguments]]")
	flag.PrintDefaults()
	os.Exit(2)
}

func readEvaluatePrint(sexpr string) bool {
	expr, err := slang.Read(sexpr)
	if err != nil {
		fmt.Printf("Read Error: %s\n", err)
		return false
	}

	result, err := slang.Evaluate(expr, env)
	if err != nil {
		fmt.Printf("Eval Error: %s\n", err)
		return false
	}

	fmt.Println(result)

	return true
}

func completer(d prompt.Document) []prompt.Suggest {
	// TODO(zachorosz): add autocomplete suggestions
	s := []prompt.Suggest{}
	return prompt.FilterHasPrefix(s, d.GetWordBeforeCursor(), true)
}

func runREPL() {
	for {
		line := prompt.Input("slang> ", completer)
		if line == replExit {
			return
		}
		readEvaluatePrint(line)
	}
}

func readFile(filename string) string {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return string(b)
}

func main() {
	var argc int
	var args []string

	flag.Usage = usage
	flag.Parse()

	if flag.NFlag() > 1 {
		flag.Usage() // Exit(2)
	}

	// filename passed as argument
	if flag.NFlag() == 0 && flag.NArg() > 0 {
		var filename = flag.Arg(0)
		contents := readFile(filename)
		program = contents

		args = flag.Args()[1:]
		argc = len(args)
	} else {
		argc = flag.NArg()
		args = flag.Args()
	}

	narg := slang.Number(argc)
	argv := make(slang.Vector, argc)
	for i, arg := range args {
		argv[i] = slang.Str(arg)
	}

	env.UseSubrPackage("", subroutines.Primitives)
	env.Define(slang.Symbol("*ARGV*"), argv)
	env.Define(slang.Symbol("*NARG*"), narg)

	if *expression != "" {
		ok := readEvaluatePrint(*expression)
		if !ok {
			os.Exit(1)
		}
	} else if program != "" {
		expressions, err := slang.ReadAll(program)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		for _, expr := range expressions {
			v, err := slang.Evaluate(expr, env)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			fmt.Println(v)
		}
	} else {
		runREPL()
	}

	os.Exit(0)
}
