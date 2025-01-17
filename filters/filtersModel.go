package models

import (
	"gh-hubbub/queries"
	"gh-hubbub/structs"
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
	filtersList  list.Model
	repository   queries.Repository
	help         help.Model
	keymap       filterKeyMap
	properties   map[string]property
	filters      filterMap
	width        int
	height       int
}

func (m *FiltersModel) SetDimensions(width, height int) {
	m.width = width
	m.height = height
}

type property struct {
	Name        string
	Description string
	Type        string
}

func NewFiltersModel(width, height int) *FiltersModel {
	fsm := NewFilterSearchModel()
	list := list.New([]list.Item{}, simpleItemDelegate{}, width, height-4)
	repository := queries.Repository{}

	help := help.New()
	keymap := filterKeyMap{}

	return &FiltersModel{
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
	case tea.KeyPressMsg:
		switch msg.String() {
		case "esc":
			return m, func() tea.Msg {
				return PreviousMessage{ModelData: m.filters}
			}
		case "ctrl+enter":
			return m, func() tea.Msg {
				return PreviousMessage{ModelData: m.filters}
			}
		}

	case AddFilterMsg:
		m.filters[msg.GetName()] = structs.Filter(msg)
		m.filterSearch = NewFilterSearchModel()
		m.filterSearch, cmd = m.filterSearch.Init()
		return m, cmd
	}

	m.filterSearch, cmd = m.filterSearch.Update(msg)

	return m, cmd
}

func NewFilter(property property, width, height int) tea.Model {
	switch property.Type {
	case "bool":
		return NewBoolModel(property.Name, false, width, height)
	case "int":
		return NewIntModel(property.Name, 0, 100000, width, height)
	case "time.Time":
		return NewDateModel(property.Name, time.Time{}, time.Now(), width, height)
	default:
		return nil
	}
}

func (m FiltersModel) View() string {
	m.filtersList = NewFiltersList(m.filters, m.width, m.height)
	filtersListView := m.filtersList.View()

	search := m.filterSearch.View()
	help := m.help.View(m.keymap)
	return lipgloss.JoinVertical(lipgloss.Left, filtersListView, search, help)
	// }
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

	list := list.New(items, simpleItemDelegate{}, width, height-8)
	list.Styles.Title = titleStyle
	list.Title = "Filters"
	list.SetShowHelp(false)
	list.SetShowStatusBar(false)
	list.SetShowTitle(true)

	return list
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
