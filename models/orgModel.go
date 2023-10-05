package models

import (
	"fmt"
	"log"

	"gh-hubbub/consts"
	"gh-hubbub/keyMaps"
	"gh-hubbub/messages"
	"gh-hubbub/structs"
	"gh-hubbub/style"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/cli/go-gh/v2/pkg/api"
	graphql "github.com/cli/shurcooL-graphql"
)

type OrgModel struct {
	Title   string
	Filters []structs.Filter

	Repositories []structs.Repository

	repoList  list.Model
	repoModel tea.Model
	help      help.Model
	keys      keyMaps.OrgKeyMap

	// Focus is the current focus of the model
	// We should just be using a state machine here
	focus   consts.Focus
	width   int
	height  int
	loaded  bool
	getting bool
}

func (m OrgModel) NewRepoSelectMsg() messages.RepoSelectMsg {
	return messages.RepoSelectMsg{
		Repository: m.Repositories[m.repoList.Index()],
		Width:      m.width / 2,
		Height:     m.height,
	}
}

func NewOrgModel(title string, width, height int) OrgModel {
	return OrgModel{
		Title:     title,
		width:     width,
		height:    height,
		help:      help.New(),
		keys:      keyMaps.NewOrgKeyMap(),
		repoModel: NewRepoModel(width/2, height),
		repoList:  list.New([]list.Item{}, list.NewDefaultDelegate(), width/2, height),
		Filters:   []structs.Filter{},
		getting:   true,
	}
}

func (m *OrgModel) FilteredRepositories() []structs.Repository {
	if len(m.Filters) == 0 {
		return m.Repositories
	}
	filteredRepos := []structs.Repository{}
	for _, repo := range m.Repositories {
		if RepoMatchesFilters(repo, m.Filters) {
			filteredRepos = append(filteredRepos, repo)
		}
	}
	return filteredRepos
}

func RepoMatchesFilters(repo structs.Repository, filters []structs.Filter) bool {
	// TODO: This is gonna get slow, fast, for big orgs. Faster pls.
	// TODO: Obviously this is also buggy if there are multiple filters, it'll only check the first one
	for _, filter := range filters {
		for _, tab := range repo.SettingsTabs {
			if tab.Name == filter.GetTab() {
				for _, setting := range tab.Settings {
					if setting.Name == filter.GetName() {
						return filter.Matches(setting)
					}
				}
			}
		}
	}
	return false
}

func (m *OrgModel) UpdateRepositories(oq structs.OrganizationQuery) {
	edges := oq.Organization.Repositories.Edges
	m.Repositories = make([]structs.Repository, len(edges))
	items := make([]list.Item, len(edges))
	for i, repoQuery := range edges {
		repo := structs.NewRepository(repoQuery.Node)
		m.Repositories[i] = repo
		items[i] = structs.NewListItem(repo.Name, repo.Url)
	}

	m.UpdateRepoList()
	m.getting = false
}

func (m *OrgModel) UpdateRepoList() {
	filteredRepositories := m.FilteredRepositories()
	items := make([]list.Item, len(filteredRepositories))
	for i, repo := range m.FilteredRepositories() {
		items[i] = structs.NewListItem(repo.Name, repo.Url)
	}

	list := list.New(items, style.DefaultDelegate, m.width, m.height-2)
	list.Title = getTitle(m.Title, m.Filters)
	list.Styles.Title = style.Title
	list.SetStatusBarItemName("Repository", "Repositories")
	list.SetShowHelp(false)
	list.SetShowTitle(true)

	m.repoList = list
}

func getTitle(t string, filters []structs.Filter) string {
	title := "Organization: " + t
	if len(filters) > 0 {
		return fmt.Sprintf("%s (%s)", title, filters[0].String())
	}
	return title
}

func (m *OrgModel) helpView() string {
	return m.help.View(m.keys)
}

func (m *OrgModel) listFocusedAndNotFiltering() bool {
	return m.focus == consts.FocusList && !m.repoList.SettingFilter()
}

func (m OrgModel) Init() tea.Cmd {
	return nil
}

func (m OrgModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {

	// Window size changed
	case tea.WindowSizeMsg:
		if !m.loaded {
			m.loaded = true
		}

	case messages.FocusMsg:
		m.focus = msg.Focus

	case messages.RepoListMsg:
		m.UpdateRepositories(msg.OrganizationQuery)
		m.repoModel, _ = m.repoModel.Update(m.NewRepoSelectMsg())

	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter:
			// If we're focussed on the list and we're not filtering, we want to focus on the repo model
			if m.listFocusedAndNotFiltering() {
				m.focus = m.focus.Next()
				return m, nil
			}
		case tea.KeyEsc:
			// Esc goes back so go to the previous model if we're focussed on the list
			if m.listFocusedAndNotFiltering() {
				return MainModel[consts.UserModelName], nil
			}
		case tea.KeyCtrlC:
			return m, tea.Quit
		}
		switch m.focus {
		case consts.FocusList:
			var tabCmd tea.Cmd

			m.repoList, cmd = m.repoList.Update(msg)
			m.repoModel, tabCmd = m.repoModel.Update(m.NewRepoSelectMsg())
			return m, tea.Batch(cmd, tabCmd)
		case consts.FocusTabs, consts.FocusFilter:
			m.repoModel, cmd = m.repoModel.Update(msg)
		}
	case messages.FilterMsg:
		switch msg.Action {
		case consts.FilterDelete:
			// Remove the filter from the list
			for i, filter := range m.Filters {
				if filter.GetTab() == msg.Filter.GetTab() && filter.GetName() == msg.Filter.GetName() {
					m.Filters = append(m.Filters[:i], m.Filters[i+1:]...)
				}
			}
		case consts.FilterAdd:
			m.Filters = []structs.Filter{msg.Filter}
			m.UpdateRepoList()
			m.repoModel, cmd = m.repoModel.Update(messages.NewConfirmFilterMsg(nil))
		}
	}

	return m, cmd
}

func (m OrgModel) View() string {
	if m.getting {
		return "getting repos ..."
	}

	var repoList = style.App.Width(half(m.width)).Render(m.repoList.View())
	var settings = style.App.Width(half(m.width)).Render(m.repoModel.View())
	help := m.helpView()
	var rightPanel = lipgloss.JoinVertical(lipgloss.Center, settings, help)

	var views = []string{repoList, rightPanel}

	return lipgloss.JoinHorizontal(lipgloss.Top, views...)
}

func (m OrgModel) GetRepositories() tea.Msg {
	client, err := api.DefaultGraphQLClient()
	if err != nil {
		return messages.AuthenticationErrorMsg{Err: err}
	}

	var organizationQuery = structs.OrganizationQuery{}

	variables := map[string]interface{}{
		"login": graphql.String(m.Title),
		"first": graphql.Int(100),
	}
	err = client.Query("OrganizationRepositories", &organizationQuery, variables)
	if err != nil {
		log.Fatal(err)
	}

	return messages.RepoListMsg{OrganizationQuery: organizationQuery}
}
