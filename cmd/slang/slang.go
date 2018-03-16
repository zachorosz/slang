package main

import (
	"flag"
	"fmt"
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
	env = slang.MakeEnv()
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

func main() {
	flag.Usage = usage
	flag.Parse()

	if flag.NFlag() > 1 {
		flag.Usage()
	}

	narg := slang.Number(flag.NArg())
	argv := make(slang.Vector, flag.NArg())
	for i, arg := range flag.Args() {
		argv[i] = arg
	}

	env.UseSubrPackage("", subroutines.Primitives)
	env.Define(slang.Symbol("*ARGV*"), argv)
	env.Define(slang.Symbol("*NARG*"), narg)

	if *expression != "" {
		ok := readEvaluatePrint(*expression)
		if !ok {
			os.Exit(1)
		}
	} else {
		runREPL()
	}

	os.Exit(0)
}
