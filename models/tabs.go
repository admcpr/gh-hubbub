package models

import (
	"gh-hubbub/structs"
	"gh-hubbub/style"

	"github.com/charmbracelet/lipgloss"
)

func RenderTabs(tabSettings []structs.SettingsTab, width, activeTab int) string {
	tabs := []string{}
	for _, t := range tabSettings {
		tabs = append(tabs, t.Name)
	}

	tabWidth := (width - 2) / len(tabs)

	var renderedTabs []string

	for i, t := range tabs {
		tabStyle := lipgloss.NewStyle().BorderForeground(style.BlueLighter).
			Border(lipgloss.NormalBorder(), true, false, false, false).
			Align(lipgloss.Center)
		isFirst, isLast, isActive := i == 0, i == len(tabs)-1, i == activeTab

		if isActive {
			tabStyle = tabStyle.Foreground(style.Pink).BorderForeground(style.Pink)
		}

		if isLast {
			tabStyle = tabStyle.MarginRight(1).
				Width(tabWidth + (width % len(tabs)))
		} else {
			tabStyle = tabStyle.Width(tabWidth)
			if isFirst {
				tabStyle = tabStyle.MarginLeft(1)
			}
		}

		renderedTabs = append(renderedTabs, tabStyle.Render(t))
	}

	row := lipgloss.JoinHorizontal(lipgloss.Top, renderedTabs...)

	return row
}
