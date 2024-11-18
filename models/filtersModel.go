package models

import (
	"gh-hubbub/queries"
	"gh-hubbub/structs"
	"time"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
)

type AddFilterMsg structs.Filter

type FiltersModel struct {
	filterSearch tea.Model
	filterModel  tea.Model
	filtersTable table.Model
	repository   queries.Repository
	help         help.Model
	keymap       filtersearchkeymap
	properties   map[string]property
	filters      []structs.Filter
	// dateRangeFilter filters.DateModel
	// intRangeFilter  filters.IntModel
	// boolFilter      filters.BoolModel
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
	keymap := filtersearchkeymap{}

	return FiltersModel{
		filterSearch: fsm,
		filtersTable: table,
		repository:   repository,
		help:         help,
		keymap:       keymap,
		properties:   make(map[string]property),
		filters:      []structs.Filter{},
	}
}

func (m FiltersModel) Init() tea.Cmd {
	cmd := m.filterSearch.Init()
	return cmd
}

func (m FiltersModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEsc:
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
		m.filters = append(m.filters, structs.Filter(msg))
		m.filterModel = nil
		m.filterSearch = NewFilterSearchModel()
		return m, m.filterSearch.Init()
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
		return m.filterSearch.View()
	}
}

type filtersListMsg structs.RepoProperties

func (m FiltersModel) filtering() bool {
	return m.filterModel != nil
}
