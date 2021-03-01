package tsp

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"sort"
)

var origin = Point{X: 0, Y: 0}

// Constants for moving
var (
	STAY = Vector{X: 0, Y: 0}
	LEFT = Vector{X: -1, Y: 0}
	DOWN = Vector{X: 0, Y: -1}
	RIGHT = Vector{X: 1, Y: 0}
	UP = Vector{X: 0, Y: 1}
)

// Constants for source code
const (
	A = 'A'
	S = 'S'
	D = 'D'
	W = 'W'
	VISIT = '$'
)

type Value rune

func (v Value) Rune() rune {
	return rune(v)
}

func (v Value) Int32() int32 {
	return int32(v)
}

func (v Value) Uint32() uint32 {
	return uint32(v)
}

type Config struct {
	StrictMode bool
	DebugMode bool
}

type Program struct {
	Instructions []Instruction
	PC int
	Stack *Stack
	Stdin *bufio.Reader
	Stdout *bufio.Writer
	Config Config
}

func (p *Program) Run() error {
	for p.PC < len(p.Instructions) {
		err := p.Step()
		if err != nil {
			if p.Config.StrictMode {
				return fmt.Errorf("runtime error: %w", err)
			} else if p.Config.DebugMode {
				fmt.Fprintln(os.Stderr, "[DEBUG] Error: ", err)
			}
		}
	}
	return nil
}

func (p *Program) Step() error {
	if p.Config.DebugMode {
		fmt.Fprintln(os.Stderr, "[DEBUG] Stack: ", p.Stack)
		fmt.Fprintf(os.Stderr, "[DEBUG] PC: %4d, %s\n", p.PC, p.Instructions[p.PC].String())
	}
	err := p.Instructions[p.PC].Exec(p)
	p.Stdout.Flush()
	p.PC++
	return err
}

func Compile(f io.Reader, config Config) (*Program, error) {
	prog := &Program{
		Instructions: make([]Instruction, 0, 0),
		Stack: NewStack(),
		Stdin: bufio.NewReader(os.Stdin),
		Stdout: bufio.NewWriter(os.Stdout),
		Config: config,
	}

	code, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, fmt.Errorf("failed to read the source: %w", err)
	}

	type ds struct{
		cd, md int
	}

	var p Point
	stepCount := 0
	visited := make(map[ds]int)
	originVisited := false
	for _, c := range code {
		v := STAY
		switch rune(c) {
		case A:
			v = LEFT
		case S:
			v = DOWN
		case D:
			v = RIGHT
		case W:
			v = UP
		case VISIT:
			cd := origin.ChebyshevDistance(p)
			md := origin.ManhattanDistance(p)
			if n, ok := visited[ds{cd, md}]; ok {
				p2 := prog.Instructions[n-1].Point
				return nil, fmt.Errorf(
					"there are 2 or more vertices having the same Chebyshev and Manhattan distances from the origin: " +
					"inst (%d (%d, %d)) and inst (%d (%d, %d))",
					n, p2.X, p2.Y,
					len(prog.Instructions) + 1, p.X, p.Y,
				)
			}
			visited[ds{cd, md}] = len(prog.Instructions) + 1
			if p == origin {
				originVisited = true
				continue
			}
			inst := MakeInstruction(p, stepCount)
			prog.Instructions = append(prog.Instructions, inst)
			stepCount = 0
		}
		if v != STAY {
			if originVisited {
				fmt.Fprintln(os.Stderr, "code remaining after visiting the origin")
				break
			}
			p = p.Add(v)
			stepCount++
		}
	}
	if !originVisited {
		return nil, fmt.Errorf("the salesperson should visit (0, 0) at the end of the code")
	}

	sort.Slice(prog.Instructions, func(i, j int) bool {
		if prog.Instructions[i].ChebyshevRadius == prog.Instructions[j].ChebyshevRadius {
			return prog.Instructions[i].ManhattanRadius < prog.Instructions[j].ManhattanRadius
		}
		return prog.Instructions[i].ChebyshevRadius < prog.Instructions[j].ChebyshevRadius
	})

	if config.DebugMode {
		fmt.Fprintln(os.Stderr, "[DEBUG] --- Instructions ---")
		for _, v := range prog.Instructions {
			fmt.Fprintln(os.Stderr, "[DEBUG]", v)
		}
		fmt.Fprintln(os.Stderr, "[DEBUG] --------------------")
	}

	return prog, nil
}
