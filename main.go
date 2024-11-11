package main

import (
	"fmt"
	"os"

	"gh-hubbub/models"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	mainModelV2 := models.NewMainModel()
	p := tea.NewProgram(mainModelV2, tea.WithAltScreen())
	// p := tea.NewProgram(models.NewFilterBooleanModel("Is this a boolean?", true))
	// p := tea.NewProgram(models.NewFilterNumberModel("What number would you like to choose?", 1, 1))
	// p := tea.NewProgram(models.NewFilterDateModel("What date would you like to choose?", time.Now(), time.Now()))

	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
