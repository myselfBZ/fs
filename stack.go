package main

import "errors"

type Stack struct{
    elements []string
}

var (
    Empty = errors.New("Stack is empty")
)


func (s *Stack) Append(e string)  {
    s.elements = append(s.elements, e)
}

func (s *Stack) Pop() error {
    if len(s.elements) == 0{
        return Empty 
    }
    s.elements = s.elements[:len(s.elements) -1]
    return nil
}
