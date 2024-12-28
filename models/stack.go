package models

import (
	"errors"
	"fmt"
	"reflect"

	tea "github.com/charmbracelet/bubbletea/v2"
)

// Stack represents a stack data structure
type Stack struct {
	elements []tea.Model
}

func callPointerMethod(i interface{}, methodName string, args ...interface{}) (interface{}, error) {
	val := reflect.ValueOf(i)

	// Get to the underlying value and create pointer if needed
	if val.Kind() != reflect.Ptr {
		// Make a new pointer to a new struct value
		ptr := reflect.New(val.Type())
		// Set the value at the pointer to our original value
		ptr.Elem().Set(val)
		// Use the pointer for method calling
		val = ptr
	}

	// Find the method
	method := val.MethodByName(methodName)
	if !method.IsValid() {
		return nil, fmt.Errorf("method %s not found", methodName)
	}

	// Convert args to reflect.Value slice
	reflectArgs := make([]reflect.Value, len(args))
	for i, arg := range args {
		reflectArgs[i] = reflect.ValueOf(arg)
	}

	// Call the method
	method.Call(reflectArgs)

	// Return the modified object
	return val.Interface(), nil
}

func (s Stack) SetDimensions(width, height int) {
	for idx := range s.elements {
		callPointerMethod(s.elements[idx], "SetDimensions", width, height)
	}
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
