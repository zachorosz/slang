package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/c-bata/go-prompt"
	"github.com/zachorosz/slang"
	"github.com/zachorosz/slang/parser"
)

const (
	replExit = "(exit)"
)

var (
	expression = flag.String("e", "", "Evaluate expression and print")
	program    = ""
	env        = slang.MakeEnv(nil)
)

func usage() {
	fmt.Println("usage: slang [[-e expression | filename] [arguments]]")
	flag.PrintDefaults()
	os.Exit(2)
}

func readEvaluatePrint(sexpr string) bool {
	expr, err := parser.Parse("REPL", sexpr)
	if err != nil {
		fmt.Printf("%s\n", err)
		return false
	}

	result, err := slang.Evaluate(expr[0], env)
	if err != nil {
		fmt.Printf("%s\n", err)
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

func setupEnv(env *slang.Env, argc int, args []string) {
	narg := slang.Number(argc)
	argv := make(slang.Vector, argc)
	for i, arg := range args {
		argv[i] = slang.Str(arg)
	}

	env.UseSubrPackage("", Primitives)
	env.Define(slang.Symbol("*ARGV*"), argv)
	env.Define(slang.Symbol("*NARG*"), narg)
}

func main() {
	flag.Usage = usage
	flag.Parse()

	if flag.NFlag() > 1 {
		flag.Usage() // Exit(2)
	}

	// filename passed as argument
	if flag.NFlag() == 0 && flag.NArg() > 0 {
		var filename = flag.Arg(0)
		program = readFile(filename)

		args := flag.Args()[1:]
		argc := len(args)
		setupEnv(&env, argc, args)

		exprs, err := parser.Parse(filename, program)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		for _, expr := range exprs {
			v, err := slang.Evaluate(expr, env)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			fmt.Println(v)
		}
	} else {
		// run REPL or evaluate expression passed via -e flag
		setupEnv(&env, flag.NArg(), flag.Args())

		if *expression != "" {
			ok := readEvaluatePrint(*expression)
			if !ok {
				os.Exit(1)
			}
		} else {
			runREPL()
		}
	}

	os.Exit(0)
}
