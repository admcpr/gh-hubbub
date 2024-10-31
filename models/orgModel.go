package models

import (
	"fmt"
	"log"
	"strings"

	"gh-hubbub/queries"
	"gh-hubbub/structs"
	"gh-hubbub/style"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/progress"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/cli/go-gh/v2/pkg/api"
	graphql "github.com/cli/shurcooL-graphql"
)

const (
	padding = 2
)

type orgQueryMsg queries.OrganizationQuery
type repoQueryMsg queries.RepositoryQuery

type OrgModel struct {
	Title     string
	repoCount int
	repos     []structs.RepoProperties

	repoList  list.Model
	repoModel RepoModel

	width  int
	height int

	progress progress.Model
}

func NewOrgModel(title string, width, height int) OrgModel {
	return OrgModel{
		Title:     title,
		width:     width,
		height:    height,
		repoModel: NewRepoModel(width/2, height),
		repoList:  list.New([]list.Item{}, list.NewDefaultDelegate(), width/2, height),
		progress:  progress.New(progress.WithDefaultGradient()),
	}
}

func (m *OrgModel) SetWidth(width int) {
	m.width = width
}

func (m *OrgModel) SetHeight(height int) {
	m.height = height
}

func (m *OrgModel) UpdateRepoList() {
	filteredRepositories := m.repos
	items := make([]list.Item, len(filteredRepositories))
	for i, repo := range m.repos {
		items[i] = structs.NewListItem(repo.Name, repo.Url)
	}

	list := list.New(items, style.DefaultDelegate, m.width/2, m.height-2)
	list.Title = m.Title
	list.Styles.Title = style.Title
	list.SetStatusBarItemName("Repository", "Repositories")
	list.SetShowHelp(false)
	list.SetShowTitle(true)

	m.repoList = list
}

func (m OrgModel) Init() tea.Cmd {
	return getRepoList(m.Title)
}

func (m OrgModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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
		m.repos = append(m.repos, structs.NewRepoProperties(msg.Repository))

		if m.repoCount == len(m.repos) {
			m.UpdateRepoList()
			m.repoModel.SelectRepo(m.repos[m.repoList.Index()])
			cmd = m.progress.SetPercent(1.0)
		} else {
			cmd = m.progress.IncrPercent(0.9 / float64(m.repoCount))
		}

		return m, cmd

	case progress.FrameMsg:
		progressModel, cmd := m.progress.Update(msg)
		m.progress = progressModel.(progress.Model)
		return m, cmd

	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEsc:
			return m, handleEscape
		case tea.KeyTab, tea.KeyShiftTab:
			m.repoModel, cmd = m.repoModel.Update(msg)
			return m, cmd
		default:
			m.repoList, cmd = m.repoList.Update(msg)
			return m, cmd
		}

	// case tea.KeyMsg:
	// 	switch msg.Type {
	// 	case tea.KeyEnter:
	// 		// If we're focussed on the list and we're not filtering, we want to focus on the repo model
	// 		if m.listFocusedAndNotFiltering() {
	// 			m.focus = m.focus.Next()
	// 			return m, nil
	// 		}
	// 	switch m.focus {
	// 	case consts.FocusList:
	// 		var tabCmd tea.Cmd

	// 		m.repoList, cmd = m.repoList.Update(msg)
	// 		m.repoModel, tabCmd = m.repoModel.Update(m.NewRepoSelectMsg())
	// 		return m, tea.Batch(cmd, tabCmd)
	// 	case consts.FocusTabs, consts.FocusFilter:
	// 		m.repoModel, cmd = m.repoModel.Update(msg)
	// 	}

	default:
		m.repoList, cmd = m.repoList.Update(msg)
	}

	return m, cmd
}

func (m OrgModel) View() string {
	if m.progress.Percent() < 1 {
		return m.ProgressView()
	}
	m.repoModel.SelectRepo(m.repos[m.repoList.Index()])

	var repoList = style.App.Width(half(m.width)).Render(m.repoList.View())
	var settings = style.App.Width(half(m.width)).Render(m.repoModel.View())
	var rightPanel = lipgloss.JoinVertical(lipgloss.Center, settings)

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

func handleEscape() tea.Msg {
	return PreviousMessage{}
}
