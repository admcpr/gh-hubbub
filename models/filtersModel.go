package models

import (
	"fmt"
	"gh-hubbub/queries"
	"gh-hubbub/structs"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/table"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type FiltersModel struct {
	textinput    textinput.Model
	filtersTable table.Model
	repository   queries.Repository
	help         help.Model
	keymap       keymap
	descriptions map[string]string
}

func NewFiltersModel() FiltersModel {
	ti := textinput.New()
	ti.Placeholder = "property name"
	ti.Prompt = "Search "
	ti.PromptStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("63"))
	ti.Cursor.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("63"))
	ti.Focus()
	ti.CharLimit = 50
	ti.Width = 20
	ti.ShowSuggestions = true

	table := table.New()
	repository := queries.Repository{}

	help := help.New()
	keymap := keymap{}

	return FiltersModel{
		textinput:    ti,
		filtersTable: table,
		repository:   repository,
		help:         help,
		keymap:       keymap,
		descriptions: make(map[string]string),
	}
}

type keymap struct{}

func (k keymap) ShortHelp() []key.Binding {
	return []key.Binding{
		key.NewBinding(key.WithKeys("tab"), key.WithHelp("tab", "complete")),
		key.NewBinding(key.WithKeys("down"), key.WithHelp("↓", "next")),
		key.NewBinding(key.WithKeys("up"), key.WithHelp("↑", "prev")),
		key.NewBinding(key.WithKeys("esc"), key.WithHelp("esc", "quit")),
	}
}
func (k keymap) FullHelp() [][]key.Binding {
	return [][]key.Binding{k.ShortHelp()}
}

func (m FiltersModel) Init() tea.Cmd {
	// TODO: get all possible filters for the repository
	return tea.Batch(getFilters, textinput.Blink)
}

func (m FiltersModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter:
			return m, nil
		case tea.KeyEsc:
			return m, handleEscape
		}
	case filtersListMessage:
		var suggestions []string
		for _, r := range msg.Properties {
			suggestions = append(suggestions, r.Name)
			m.descriptions[r.Name] = r.Description
		}
		m.textinput.SetSuggestions(suggestions)
	}

	var cmd tea.Cmd
	m.textinput, cmd = m.textinput.Update(msg)
	return m, cmd
}

func (m FiltersModel) View() string {
	return fmt.Sprintf(
		"Pick a Property :\n\n%s\n\n   %s\n\n%s\n\n",
		m.textinput.View(),
		m.LookupDescription(),
		m.help.View(m.keymap),
	)
}

type filtersListMessage structs.RepoProperties

func (m FiltersModel) LookupDescription() string {
	return m.descriptions[m.textinput.CurrentSuggestion()]
}

func getFilters() tea.Msg {
	rq := queries.Repository{}
	rp := structs.NewRepoProperties(rq)

	return filtersListMessage(rp)
}
