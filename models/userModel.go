package models

import (
	"fmt"

	"gh-hubbub/structs"
	"gh-hubbub/style"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/cli/go-gh/v2/pkg/api"
)

type ErrMsg struct{ Err error }
type OrgListMsg struct{ Organisations []structs.Organisation }

type UserModel struct {
	organisations  []structs.Organisation
	User           structs.User
	SelectedOrgUrl string
	list           list.Model
}

func NewUserModel(user structs.User, width, height int) UserModel {
	userList := list.New([]list.Item{}, list.NewDefaultDelegate(), width, height)

	userList.Title = "User: " + user.Name
	userList.SetStatusBarItemName("Organisation", "Organisations")
	userList.Styles.Title = style.Title
	userList.SetShowTitle(true)

	return UserModel{User: user, list: userList}
}

func (m *UserModel) SetWidth(width int) {
	if len(m.list.Items()) > 0 {
		m.list.SetWidth(width)
	}
}

func (m *UserModel) SetHeight(height int) {
	if len(m.list.Items()) > 0 {
		m.list.SetHeight(height)
	}
}

func (m UserModel) Init() tea.Cmd {
	return getOrganisations
}

func (m UserModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case OrgListMsg:
		m.organisations = msg.Organisations
		items := make([]list.Item, len(m.organisations))
		for i, org := range m.organisations {
			items[i] = structs.NewListItem(org.Login, org.Url)

		}

		cmd = m.list.SetItems(items)

		return m, cmd
	}

	m.list, cmd = m.list.Update(msg)

	return m, cmd
}

func (m UserModel) View() string {
	return style.App.Render(m.list.View())
}

func (m UserModel) SelectedOrg() structs.Organisation {
	return m.organisations[m.list.Index()]
}

func getOrganisations() tea.Msg {
	client, err := api.DefaultRESTClient()
	if err != nil {
		return AuthenticationErrorMsg{Err: err}
	}
	response := []structs.Organisation{}

	err = client.Get("user/orgs", &response)
	if err != nil {
		fmt.Println(err)
		return ErrMsg{Err: err}
	}

	return OrgListMsg{Organisations: response}
}
