package models

import (
	"gh-hubbub/consts"
	"gh-hubbub/filters"
	"gh-hubbub/keyMaps"
	"gh-hubbub/structs"
	"gh-hubbub/style"
	"time"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type RepoModel struct {
	repository  structs.RepositorySettings
	FilterModel tea.Model

	settingsTable table.Model

	help help.Model
	keys keyMaps.RepoKeyMap

	activeTab int
	focus     consts.Focus
	loaded    bool
	width     int
	height    int
}

func NewRepoModel(width, height int) RepoModel {
	return RepoModel{
		repository: structs.RepositorySettings{},
		width:      width,
		height:     height,
		help:       help.New(),
		keys:       keyMaps.NewRepoKeyMap(),
	}
}

func (m RepoModel) Init() tea.Cmd {
	return nil
}

func (m *RepoModel) filterHasFocus() bool {
	return m.focus == consts.FocusFilter
}

func (m *RepoModel) SelectRepo(repository structs.RepositorySettings, width, height int) {
	m.repository = repository
	m.settingsTable = NewSettingsTable(m.repository.SettingsTabs[m.activeTab].Settings, width)

	m.width = width
	m.height = height
}

func (m *RepoModel) SelectTab(index int) {
	m.activeTab = index
	m.settingsTable = NewSettingsTable(m.repository.SettingsTabs[m.activeTab].Settings, m.width)
}

func (m *RepoModel) InitFilterEditor() {
	tab := m.repository.SettingsTabs[m.activeTab]
	index := m.settingsTable.Cursor()
	setting := tab.Settings[index]

	switch value := setting.Value.(type) {
	case bool:
		m.FilterModel = filters.NewBoolModel(tab.Name, setting.Name, value)
	case int:
		m.FilterModel = filters.NewIntModel(tab.Name, setting.Name, value, value)
	case time.Time:
		m.FilterModel = filters.NewDateModel(tab.Name, setting.Name, value, value)
	}
}

func (m RepoModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		if !m.loaded {
			m.loaded = true
		}
		return m, nil
	case RepoSelectMsg:
		m.SelectRepo(msg.Repository, msg.Width, msg.Height)
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEsc:
			return m, m.FocusList
		default:
			if m.filterHasFocus() {
				m.FilterModel, cmd = m.FilterModel.Update(msg)
			} else {
				m, cmd = m.UpdateRepoModel(msg.Type)
			}
		}
	case FocusMsg:
		m.focus = msg.Focus
	case filters.FilterMsg:
		switch msg.Action {
		case consts.FilterConfirm:
			m.focus = consts.FocusTabs
			cmd = m.SendFocusMsg
		}
	}

	return m, cmd
}

func (m RepoModel) UpdateRepoModel(keyType tea.KeyType) (RepoModel, tea.Cmd) {
	var cmd tea.Cmd

	switch keyType {
	case tea.KeyEnter:
		if !m.filterHasFocus() {
			m.InitFilterEditor()
			m.focus = consts.FocusFilter
			cmd = m.SendFocusMsg
		}
	case tea.KeyEsc:
		cmd = m.FocusList
	case tea.KeyRight:
		m.SelectTab(min(m.activeTab+1, len(m.repository.SettingsTabs)-1))
	case tea.KeyLeft:
		m.SelectTab(max(m.activeTab-1, 0))
	case tea.KeyDown:
		m.settingsTable.MoveDown(1)
	case tea.KeyUp:
		m.settingsTable.MoveUp(1)
	}

	return m, cmd
}

func (m RepoModel) View() string {
	if m.repository.SettingsTabs == nil || len(m.repository.SettingsTabs) == 0 {
		// Can this ever happen ????
		return ""
	}

	var tabs = RenderTabs(m.repository.SettingsTabs, m.width, m.activeTab)
	if m.filterHasFocus() {
		filter := lipgloss.NewStyle().Width(m.width - 2).Height(m.height - 7).Render(m.FilterModel.View())
		return lipgloss.JoinVertical(lipgloss.Left, tabs, filter)
	} else {
		settings := style.Settings.Width(m.width - 2).Height(m.height - 7).Render(m.settingsTable.View())
		return lipgloss.JoinVertical(lipgloss.Left, tabs, settings)
	}
}

func NewSettingsTable(activeSettings []structs.Setting, width int) table.Model {
	widthWithoutBorder := width - 2
	quarterWidth := quarter(widthWithoutBorder)

	columns := []table.Column{
		{Title: "Setting", Width: (widthWithoutBorder - quarterWidth)},
		{Title: "Value", Width: quarterWidth}}

	rows := make([]table.Row, len(activeSettings))
	for i, setting := range activeSettings {
		rows[i] = table.Row{setting.Name, setting.String()}
	}

	return table.New(table.WithColumns(columns),
		table.WithRows(rows), table.WithFocused(true), table.WithStyles(GetTableStyles()))
}

func GetTableStyles() table.Styles {
	return table.Styles{
		Selected: style.TableSelected,
		Header:   style.TableHeader,
		Cell:     style.TableCell,
	}
}

func (m *RepoModel) SendFocusMsg() tea.Msg {
	return NewFocusMsg(m.focus)
}

func (m *RepoModel) FocusList() tea.Msg {
	m.focus = consts.FocusList
	return m.SendFocusMsg()
}
