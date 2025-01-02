package models

import (
	"fmt"
	"log"
	"sort"
	"strings"

	"gh-hubbub/queries"
	"gh-hubbub/structs"

	"github.com/charmbracelet/bubbles/v2/list"
	"github.com/charmbracelet/bubbles/v2/progress"
	tea "github.com/charmbracelet/bubbletea/v2"
	"github.com/charmbracelet/lipgloss/v2"
	"github.com/cli/go-gh/v2/pkg/api"
	graphql "github.com/cli/shurcooL-graphql"
)

type orgQueryMsg queries.OrganizationQuery
type repoQueryMsg queries.RepositoryQuery

type OrgModel struct {
	Title     string
	repoCount int
	repos     []structs.RepoProperties
	filters   filterMap

	repoList  list.Model
	repoModel RepoModel

	width  int
	height int

	progress progress.Model
}

func NewOrgModel(title string, width, height int) *OrgModel {
	return &OrgModel{
		Title:     title,
		width:     width,
		height:    height,
		repoModel: NewRepoModel(width/2, height),
		repoList:  list.New([]list.Item{}, simpleItemDelegate{}, width/2, height),
		progress:  progress.New(progress.WithDefaultGradient()),
	}
}

func (m *OrgModel) SetDimensions(width, height int) {
	m.width = width
	m.height = height
}

func (m *OrgModel) populateRepoList() {
	filteredRepositories := m.filters.filterRepos(m.repos)
	items := make([]list.Item, len(filteredRepositories))
	for i, repo := range filteredRepositories {
		items[i] = simpleItem(repo.Name)
	}

	sort.Slice(items, func(i, j int) bool {
		return items[i].(simpleItem) < items[j].(simpleItem)
	})

	list := list.New(items, simpleItemDelegate{}, m.width/2, m.height-2)
	list.Title = fmt.Sprintf("%s Filters: %d", m.Title, len(m.filters))
	list.Styles.Title = titleStyle
	list.SetStatusBarItemName("Repository", "Repositories")
	list.SetShowHelp(false)
	list.SetShowTitle(true)

	m.repoList = list
	m.repoModel.SelectRepo(m.repos[m.repoList.Index()])
}

func (m OrgModel) Init() (tea.Model, tea.Cmd) {
	return m, getRepoList(m.Title)
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
			m.populateRepoList()
			cmd = m.progress.SetPercent(1.0)
		} else {
			cmd = m.progress.IncrPercent(0.9 / float64(m.repoCount))
		}
		return m, cmd

	case filtersMsg:
		m.filters = filterMap(msg)
		m.populateRepoList()
		return m, nil

	case progress.FrameMsg:
		progressModel, cmd := m.progress.Update(msg)
		m.progress = progressModel.(progress.Model)
		return m, cmd

	case tea.KeyPressMsg:
		switch msg.String() {
		case "F", "f":
			return m, handleNext
		}
		switch msg.String() {
		case "f":
			return m, func() tea.Msg {
				return NextMessage{ModelData: m.filters}
			}
		case "esc":
			return m, func() tea.Msg {
				return PreviousMessage{}
			}
		case "tab", "shift+tab":
			repoModel, cmd := m.repoModel.Update(msg)
			m.repoModel = repoModel.(RepoModel)
			return m, cmd
		default:
			m.repoList, cmd = m.repoList.Update(msg)
			return m, cmd
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
	m.repoModel.SelectRepo(m.repos[m.repoList.Index()])

	var repoList = appStyle.Width(half(m.width)).Render(m.repoList.View())
	var settings = appStyle.Width(half(m.width)).Render(m.repoModel.View())
	var rightPanel = lipgloss.JoinVertical(lipgloss.Center, settings)

	var views = []string{repoList, rightPanel}

	return lipgloss.JoinHorizontal(lipgloss.Top, views...)
}

func (m OrgModel) ProgressView() string {
	pad := strings.Repeat(" ", 2)
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
