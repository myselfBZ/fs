package main

type Stack struct{
    elements []string
}

func (s *Stack) Append(e string)  {
    s.elements = append(s.elements, e)
}

func (s *Stack) Pop(){
    s.elements = s.elements[len(s.elements)-1:]
}
