package main

import (
	"errors"

	"github.com/myselfBZ/fs/entries"
)

type Stack struct{
    dirs []entries.Directory
}

var (
    Empty = errors.New("Stack is empty")
)


func (s *Stack) Append(d entries.Directory)  {
    s.dirs = append(s.dirs, d)
}

func (s *Stack) Pop() error {
    if len(s.dirs) == 0{
        return Empty 
    }
    s.dirs = s.dirs[:len(s.dirs) -1]
    return nil
}
