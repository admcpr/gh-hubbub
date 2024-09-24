package models

import (
	"gh-hubbub/structs"
	"gh-hubbub/style"

	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type RepoModelList struct {
	repository    structs.RepositorySettings
	settingsTable table.Model

	activeTab int
	width     int
	height    int
}

func NewRepoModelList(width, height int) RepoModelList {
	return RepoModelList{
		repository: structs.RepositorySettings{},
		width:      width,
		height:     height,
	}
}

func (m *RepoModelList) SetWidth(width int) {
	m.width = width
}

func (m *RepoModelList) SetHeight(height int) {
	m.height = height
}

func (m RepoModelList) Init() tea.Cmd {
	return nil
}

func (m *RepoModelList) SelectRepo(repository structs.RepositorySettings) {
	m.repository = repository
	m.settingsTable = NewSettingsList(m.repository.SettingsTabs[m.activeTab].Settings, m.width)
}

func (m *RepoModelList) SelectTab(index int) {
	m.activeTab = index
	m.settingsTable = NewSettingsList(m.repository.SettingsTabs[m.activeTab].Settings, m.width)
}

func (m RepoModelList) Update(msg tea.Msg) (RepoModelList, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyTab:
			m.SelectTab(min(m.activeTab+1, len(m.repository.SettingsTabs)-1))
		case tea.KeyShiftTab:
			m.SelectTab(max(m.activeTab-1, 0))
		}
	}
	return m, cmd
}

func (m RepoModelList) View() string {
	if len(m.repository.SettingsTabs) == 0 {
		// Can this ever happen ????
		return ""
	}

	// frameWidth, frameHeight := style.Settings.GetFrameSize()

	settings := style.App.
		Width(m.width).
		// Height(5).
		Render(m.settingsTable.View())

	return lipgloss.JoinVertical(lipgloss.Left, settings)
}

func NewSettingsList(activeSettings []structs.Setting, width int) table.Model {
	halfWidth := half(width)

	columns := []table.Column{
		{Title: "", Width: halfWidth},
		{Title: "", Width: halfWidth}}

	rows := make([]table.Row, len(activeSettings))
	for i, setting := range activeSettings {
		rows[i] = table.Row{setting.Name, setting.String()}
	}

	// table := table.New(table.WithColumns(columns), table.WithRows(rows),
	// 	table.WithFocused(true), table.WithStyles(table.DefaultStyles()))

	table := table.New(table.WithColumns(columns), table.WithRows(rows))

	return table
}
