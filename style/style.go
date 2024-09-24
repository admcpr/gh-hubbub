package style

import (
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/lipgloss"
)

var (
	Pink         = lipgloss.Color("#f72585")
	Purple       = lipgloss.Color("#7209b7")
	PurpleDarker = lipgloss.Color("#3a0ca3")
	Blue         = lipgloss.Color("#4361ee")
	BlueLighter  = lipgloss.Color("#4cc9f0")
	White        = lipgloss.Color("#dddddd")

	App = lipgloss.NewStyle().Padding(0, 0).Foreground(White).BorderForeground(Blue)

	Tab = lipgloss.NewStyle().BorderForeground(BlueLighter).
		Border(lipgloss.NormalBorder(), true, false, false, false).
		Align(lipgloss.Center)
	TabActive = Tab.Foreground(Pink).BorderForeground(Pink)

	Title = lipgloss.NewStyle().
		Foreground(Blue).
		BorderForeground(BlueLighter).
		Border(lipgloss.NormalBorder(), false, false, true, false).
		Padding(1, 1, 0, 1)

	Error = lipgloss.NewStyle().Foreground(PurpleDarker)

	DefaultDelegate = BuildDefaultDelegate()
)

func BuildDefaultDelegate() list.DefaultDelegate {
	defaultDelegate := list.NewDefaultDelegate()
	defaultDelegate.Styles.SelectedTitle.Foreground(Pink)
	defaultDelegate.Styles.SelectedTitle.BorderForeground(Pink)
	defaultDelegate.Styles.SelectedDesc.Foreground(Purple)
	defaultDelegate.Styles.SelectedDesc.BorderForeground(Pink)

	return defaultDelegate
}
