package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/zachorosz/slang"
)

func readEvaluatePrint(sexpr string, env slang.Env) bool {
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

func main() {
	env := &slang.StandardEnv

	var expr string

	flag.Usage = func() {
		fmt.Printf("Usage: %s [ option | filename ] [ arguments ]\n", os.Args[0])

		flag.PrintDefaults()
	}

	flag.StringVar(&expr, "e", "", "Evaluate expression and print")
	flag.Parse()

	if flag.NFlag() > 1 {
		flag.Usage()
		os.Exit(1)
	}

	if expr != "" {
		ok := readEvaluatePrint(expr, *env)
		if !ok {
			os.Exit(1)
		}
	}

	os.Exit(0)
}
