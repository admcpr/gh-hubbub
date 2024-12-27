package models

import (
	"gh-hubbub/structs"
	"gh-hubbub/style"

	"github.com/charmbracelet/bubbles/v2/table"
	tea "github.com/charmbracelet/bubbletea/v2"
	"github.com/charmbracelet/lipgloss/v2"
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
		repository: structs.RepoProperties{Properties: map[string]structs.RepoProperty{}, PropertyGroups: map[string][]structs.RepoProperty{}},
		width:      width,
		height:     height,
	}
}

func (m *RepoModel) SetDimensions(width, height int) {
	m.width = width
	m.height = height
	m.repoHeader.SetDimensions(width, height)
}

func (m RepoModel) Init() (tea.Model, tea.Cmd) {
	return m, nil
}

func (m *RepoModel) SelectRepo(repository structs.RepoProperties) {
	m.repository = repository
	key := m.repository.GroupKeys[m.activeTab]
	m.repoHeader = NewRepoHeaderModel(m.width, m.repository.GroupKeys, m.activeTab)
	m.settingsTable = NewSettingsTable(m.repository.PropertyGroups[key], m.width)
}

func (m *RepoModel) SelectTab(index int) {
	m.activeTab = index
	key := m.repository.GroupKeys[index]
	m.settingsTable = NewSettingsTable(m.repository.PropertyGroups[key], m.width)
}

func (m RepoModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyReleaseMsg:
		switch msg.String() {
		case "tab":
			if m.activeTab < len(m.repository.PropertyGroups)-1 {
				m.SelectTab(m.activeTab + 1)
			} else {
				m.SelectTab(0)
			}
			repoHeader, _ := m.repoHeader.Update(TabSelectMessage{Index: m.activeTab})
			m.repoHeader = repoHeader.(RepoHeaderModel)
		case "shift+tab":
			if m.activeTab > 0 {
				m.SelectTab(m.activeTab - 1)
			} else {
				m.SelectTab(len(m.repository.PropertyGroups) - 1)
			}
			repoHeader, _ := m.repoHeader.Update(TabSelectMessage{Index: m.activeTab})
			m.repoHeader = repoHeader.(RepoHeaderModel)
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
