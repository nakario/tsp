package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/nakario/tsp"
)

var (
	debug *bool = flag.Bool("debug", false, "enable debug output")
	strict *bool = flag.Bool("strict", false, "enable strict mode: stop execution with any error")
)

func compile(path string) (*tsp.Program, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open the source: %w", err)
	}
	defer f.Close()

	prog, err := tsp.Compile(f, tsp.Config{DebugMode: *debug, StrictMode: *strict})
	if err != nil {
		return nil, fmt.Errorf("failed to compile the source: %w", err)
	}
	return prog, nil
}

func main() {
	flag.Parse()
	if flag.NArg() != 1 {
		flag.Usage()
		return
	}

	prog, err := compile(flag.Arg(0))
	if err != nil {
		fmt.Fprintln(os.Stderr, "Compile error:", err)
		return
	}

	if err := prog.Run(); err != nil {
		fmt.Fprintln(os.Stderr, "Runtime error:", err)
	}
}
