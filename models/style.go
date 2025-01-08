package models

import (
	"image/color"

	"github.com/charmbracelet/bubbles/v2/list"
	"github.com/charmbracelet/lipgloss/v2"
)

type colors struct {
	name         string
	black        color.Color
	red          color.Color
	green        color.Color
	yellow       color.Color
	blue         color.Color
	purple       color.Color
	cyan         color.Color
	white        color.Color
	brightBlack  color.Color
	brightRed    color.Color
	brightGreen  color.Color
	brightYellow color.Color
	brightBlue   color.Color
	brightPurple color.Color
	brightCyan   color.Color
	brightWhite  color.Color
	background   color.Color
	foreground   color.Color
}

func NewColors(darkmode bool) colors {
	colors := colors{
		name:         "PencilDark",
		black:        lipgloss.Color("#212121"),
		red:          lipgloss.Color("#c30771"),
		green:        lipgloss.Color("#10a778"),
		yellow:       lipgloss.Color("#a89c14"),
		blue:         lipgloss.Color("#008ec4"),
		purple:       lipgloss.Color("#523c79"),
		cyan:         lipgloss.Color("#20a5ba"),
		white:        lipgloss.Color("#d9d9d9"),
		brightBlack:  lipgloss.Color("#424242"),
		brightRed:    lipgloss.Color("#fb007a"),
		brightGreen:  lipgloss.Color("#5fd7af"),
		brightYellow: lipgloss.Color("#f3e430"),
		brightBlue:   lipgloss.Color("#20bbfc"),
		brightPurple: lipgloss.Color("#6855de"),
		brightCyan:   lipgloss.Color("#4fb8cc"),
		brightWhite:  lipgloss.Color("#f1f1f1"),
		background:   lipgloss.Color("#212121"),
		foreground:   lipgloss.Color("#f1f1f1"),
	}
	if darkmode {
		colors.name = "3024 Night"
		colors.background = lipgloss.Color("#090300")
		colors.foreground = lipgloss.Color("#a5a2a2")
	}
	return colors
}

var (
	AppColors = NewColors(true)

	appStyle = lipgloss.NewStyle().Padding(0, 0).
			Foreground(AppColors.foreground).
			BorderForeground(AppColors.blue).
			Border(lipgloss.RoundedBorder(), false)

	tabStyle = appStyle.Border(lipgloss.NormalBorder(), true, false, false, false).
			Align(lipgloss.Center)

	activeTabStyle = tabStyle.BorderForeground(AppColors.brightBlue).
			Foreground(AppColors.cyan)

	titleStyle = appStyle.Foreground(AppColors.blue).
			BorderForeground(AppColors.brightBlue).
			Border(lipgloss.NormalBorder(), false, false, true, true).
			Padding(0, 1, 0, 1)

	errorStyle = lipgloss.NewStyle().Foreground(AppColors.red)

	promptStyle = appStyle.Width(7).
			Align(lipgloss.Right).
			PaddingRight(1).
			MarginTop(1)

	textStyle = appStyle.Foreground(AppColors.foreground).
			PaddingLeft(1)

	cursorStyle = appStyle

	buttonStyle = appStyle.
			BorderForeground(AppColors.purple).
			Padding(0, 3).
			Margin(2)

	activeButtonStyle = buttonStyle.
				Foreground(AppColors.foreground).
				Background(AppColors.cyan).
				Underline(true)

	itemStyle         = lipgloss.NewStyle().PaddingLeft(2)
	selectedItemStyle = itemStyle.
				PaddingLeft(1).
				Foreground(AppColors.cyan).
				BorderForeground(AppColors.cyan).
				Border(lipgloss.NormalBorder(), false, false, false, true)

	modalTitleStyle = titleStyle.
			Align(lipgloss.Center).
			Foreground(AppColors.blue).
			BorderForeground(AppColors.brightGreen).
			Border(lipgloss.DoubleBorder(), false, false, true, false).
			Width(60)

	modalStyle = appStyle.
			BorderForeground(AppColors.green).
			Padding(0)

	DefaultDelegate = BuildDefaultDelegate()
)

func BuildDefaultDelegate() list.DefaultDelegate {
	d := list.NewDefaultDelegate()
	d.Styles.SelectedTitle = d.Styles.SelectedTitle.
		Foreground(AppColors.cyan).
		BorderForeground(AppColors.cyan)
	d.Styles.SelectedDesc = d.Styles.SelectedTitle

	return d
}
