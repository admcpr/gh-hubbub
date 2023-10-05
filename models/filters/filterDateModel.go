package filters

import (
	"fmt"
	"gh-hubbub/messages"

	"strconv"
	"strings"
	"time"

	"gh-hubbub/structs"
	"gh-hubbub/style"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type DateModel struct {
	Tab       string
	Title     string
	fromInput textinput.Model
	toInput   textinput.Model
}

func dateValidator(s, prompt string) error {

	errMsg := fmt.Errorf("please enter a YYYY-MM-DD date for `%s`", prompt)

	// Can't be longer than 10 characters
	if len(s) > 10 {
		return errMsg
	}

	// Only dashes `-` and numbers are allowed
	withoutDashes := strings.ReplaceAll(s, "-", "")
	_, err := strconv.Atoi(withoutDashes)
	if err != nil {
		return errMsg
	}

	// Date needs to be in ISO format e.g. 2006-01-02
	if len(s) == 10 {
		_, err := time.Parse("2006-01-02", s)
		if err != nil {
			return errMsg
		}
	}

	// TODO: control the format so we can't enter e.g. 99999999 or ------

	return nil
}

func NewDateModel(tab, title string, from, to time.Time) DateModel {
	m := DateModel{
		Tab:       tab,
		Title:     title,
		fromInput: textinput.New(),
		toInput:   textinput.New(),
	}

	m.fromInput.Placeholder = from.Format("2006-01-02")
	m.fromInput.Prompt = "From:"
	m.fromInput.CharLimit = 10
	m.fromInput.Validate = func(s string) error { return dateValidator(s, m.fromInput.Prompt) }

	m.toInput.Placeholder = to.Format("2006-01-02")
	m.toInput.Prompt = "To:"
	m.toInput.CharLimit = 10
	m.toInput.Validate = func(s string) error { return dateValidator(s, m.toInput.Prompt) }

	m.fromInput.Focus()

	return m
}

func (m DateModel) Init() tea.Cmd {
	return textinput.Blink
}

func (m DateModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch msg.String() {
		case tea.KeyEnter.String():
			// TODO: validate
			return m, m.SendAddFilterMsg
		case tea.KeyTab.String():
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

func (m DateModel) View() string {
	errorText := ""
	if m.fromInput.Err != nil {
		errorText = "\n" + style.Error.Render(m.fromInput.Err.Error())
	}
	if m.toInput.Err != nil {
		errorText = "\n" + style.Error.Render(m.toInput.Err.Error())
	}
	return m.Title + " " + m.fromInput.View() + " " + m.toInput.View() + errorText
}

func (m *DateModel) Focus() tea.Cmd {
	return m.fromInput.Focus()
}

func (m *DateModel) GetValue() (time.Time, time.Time, error) {
	fromError := m.fromInput.Validate(m.fromInput.Value())
	if fromError != nil {
		return time.Time{}, time.Time{}, fromError
	}

	toError := m.toInput.Validate(m.toInput.Value())
	if toError != nil {
		return time.Time{}, time.Time{}, toError
	}

	from, _ := time.Parse("2006-01-02", m.fromInput.Value())
	to, _ := time.Parse("2006-01-02", m.toInput.Value())

	return from, to, nil
}

func (m DateModel) SendAddFilterMsg() tea.Msg {
	from, to, _ := m.GetValue()

	return messages.NewAddFilterMsg(structs.NewFilterDate(m.Tab, m.Title, from, to))
}
