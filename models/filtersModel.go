package models

import (
	"gh-hubbub/queries"
	"gh-hubbub/structs"
	"time"

	"github.com/charmbracelet/bubbles/v2/help"
	"github.com/charmbracelet/bubbles/v2/key"
	"github.com/charmbracelet/bubbles/v2/table"
	tea "github.com/charmbracelet/bubbletea/v2"
)

type AddFilterMsg structs.Filter

type FiltersModel struct {
	filterSearch tea.Model
	filterModel  tea.Model
	filtersTable table.Model
	repository   queries.Repository
	help         help.Model
	keymap       filterKeyMap
	properties   map[string]property
	filters      map[string]structs.Filter
}

type property struct {
	Name        string
	Description string
	Type        string
}

func NewFiltersModel() FiltersModel {
	fsm := NewFilterSearchModel()
	table := table.New()
	repository := queries.Repository{}

	help := help.New()
	keymap := filterKeyMap{}

	return FiltersModel{
		filterSearch: fsm,
		filtersTable: table,
		repository:   repository,
		help:         help,
		keymap:       keymap,
		properties:   make(map[string]property),
		filters:      make(map[string]structs.Filter),
	}
}

func (m FiltersModel) Init() (tea.Model, tea.Cmd) {
	_, cmd := m.filterSearch.Init()
	return m, cmd
}

func (m FiltersModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			if m.filtering() {
				m.filterModel = nil
				return m, nil
			} else {
				return m, handleEscape
			}
		}
	case PropertySelectedMsg:
		m.filterModel = NewFilter(property(msg))
		return m, nil

	case AddFilterMsg:
		m.filters[msg.GetName()] = structs.Filter(msg)
		m.filterModel = nil
		m.filterSearch = NewFilterSearchModel()
		return m.filterSearch.Init()
	}

	if m.filtering() {
		m.filterModel, cmd = m.filterModel.Update(msg)
	} else {
		m.filterSearch, cmd = m.filterSearch.Update(msg)
	}

	return m, cmd
}

func NewFilter(property property) tea.Model {
	switch property.Type {
	case "bool":
		return NewBoolModel(property.Name, false)
	case "int":
		return NewIntModel(property.Name, 0, 100000)
	case "time.Time":
		return NewDateModel(property.Name, time.Time{}, time.Now())
	default:
		return nil
	}
}

func (m FiltersModel) View() string {
	if m.filtering() {
		return m.filterModel.View()
	} else {
		filterTableView := ""
		if len(m.filters) > 0 {
			m.filtersTable = NewFilterTable(m.filters, 80)
			filterTableView = m.filtersTable.View()
		}
		return m.filterSearch.View() + "\n\n" + filterTableView + "\n\n" + m.help.View(m.keymap)
	}
}

type filtersListMsg structs.RepoProperties

func NewFilterTable(filters map[string]structs.Filter, width int) table.Model {
	halfWidth := half(width)

	columns := []table.Column{
		{Title: "Name", Width: halfWidth},
		{Title: "Filter", Width: halfWidth}}

	rows := make([]table.Row, 0, len(filters))
	for _, filter := range filters {
		rows = append(rows, table.Row{filter.GetName(), filter.String()})
	}

	// Make table cells not wrap
	s := table.DefaultStyles()
	s.Cell = s.Cell.MaxWidth(halfWidth).Inline(true)

	table := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithStyles(s),
	)

	return table
}

func (m FiltersModel) filtering() bool {
	return m.filterModel != nil
}

type filterKeyMap struct{}

func (k filterKeyMap) ShortHelp() []key.Binding {
	return []key.Binding{
		key.NewBinding(key.WithKeys("tab"), key.WithHelp("tab", "complete")),
		key.NewBinding(key.WithKeys("down"), key.WithHelp("↓", "next suggestion")),
		key.NewBinding(key.WithKeys("up"), key.WithHelp("↑", "prev suggestion")),
		key.NewBinding(key.WithKeys("esc"), key.WithHelp("esc", "back")),
	}
}
func (k filterKeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{k.ShortHelp()}
}
