package user

import (
	"fmt"
	"sort"

	"gh-hubbub/orgs"
	"gh-hubbub/shared"

	"github.com/charmbracelet/bubbles/v2/list"
	tea "github.com/charmbracelet/bubbletea/v2"
	"github.com/cli/go-gh/v2/pkg/api"
)

type ErrMsg struct{ Err error }
type OrgListMsg struct{ Organisations []orgs.Organisation }

type UserModel struct {
	organisations  []orgs.Organisation
	User           User
	SelectedOrgUrl string
	list           list.Model
	width          int
	height         int
}

func NewUserModel(user User, width, height int) UserModel {
	userList := list.New([]list.Item{}, shared.DefaultDelegate, width, height)

	userList.Title = "User: " + user.Name
	userList.SetStatusBarItemName("Organization", "Organizations")
	userList.Styles.Title = shared.TitleStyle
	userList.SetShowTitle(true)

	return UserModel{User: user, list: userList, width: width, height: height}
}

func (m *UserModel) SetDimensions(width, height int) {
	m.width = width
	m.height = height
}

func (m UserModel) Init() (tea.Model, tea.Cmd) {
	return m, getOrganisations
}

func (m UserModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case OrgListMsg:
		m.organisations = msg.Organisations
		sort.Slice(m.organisations, func(i, j int) bool {
			return m.organisations[i].Login < m.organisations[j].Login
		})

		items := make([]list.Item, len(m.organisations))
		for i, org := range m.organisations {
			items[i] = shared.NewListItem(org.Login, org.Url)
		}

		cmd = m.list.SetItems(items)

		return m, cmd
	case tea.KeyPressMsg:
		switch msg.String() {
		case "enter":
			selectedOrg := m.organisations[m.list.Index()].Login
			cmd = func() tea.Msg {
				return shared.NextMessage{ModelData: selectedOrg}
			}
			return m, cmd
		default:
			m.list, cmd = m.list.Update(msg)
			return m, cmd
		}
	}

	m.list, cmd = m.list.Update(msg)

	return m, cmd
}

func (m UserModel) View() string {
	m.list.SetWidth(m.width)
	m.list.SetHeight(m.height)
	return shared.AppStyle.Render(m.list.View())
}

func (m UserModel) SelectedOrg() orgs.Organisation {
	return m.organisations[m.list.Index()]
}

func getOrganisations() tea.Msg {
	client, err := api.DefaultRESTClient()
	if err != nil {
		return AuthenticationErrorMsg{Err: err}
	}
	response := []orgs.Organisation{}

	err = client.Get("user/orgs", &response)
	if err != nil {
		fmt.Println(err)
		return ErrMsg{Err: err}
	}

	return OrgListMsg{Organisations: response}
}
