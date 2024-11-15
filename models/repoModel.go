package models

import (
	"gh-hubbub/structs"
	"gh-hubbub/style"

	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type RepoModel struct {
	repoHeader    RepoHeaderModel
	repository    structs.RepoProperties
	settingsTable table.Model
	activeTab     int
	width         int
	height        int
}

func NewRepoModel(width, height int) RepoModel {
	return RepoModel{
		repoHeader: NewRepoHeaderModel(width, []string{}, 0),
		repository: structs.RepoProperties{Properties: []structs.RepoProperty{}, PropertyGroups: map[string][]structs.RepoProperty{}},
		width:      width,
		height:     height,
	}
}

func (m *RepoModel) SetWidth(width int) {
	m.width = width
	m.repoHeader.SetWidth(width)
}

func (m *RepoModel) SetHeight(height int) {
	m.height = height
}

func (m RepoModel) Init() tea.Cmd {
	return nil
}

func (m *RepoModel) SelectRepo(repository structs.RepoProperties) {
	m.repository = repository
	key := m.repository.Keys[m.activeTab]
	m.repoHeader = NewRepoHeaderModel(m.width, m.repository.Keys, m.activeTab)
	m.settingsTable = NewSettingsTable(m.repository.PropertyGroups[key], m.width)
}

func (m *RepoModel) SelectTab(index int) {
	m.activeTab = index
	key := m.repository.Keys[index]
	m.settingsTable = NewSettingsTable(m.repository.PropertyGroups[key], m.width)
}

func (m RepoModel) Update(msg tea.Msg) (RepoModel, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyTab:
			if m.activeTab < len(m.repository.PropertyGroups)-1 {
				m.SelectTab(m.activeTab + 1)
			} else {
				m.SelectTab(0)
			}
			m.repoHeader, _ = m.repoHeader.Update(TabSelectMessage{Index: m.activeTab})
		case tea.KeyShiftTab:
			if m.activeTab > 0 {
				m.SelectTab(m.activeTab - 1)
			} else {
				m.SelectTab(len(m.repository.PropertyGroups) - 1)
			}
			m.repoHeader, _ = m.repoHeader.Update(TabSelectMessage{Index: m.activeTab})
		}
	}
	return m, cmd
}

func (m RepoModel) View() string {
	// frameWidth, frameHeight := style.Settings.GetFrameSize()
	settings := style.App.
		Width(m.width).
		// Height(5).
		Render(m.settingsTable.View())

	return lipgloss.JoinVertical(lipgloss.Left, m.repoHeader.View(), settings)
}

func NewSettingsTable(activeSettings []structs.RepoProperty, width int) table.Model {
	halfWidth := half(width)

	columns := []table.Column{
		{Title: "", Width: halfWidth},
		{Title: "", Width: halfWidth}}

	rows := make([]table.Row, len(activeSettings))
	for i, setting := range activeSettings {
		rows[i] = table.Row{setting.Name, setting.String()}
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

func handleNext() tea.Msg {
	return NextMessage{}
}
