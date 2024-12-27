package models

import (
	"gh-hubbub/queries"
	"gh-hubbub/structs"
	"gh-hubbub/style"
	"sort"
	"time"

	"github.com/charmbracelet/bubbles/v2/help"
	"github.com/charmbracelet/bubbles/v2/key"
	"github.com/charmbracelet/bubbles/v2/list"
	tea "github.com/charmbracelet/bubbletea/v2"
	"github.com/charmbracelet/lipgloss/v2"
)

type AddFilterMsg structs.Filter
type filterMap map[string]structs.Filter

type FiltersModel struct {
	filterSearch tea.Model
	filterModel  tea.Model
	filtersList  list.Model
	repository   queries.Repository
	help         help.Model
	keymap       filterKeyMap
	properties   map[string]property
	filters      filterMap
	width        int
	height       int
}

func (m FiltersModel) SetDimensions(width, height int) {
	m.width = width
	m.height = height
}

type property struct {
	Name        string
	Description string
	Type        string
}

func NewFiltersModel(width, height int) FiltersModel {
	fsm := NewFilterSearchModel()
	list := list.New([]list.Item{}, simpleItemDelegate{}, width, height-4)
	repository := queries.Repository{}

	help := help.New()
	keymap := filterKeyMap{}

	return FiltersModel{
		filterSearch: fsm,
		filtersList:  list,
		repository:   repository,
		help:         help,
		keymap:       keymap,
		properties:   make(map[string]property),
		filters:      make(map[string]structs.Filter),
		width:        width,
		height:       height,
	}
}

func (m FiltersModel) Init() (tea.Model, tea.Cmd) {
	_, cmd := m.filterSearch.Init()
	return m, cmd
}

func (m FiltersModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyReleaseMsg:
		switch msg.String() {
		case "esc":
			if m.filtering() {
				m.filterModel = nil
				return m, nil
			} else {
				return m, func() tea.Msg {
					return PreviousMessage{ModelData: m.filters}
				}
			}
		case "ctrl+enter":
			return m, func() tea.Msg {
				return PreviousMessage{ModelData: m.filters}
			}
		}
	case PropertySelectedMsg:
		m.filterModel = NewFilter(property(msg))
		return m, nil

	case AddFilterMsg:
		m.filters[msg.GetName()] = structs.Filter(msg)
		m.filterModel = nil
		m.filterSearch = NewFilterSearchModel()
		m.filterSearch, cmd = m.filterSearch.Init()
		return m, cmd
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
		filtersListView := ""
		if len(m.filters) > 0 {
			m.filtersList = NewFiltersList(m.filters, m.width, m.height)
			filtersListView = m.filtersList.View()
		}
		search := m.filterSearch.View()
		help := m.help.View(m.keymap)
		return lipgloss.JoinVertical(lipgloss.Left, search, filtersListView, help)
	}
}

type filtersListMsg structs.RepoProperties

func NewFiltersList(filters map[string]structs.Filter, width, height int) list.Model {
	items := make([]list.Item, len(filters))
	i := 0
	for _, filter := range filters {
		items[i] = simpleItem(filter.GetName())
		i++
	}

	sort.Slice(items, func(i, j int) bool {
		return items[i].(simpleItem) < items[j].(simpleItem)
	})

	list := list.New(items, simpleItemDelegate{}, width, height-4)
	list.Styles.Title = style.Title
	list.Title = "Filters"
	list.SetShowHelp(false)
	list.SetShowTitle(true)

	return list
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
