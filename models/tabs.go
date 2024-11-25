package models

import (
	"gh-hubbub/style"

	"github.com/charmbracelet/lipgloss/v2"
)

func RenderTabs(tabs []string, width, activeTab int) string {
	tabWidth := (width - 4) / len(tabs)

	var renderedTabs []string

	for i, t := range tabs {
		tabStyle := style.Tab
		isFirst, isLast, isActive := i == 0, i == len(tabs)-1, i == activeTab

		if isActive {
			tabStyle = style.TabActive
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
