package models

import (
	"gh-hubbub/structs"
	"sort"

	"github.com/charmbracelet/bubbles/v2/list"
	tea "github.com/charmbracelet/bubbletea/v2"
	"github.com/charmbracelet/lipgloss/v2"
)

type RepoModel struct {
	repoHeader   RepoHeaderModel
	repository   structs.RepoProperties
	settingsList list.Model
	activeTab    int
	width        int
	height       int
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
	m.settingsList = NewSettingsList(m.repository.PropertyGroups[key], m.width, m.height)
}

func (m *RepoModel) SelectTab(index int) {
	m.activeTab = index
	key := m.repository.GroupKeys[index]
	m.settingsList = NewSettingsList(m.repository.PropertyGroups[key], m.width, m.height)
}

func (m RepoModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyPressMsg:
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
	// settings := appStyle.
	// 	Width(m.width).
	// 	Render(m.settingsTable.View())

	settings := m.settingsList.View()

	return lipgloss.JoinVertical(lipgloss.Left, settings)
}

func NewSettingsList(activeSettings []structs.RepoProperty, width, height int) list.Model {
	sort.Slice(activeSettings, func(i, j int) bool {
		return activeSettings[i].Name < activeSettings[j].Name
	})

	items := make([]list.Item, len(activeSettings))

	for i, setting := range activeSettings {
		items[i] = structs.NewListItem(setting.Name, setting.String())
	}

	delegate := DefaultDelegate
	delegate.Styles.SelectedDesc = delegate.Styles.NormalDesc
	delegate.Styles.SelectedTitle = delegate.Styles.NormalTitle
	delegate.SetSpacing(0)

	list := list.New(items, delegate, width, height)
	list.SetShowFilter(false)
	list.SetShowHelp(false)
	list.SetShowPagination(false)
	list.SetShowStatusBar(false)
	list.Styles.Title = titleStyle
	list.SetHeight(height)

	return list
}

func handleNext() tea.Msg {
	return NextMessage{}
}
