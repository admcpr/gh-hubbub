package models

import (
	"gh-hubbub/consts"
	"gh-hubbub/filters"
	"gh-hubbub/structs"
	"gh-hubbub/style"
	"time"

	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type RepoModel struct {
	repository  structs.RepositorySettings
	FilterModel tea.Model

	settingsTable table.Model

	activeTab int
	focus     consts.Focus
	width     int
	height    int
}

func NewRepoModel(width, height int) RepoModel {
	return RepoModel{
		repository: structs.RepositorySettings{},
		width:      width,
		height:     height,
	}
}

func (m *RepoModel) SetWidth(width int) {
	m.width = width
}

func (m *RepoModel) SetHeight(height int) {
	m.height = height
}

func (m RepoModel) Init() tea.Cmd {
	return nil
}

func (m *RepoModel) filterHasFocus() bool {
	return m.focus == consts.FocusFilter
}

func (m *RepoModel) SelectRepo(repository structs.RepositorySettings) {
	m.repository = repository
	m.settingsTable = NewSettingsTable(m.repository.SettingsTabs[m.activeTab].Settings, m.width)
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

func (m RepoModel) Update(msg tea.Msg) (RepoModel, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
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
	}
	return m, cmd
}

func (m RepoModel) UpdateRepoModel(keyType tea.KeyType) (RepoModel, tea.Cmd) {
	var cmd tea.Cmd

	switch keyType {
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
	if len(m.repository.SettingsTabs) == 0 {
		// Can this ever happen ????
		return ""
	}

	frameWidth, frameHeight := style.Settings.GetFrameSize()
	var tabs = RenderTabs(m.repository.SettingsTabs, m.width, m.activeTab)

	// return style.Settings.Width(m.width - frameWidth).Height(m.height - frameHeight).Render(m.settingsTable.View())

	// if m.filterHasFocus() {
	// 	filter := lipgloss.NewStyle().Width(m.width - 2).Height(20).Render(m.FilterModel.View())
	// 	// filter := lipgloss.NewStyle().Width(m.width - 2).Height(m.height - 7).Render(m.FilterModel.View())
	// 	return lipgloss.JoinVertical(lipgloss.Left, tabs, filter)
	// } else {
	// 	settings := style.Settings.Width(m.width - 2).Height(20).Render(m.settingsTable.View())
	settings := style.Settings.Width(m.width - frameWidth).Height(m.height - frameHeight - 2).Render(m.settingsTable.View())
	return lipgloss.JoinVertical(lipgloss.Left, tabs, settings)
	// }
}

func NewSettingsTable(activeSettings []structs.Setting, width int) table.Model {
	halfWidth := half(width) - 1

	columns := []table.Column{
		{Title: "", Width: halfWidth},
		{Title: "", Width: halfWidth}}

	rows := make([]table.Row, len(activeSettings))
	for i, setting := range activeSettings {
		rows[i] = table.Row{setting.Name, setting.String()}
	}

	table := table.New(table.WithColumns(columns), table.WithRows(rows),
		table.WithFocused(true), table.WithStyles(table.DefaultStyles()))

	table.SetHeight(10)

	return table
}

func (m *RepoModel) SendFocusMsg() tea.Msg {
	return NewFocusMsg(m.focus)
}

func (m *RepoModel) FocusList() tea.Msg {
	m.focus = consts.FocusList
	return m.SendFocusMsg()
}
