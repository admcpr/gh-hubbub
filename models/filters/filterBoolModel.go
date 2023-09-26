package filters

import (
	"gh-hubbub/messages"

	"gh-hubbub/structs"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type BoolModel struct {
	Tab   string
	Title string
	input textinput.Model
}

func NewBoolModel(tab, title string, value bool) BoolModel {
	m := BoolModel{
		Tab:   tab,
		Title: title,
		input: textinput.New(),
	}

	m.input.SetValue(structs.YesNo(value))
	m.input.Focus()

	return m
}

func (m BoolModel) Init() tea.Cmd {
	return textinput.Blink
}

func (m BoolModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch msg.String() {
		case tea.KeyEnter.String():
			return m, m.SendAddFilterMsg
		case "y", "Y":
			m.input.SetValue("Yes")
		case "n", "N":
			m.input.SetValue("No")
		}
	}

	return m, cmd
}

func (m BoolModel) View() string {
	return m.Title + " " + m.input.View()
}

func (m *BoolModel) GetValue() bool {
	return m.input.Value() == "Yes"
}

func (m *BoolModel) Focus() tea.Cmd {
	return m.input.Focus()
}

func (m BoolModel) SendAddFilterMsg() tea.Msg {
	return messages.NewAddFilterMsg(structs.NewFilterBool(m.Tab, m.Title, m.GetValue()))
}
