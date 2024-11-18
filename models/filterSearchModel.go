package models

import (
	"fmt"
	"gh-hubbub/queries"
	"gh-hubbub/structs"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type FilterSearchModel struct {
	textinput  textinput.Model
	repository queries.Repository
	help       help.Model
	keymap     filtersearchkeymap
	properties map[string]property
}

func NewFilterSearchModel() FilterSearchModel {
	ti := textinput.New()
	ti.Placeholder = "Type to search"
	ti.Prompt = "Property Name: "
	ti.PromptStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("63"))
	ti.Cursor.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("63"))
	ti.Focus()
	ti.CharLimit = 50
	ti.Width = 20
	ti.ShowSuggestions = true

	repository := queries.Repository{}

	help := help.New()
	keymap := filtersearchkeymap{}

	return FilterSearchModel{
		textinput:  ti,
		repository: repository,
		help:       help,
		keymap:     keymap,
		properties: make(map[string]property),
	}
}

type PropertySelectedMsg property

func (m FilterSearchModel) Init() tea.Cmd {
	return tea.Batch(getFilters, textinput.Blink)
}

func (m FilterSearchModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter:
			_, exists := m.CurrentPropertySuggestion()
			if exists {
				return m, m.SendPropertyMsg
			}
			return m, nil
		}
	case filtersListMsg:
		var suggestions []string
		for _, r := range msg.Properties {
			suggestions = append(suggestions, r.Name)
			m.properties[r.Name] = property{r.Name, r.Description, r.Type}
		}
		m.textinput.SetSuggestions(suggestions)
	}

	var cmd tea.Cmd
	m.textinput, cmd = m.textinput.Update(msg)

	return m, cmd
}

func (m FilterSearchModel) View() string {
	sugestedType := ""
	prop, exists := m.CurrentPropertySuggestion()
	if exists {
		sugestedType = prop.Type
	}

	return fmt.Sprintf(
		"Pick a Property :\n\n%s\n\n   %s\n\n%s\n\n%s\n",
		m.textinput.View(),
		m.LookupDescription(),
		m.help.View(m.keymap),
		sugestedType,
	)
}

func (m FilterSearchModel) LookupDescription() string {
	prop, exists := m.properties[m.textinput.CurrentSuggestion()]
	if exists {
		return prop.Description
	} else {
		return ""
	}
}

func (m FilterSearchModel) CurrentPropertySuggestion() (property, bool) {
	prop, exists := m.properties[m.textinput.CurrentSuggestion()]
	return prop, exists
}

func (m FilterSearchModel) SendPropertyMsg() tea.Msg {
	property, _ := m.CurrentPropertySuggestion()
	return PropertySelectedMsg(property)
}

func getFilters() tea.Msg {
	rq := queries.Repository{}
	rp := structs.NewRepoProperties(rq)

	return filtersListMsg(rp)
}

type filtersearchkeymap struct{}

func (k filtersearchkeymap) ShortHelp() []key.Binding {
	return []key.Binding{
		key.NewBinding(key.WithKeys("tab"), key.WithHelp("tab", "complete")),
		key.NewBinding(key.WithKeys("down"), key.WithHelp("↓", "next suggestion")),
		key.NewBinding(key.WithKeys("up"), key.WithHelp("↑", "prev suggestion")),
		key.NewBinding(key.WithKeys("esc"), key.WithHelp("esc", "back")),
	}
}
func (k filtersearchkeymap) FullHelp() [][]key.Binding {
	return [][]key.Binding{k.ShortHelp()}
}
