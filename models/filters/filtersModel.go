package filters

import (
	"gh-hubbub/structs"

	tea "github.com/charmbracelet/bubbletea"
)

type FiltersModel struct {
	Filters []structs.Filter
}

func (m FiltersModel) Init() tea.Cmd {
	return nil
}

func (m FiltersModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m, nil
}

func (m FiltersModel) View() string {
	text := ""

	for _, filter := range m.Filters {
		text += filter.String() + " "
	}

	return text
}
