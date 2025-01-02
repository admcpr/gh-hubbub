package models

import (
	"github.com/charmbracelet/bubbles/v2/paginator"
	tea "github.com/charmbracelet/bubbletea/v2"
)

type TabSelectMessage struct{ Index int }

type RepoHeaderModel struct {
	titles    []string
	paginator paginator.Model
	width     int
	height    int
}

func NewRepoHeaderModel(width int, titles []string, index int) RepoHeaderModel {
	p := paginator.New()
	p.Type = paginator.Dots
	p.PerPage = 1
	p.SetTotalPages(len(titles))
	p.Page = index

	return RepoHeaderModel{
		titles:    titles,
		paginator: p,
		width:     width,
	}
}

func (m *RepoHeaderModel) SetDimensions(width, height int) {
	m.width = width
	m.height = height
}

func (m RepoHeaderModel) Init() (tea.Model, tea.Cmd) {
	return m, nil
}

func (m RepoHeaderModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m, nil
}

func (m RepoHeaderModel) View() string {
	heading := headerStyle.Render(m.titles[m.paginator.Page])
	return heading + "\n" + m.paginator.View()
}
