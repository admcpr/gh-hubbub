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
		name:         "3024 Day",
		black:        lipgloss.Color("#090300"),
		red:          lipgloss.Color("#db2d20"),
		green:        lipgloss.Color("#01a252"),
		yellow:       lipgloss.Color("#fded02"),
		blue:         lipgloss.Color("#01a0e4"),
		purple:       lipgloss.Color("#a16a94"),
		cyan:         lipgloss.Color("#b5e4f4"),
		white:        lipgloss.Color("#a5a2a2"),
		brightBlack:  lipgloss.Color("#5c5855"),
		brightRed:    lipgloss.Color("#e8bbd0"),
		brightGreen:  lipgloss.Color("#3a3432"),
		brightYellow: lipgloss.Color("#4a4543"),
		brightBlue:   lipgloss.Color("#807d7c"),
		brightPurple: lipgloss.Color("#d6d5d4"),
		brightCyan:   lipgloss.Color("#cdab53"),
		brightWhite:  lipgloss.Color("#f7f7f7"),
		background:   lipgloss.Color("#f7f7f7"),
		foreground:   lipgloss.Color("#4a4543"),
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
			Foreground(AppColors.white).
			BorderForeground(AppColors.blue)

	tabStyle = lipgloss.NewStyle().BorderForeground(AppColors.brightBlue).
			Border(lipgloss.NormalBorder(), true, false, false, false).
			Align(lipgloss.Center)

	activeTabStyle = tabStyle.Foreground(AppColors.cyan).
			BorderForeground(AppColors.cyan)

	titleStyle = lipgloss.NewStyle().
			Foreground(AppColors.blue).
			BorderForeground(AppColors.brightBlue).
			Border(lipgloss.NormalBorder(), false, false, true, false).
			Padding(1, 1, 0, 1)

	errorStyle = lipgloss.NewStyle().Foreground(AppColors.purple)

	promptStyle = lipgloss.NewStyle().
			Foreground(AppColors.foreground).
			Width(7).
			Align(lipgloss.Right).
			PaddingRight(1).
			MarginTop(1)

	textStyle = lipgloss.NewStyle().
			Foreground(AppColors.foreground).
			PaddingLeft(1)

	cursorStyle = lipgloss.NewStyle().
			Foreground(AppColors.foreground)

	buttonStyle = lipgloss.NewStyle().
			Foreground(AppColors.foreground).
			BorderForeground(AppColors.purple).
			Padding(0, 3).
			Margin(2)

	activeButtonStyle = buttonStyle.
				Foreground(AppColors.foreground).
				Background(AppColors.cyan).
				Underline(true)

	itemStyle         = lipgloss.NewStyle().PaddingLeft(4)
	selectedItemStyle = lipgloss.NewStyle().PaddingLeft(2).
				Foreground(AppColors.cyan)

	headerStyle = lipgloss.NewStyle().
			Foreground(AppColors.foreground).
			Background(AppColors.background)

	modalTitleStyle = lipgloss.NewStyle().
			Align(lipgloss.Center).
			Foreground(AppColors.blue).
			BorderForeground(AppColors.brightBlue).
			Border(lipgloss.DoubleBorder(), false, false, true, false).
			Width(60)

	modalStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(AppColors.blue).
			Padding(0)

	DefaultDelegate = BuildDefaultDelegate()
)

func BuildDefaultDelegate() list.DefaultDelegate {
	defaultDelegate := list.NewDefaultDelegate()
	defaultDelegate.Styles.SelectedTitle.Foreground(AppColors.cyan)
	defaultDelegate.Styles.SelectedTitle.BorderForeground(AppColors.cyan)
	defaultDelegate.Styles.SelectedDesc.Foreground(AppColors.purple)
	defaultDelegate.Styles.SelectedDesc.BorderForeground(AppColors.cyan)

	return defaultDelegate
}
