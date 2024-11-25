package models

import (
	"fmt"
	"gh-hubbub/structs"
	"strconv"

	"github.com/charmbracelet/bubbles/v2/textinput"
	tea "github.com/charmbracelet/bubbletea/v2"
)

type IntModel struct {
	Name      string
	fromInput textinput.Model
	toInput   textinput.Model
}

func intValidator(s, prompt string) error {
	_, err := strconv.Atoi(s)
	if err != nil {
		return fmt.Errorf("please enter an integer for the `%s` value", prompt)
	}

	return nil
}

func NewIntModel(title string, from, to int) IntModel {
	m := IntModel{
		Name:      title,
		fromInput: textinput.New(),
		toInput:   textinput.New(),
	}

	m.fromInput.Placeholder = fmt.Sprint(from)
	m.fromInput.Prompt = "From: "
	m.fromInput.CharLimit = 4
	m.fromInput.Validate = func(s string) error { return intValidator(s, m.fromInput.Prompt) }

	m.toInput.Placeholder = fmt.Sprint(to)
	m.toInput.Prompt = "To: "
	m.toInput.CharLimit = 4
	m.toInput.Validate = func(s string) error { return intValidator(s, m.toInput.Prompt) }

	m.fromInput.Focus()

	return m
}

func (m IntModel) Init() (tea.Model, tea.Cmd) {
	return m, textinput.Blink
}

func (m IntModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			return m, m.SendAddFilterMsg
		case "tab":
			if m.fromInput.Focused() {
				m.fromInput.Blur()
				m.toInput.Focus()
			} else {
				m.toInput.Blur()
				m.fromInput.Focus()
			}
		default:
			if m.fromInput.Focused() {
				m.fromInput, cmd = m.fromInput.Update(msg)
			} else {
				m.toInput, cmd = m.toInput.Update(msg)
			}
		}
	}

	return m, cmd
}

func (m IntModel) View() string {
	return m.Name + " " + m.fromInput.View() + " " + m.toInput.View()
}

func (m *IntModel) GetValue() (int, int) {
	from, _ := strconv.Atoi(m.fromInput.Value())
	to, _ := strconv.Atoi(m.toInput.Value())
	return from, to
}

func (m IntModel) SendAddFilterMsg() tea.Msg {
	from, to := m.GetValue()
	return structs.NewFilterInt(m.Name, from, to)
}
