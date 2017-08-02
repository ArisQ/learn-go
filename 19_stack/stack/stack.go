package stack

import (
	"errors"
	"fmt"
	"runtime"
	"strconv"
)

const stackLimit int = 4

type stack struct {
	data [stackLimit]int
	idx  int
}

type Stack struct {
	stack
}

func NewStack() (s *stack) {
	s = new(stack)
	runtime.SetFinalizer(s, func(s *stack) {
		fmt.Println("stack finalizer")
	})
	return
}

func (s *stack) Push(d int) error {
	if s.idx >= stackLimit {
		return errors.New("stack full")
	}
	s.data[s.idx] = d
	s.idx++
	return nil
}

func (s *stack) Pop() (d int, e error) {
	if s.idx == 0 {
		e = errors.New("stack empty")
		return
	}
	s.idx--
	d = s.data[s.idx]
	return
}

func (s *stack) String() string {
	var str string
	if s.idx == 0 {
		str = "stack is empty."
		return str
	}
	str = "stack:"
	for i := 0; i < s.idx; i++ {
		str += "[" + strconv.Itoa(i) + ":" + strconv.Itoa(s.data[i]) + "]"
	}
	return str
}
