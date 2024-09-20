package models

import (
	"errors"
	"reflect"

	tea "github.com/charmbracelet/bubbletea"
)

// Stack represents a stack data structure
type Stack struct {
	elements []tea.Model
}

// Push adds an element to the top of the stack
func (s *Stack) Push(element tea.Model) {
	s.elements = append(s.elements, element)
}

// Pop removes and returns the top element of the stack
func (s *Stack) Pop() (tea.Model, error) {
	if len(s.elements) == 0 {
		return nil, errors.New("stack is empty")
	}
	element := s.elements[len(s.elements)-1]
	s.elements = s.elements[:len(s.elements)-1]
	return element, nil
}

// Peek returns the top element of the stack without removing it
func (s *Stack) Peek() (tea.Model, error) {
	if len(s.elements) == 0 {
		return nil, errors.New("stack is empty")
	}
	element := s.elements[len(s.elements)-1]
	return element, nil
}

func (s *Stack) Len() int {
	return len(s.elements)
}

func (s *Stack) TypeOfHead() reflect.Type {
	return reflect.TypeOf(s.elements[len(s.elements)-1])
}
