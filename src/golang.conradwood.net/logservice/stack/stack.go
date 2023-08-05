package stack

import (
	"sync"
)

var (
	stacklock sync.Mutex
	stacks    = make(map[string]*Stack)
)

const (
	STACKSIZE = 100
)

type Stack struct {
	lock    sync.Mutex
	entries []string
}

func Get(stackname string) *Stack {
	stacklock.Lock()
	defer stacklock.Unlock()
	stack := stacks[stackname]
	if stack == nil {
		stack = &Stack{}
		stacks[stackname] = stack
	}
	return stack
}

func (s *Stack) Add(line string) {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.entries = append(s.entries, line)
	for len(s.entries) > STACKSIZE {
		s.entries = s.entries[1:]
	}
}

func (s *Stack) Get() []string {
	return s.entries
}
