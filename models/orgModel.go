package models

import (
	"fmt"
	"log"
	"strings"

	"gh-hubbub/consts"
	"gh-hubbub/filters"
	"gh-hubbub/keyMaps"
	"gh-hubbub/queries"
	"gh-hubbub/structs"
	"gh-hubbub/style"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/progress"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/cli/go-gh/v2/pkg/api"
	graphql "github.com/cli/shurcooL-graphql"
)

const (
	padding  = 2
	maxWidth = 80
)

type orgQueryMsg queries.OrganizationQuery
type repoQueryMsg queries.RepositoryQuery
type FocusMsg struct{ Focus consts.Focus }

func NewFocusMsg(focus consts.Focus) FocusMsg {
	return FocusMsg{Focus: focus}
}

type OrgModel struct {
	Title   string
	Filters []filters.Filter

	repoCount int
	repos     []structs.RepositorySettings

	repoList  list.Model
	repoModel RepoModel
	help      help.Model
	keys      keyMaps.OrgKeyMap

	// Focus is the current focus of the model
	// We should just be using a state machine here
	focus  consts.Focus
	width  int
	height int

	progress progress.Model
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
		Filters:   []filters.Filter{},
		progress:  progress.New(progress.WithDefaultGradient()),
	}
}

func (m *OrgModel) FilteredRepositories() []structs.RepositorySettings {
	if len(m.Filters) == 0 {
		return m.repos
	}
	filteredRepos := []structs.RepositorySettings{}
	for _, repo := range m.repos {
		if RepoMatchesFilters(repo, m.Filters) {
			filteredRepos = append(filteredRepos, repo)
		}
	}
	return filteredRepos
}

func RepoMatchesFilters(repo structs.RepositorySettings, filters []filters.Filter) bool {
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

func getTitle(t string, filters []filters.Filter) string {
	title := "Organization: " + t
	if len(filters) > 0 {
		return fmt.Sprintf("%s (%s)", title, filters[0].String())
	}
	return title
}

func (m *OrgModel) helpView() string {
	return m.help.View(m.keys)
}

func (m OrgModel) Init() tea.Cmd {
	return getRepoList(m.Title)
}

func (m OrgModel) Update(msg tea.Msg) (OrgModel, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case orgQueryMsg:
		repos := msg.Organization.Repositories.Nodes
		cmds := []tea.Cmd{m.progress.SetPercent(0.1)}
		m.repoCount = len(msg.Organization.Repositories.Nodes)
		for _, repo := range repos {
			cmds = append(cmds, getRepoDetails(m.Title, repo.Name))
		}
		return m, tea.Batch(cmds...)

	case repoQueryMsg:
		m.repos = append(m.repos, structs.NewRepository(msg.Repository))

		if m.repoCount == len(m.repos) {
			m.UpdateRepoList()
			m.repoModel.SelectRepo(m.repos[m.repoList.Index()], m.width, m.height)
			cmd = m.progress.SetPercent(1.0)
		} else {
			cmd = m.progress.IncrPercent(0.9 / float64(m.repoCount))
		}

		return m, cmd

	case progress.FrameMsg:
		progressModel, cmd := m.progress.Update(msg)
		m.progress = progressModel.(progress.Model)
		return m, cmd

	case FocusMsg:
		m.focus = msg.Focus

	// case tea.KeyMsg:
	// 	switch msg.Type {
	// 	case tea.KeyEnter:
	// 		// If we're focussed on the list and we're not filtering, we want to focus on the repo model
	// 		if m.listFocusedAndNotFiltering() {
	// 			m.focus = m.focus.Next()
	// 			return m, nil
	// 		}
	// 	case tea.KeyEsc:
	// 		// Esc goes back so go to the previous model if we're focussed on the list
	// 		if m.listFocusedAndNotFiltering() {
	// 			return MainModel[consts.UserModelName], nil
	// 		}
	// 	case tea.KeyCtrlC:
	// 		return m, tea.Quit
	// 	}
	// 	switch m.focus {
	// 	case consts.FocusList:
	// 		var tabCmd tea.Cmd

	// 		m.repoList, cmd = m.repoList.Update(msg)
	// 		m.repoModel, tabCmd = m.repoModel.Update(m.NewRepoSelectMsg())
	// 		return m, tea.Batch(cmd, tabCmd)
	// 	case consts.FocusTabs, consts.FocusFilter:
	// 		m.repoModel, cmd = m.repoModel.Update(msg)
	// 	}
	case filters.FilterMsg:
		switch msg.Action {
		case consts.FilterDelete:
			// Remove the filter from the list
			for i, filter := range m.Filters {
				if filter.GetTab() == msg.Filter.GetTab() && filter.GetName() == msg.Filter.GetName() {
					m.Filters = append(m.Filters[:i], m.Filters[i+1:]...)
				}
			}
		case consts.FilterAdd:
			m.Filters = []filters.Filter{msg.Filter}
			m.UpdateRepoList()
			m.repoModel, cmd = m.repoModel.Update(filters.NewConfirmFilterMsg(nil))
		}

	default:
		m.repoList, cmd = m.repoList.Update(msg)
	}

	return m, cmd
}

func (m OrgModel) View() string {
	if m.progress.Percent() < 1 {
		return m.ProgressView()
	}
	m.repoModel.SelectRepo(m.repos[m.repoList.Index()], m.width, m.height)
	var repoList = style.App.Width(half(m.width)).Render(m.repoList.View())
	var settings = style.App.Width(half(m.width)).Render(m.repoModel.View())
	help := m.helpView()
	var rightPanel = lipgloss.JoinVertical(lipgloss.Center, settings, help)

	var views = []string{repoList, rightPanel}

	return lipgloss.JoinHorizontal(lipgloss.Top, views...)
}

func (m OrgModel) ProgressView() string {
	pad := strings.Repeat(" ", padding)
	progress := "\n" + pad + m.progress.View() + "\n\n" + pad + "Getting repositories ... "

	if m.repoCount < 1 {
		return progress
	}
	return progress + fmt.Sprintf("%d of %d", len(m.repos), m.repoCount)
}

func getRepoDetails(owner string, name string) tea.Cmd {
	return func() tea.Msg {
		client, err := api.DefaultGraphQLClient()
		if err != nil {
			log.Fatal(err)
		}
		repoQuery := queries.RepositoryQuery{}

		variables := map[string]interface{}{
			"owner": graphql.String(owner),
			"name":  graphql.String(name),
			// "first": graphql.Int(100),
		}
		err = client.Query("Repository", &repoQuery, variables)
		if err != nil {
			log.Fatal(err)
		}

		return repoQueryMsg(repoQuery)
	}
}

func getRepoList(login string) tea.Cmd {
	return func() tea.Msg {
		client, err := api.DefaultGraphQLClient()
		if err != nil {
			return AuthenticationErrorMsg{Err: err}
		}

		var organizationQuery = queries.OrganizationQuery{}

		variables := map[string]interface{}{
			"login": graphql.String(login),
			"first": graphql.Int(100),
		}
		err = client.Query("OrganizationRepositories", &organizationQuery, variables)
		if err != nil {
			log.Fatal(err)
		}

		return orgQueryMsg(organizationQuery)
	}
}
