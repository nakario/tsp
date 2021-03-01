package tsp

import (
	"errors"
	"fmt"
	"io"
	"sort"
)

type Instruction struct {
	ChebyshevRadius int
	ManhattanRadius int
	Point Point
	Exec func(p *Program) error
	Str string
}

func (i Instruction) String() string {
	return i.Str
}

func MakeExecPush(n int) func(*Program) error {
	return func(p *Program) error {
		p.Stack.Push(Value(n))
		return nil
	}
}

func ExecDup(p *Program) error {
	return p.Stack.Dup()
}

func MakeExecSwap(n int) func(*Program) error {
	return func(p *Program) error {
		return p.Stack.Swap(n)
	}
}

func ExecSub(p *Program) error {
	return p.Stack.Sub()
}

func ExecGreater(p *Program) error {
	return p.Stack.Greater()
}

func MakeExecJumpZero(n int) func(*Program) error {
	return func(p *Program) error {
		v, err := p.Stack.Top()
		if err != nil {
			return fmt.Errorf("failed to get the top value of the stack: %w", err)
		}
		if v == 0 {
			p.PC = -1 + sort.Search(len(p.Instructions), func(i int) bool {
				return n <= p.Instructions[i].ChebyshevRadius
			})
		}
		return nil
	}
}

func ExecGetChar(p *Program) error {
	c, _, err := p.Stdin.ReadRune()
	if errors.Is(err, io.EOF) {
		c = 0
	} else if err != nil {
		return fmt.Errorf("failed to get a char: %w", err)
	}
	p.Stack.Push(Value(c))
	return nil
}

func ExecPutChar(p *Program) error {
	c, err := p.Stack.Pop()
	if err != nil {
		return fmt.Errorf("failed to pop a char: %w", err)
	}
	_, err = p.Stdout.WriteString(string([]rune{c.Rune()}))
	if err != nil {
		return fmt.Errorf("failed to write a char: %w", err)
	}
	return nil
}

func ExecNop(p *Program) error {
	return nil
}

func MakeInstruction(p Point, stepCount int) Instruction {
	var exec func(p *Program) error = ExecNop
	str := fmt.Sprintf("nop (%d, %d)", p.X, p.Y)
	switch area := p.Area(); area {
	case AreaENE:
		exec = MakeExecPush(stepCount)
		str = fmt.Sprintf("push %d (%d, %d)", stepCount, p.X, p.Y)
	case AreaNNE:
		exec = ExecDup
		str = fmt.Sprintf("dup (%d, %d)", p.X, p.Y)
	case AreaNNW:
		exec = MakeExecSwap(stepCount)
		str = fmt.Sprintf("swap %d (%d, %d)", stepCount, p.X, p.Y)
	case AreaWNW:
		exec = ExecSub
		str = fmt.Sprintf("sub (%d, %d)", p.X, p.Y)
	case AreaWSW:
		exec = ExecGreater
		str = fmt.Sprintf("greater (%d, %d)", p.X, p.Y)
	case AreaSSW:
		exec = MakeExecJumpZero(stepCount)
		str = fmt.Sprintf("jump-zero %d (%d, %d)", stepCount, p.X, p.Y)
	case AreaSSE:
		exec = ExecGetChar
		str = fmt.Sprintf("getchar (%d, %d)", p.X, p.Y)
	case AreaESE:
		exec = ExecPutChar
		str = fmt.Sprintf("putchar (%d, %d)", p.X, p.Y)
	}
	return Instruction{
		ChebyshevRadius: origin.ChebyshevDistance(p),
		ManhattanRadius: origin.ManhattanDistance(p),
		Point: p,
		Exec: exec,
		Str: str,
	}
}
