package filters

import (
	"gh-hubbub/repos"
	"gh-hubbub/shared"

	"github.com/charmbracelet/bubbles/v2/textinput"
	tea "github.com/charmbracelet/bubbletea/v2"
	"github.com/charmbracelet/lipgloss/v2"
)

type FilterSearchModel struct {
	textinput  textinput.Model
	repository repos.Repository
	properties map[string]Property
}

func NewFilterSearchModel() FilterSearchModel {
	ti := textinput.New()
	ti.Placeholder = "Type to search"
	ti.Prompt = "Add filter: "
	ti.PromptStyle = shared.PromptStyle.Width(len(ti.Prompt))
	ti.Cursor.Style = shared.CursorStyle
	ti.Focus()
	ti.CharLimit = 50
	ti.SetWidth(20)
	ti.ShowSuggestions = true

	repository := repos.Repository{}

	return FilterSearchModel{
		textinput:  ti,
		repository: repository,
		properties: make(map[string]Property),
	}
}

type PropertySelectedMsg Property

func (m FilterSearchModel) Init() (tea.Model, tea.Cmd) {
	return m, tea.Batch(getFilters, textinput.Blink)
}

func (m FilterSearchModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		switch msg.String() {
		case "enter":
			_, exists := m.CurrentPropertySuggestion()
			if exists {
				return m, m.SendNextMsg
			}
			return m, nil
		}
	case filtersListMsg:
		var suggestions []string
		for _, r := range msg.Properties {
			suggestions = append(suggestions, r.Name)
			m.properties[r.Name] = Property{r.Name, r.Description, r.Type}
		}
		m.textinput.SetSuggestions(suggestions)
	}

	var cmd tea.Cmd
	m.textinput, cmd = m.textinput.Update(msg)

	return m, cmd
}

func (m FilterSearchModel) View() string {
	return lipgloss.JoinVertical(lipgloss.Left, m.textinput.View(), "\n", m.LookupDescription())
}

func (m FilterSearchModel) LookupDescription() string {
	prop, exists := m.properties[m.textinput.CurrentSuggestion()]
	if exists {
		return prop.Description
	} else {
		return ""
	}
}

func (m FilterSearchModel) CurrentPropertySuggestion() (Property, bool) {
	prop, exists := m.properties[m.textinput.CurrentSuggestion()]
	return prop, exists
}

func (m FilterSearchModel) SendNextMsg() tea.Msg {
	property, _ := m.CurrentPropertySuggestion()
	return shared.NextMessage{ModelData: property}
}

func getFilters() tea.Msg {
	rq := repos.Repository{}
	rp := repos.NewRepoConfig(rq)

	return filtersListMsg(rp)
}
