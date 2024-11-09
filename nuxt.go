package main

import (
	"fmt"

	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// Create a new Nuxt.js project
func CreateNuxtProject(name string) string {
	ClearScreen()

	s := spinner.New()
	s.Spinner = spinner.Dot
	s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))

	p := tea.NewProgram(initialSpinnerModel(s, name, "Creating Nuxt project"))
	p.Run()

	message := fmt.Sprintf("Created Nuxt project: %s\n", name)
	fmt.Print(message)
	return message
}

// Create a new Nuxt.js with MySQL project
func CreateNuxtWithMySQL(name string) string {
	ClearScreen()

	s := spinner.New()
	s.Spinner = spinner.Dot
	s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))

	p := tea.NewProgram(initialSpinnerModel(s, name, "Creating Nuxt with MySQL project"))
	p.Run()

	message := fmt.Sprintf("Created Nuxt with MySQL project: %s\n", name)
	fmt.Print(message)
	return message
}

// Create a new Nuxt.js with Pocketbase project
func CreateNuxtWithPocketbase(name string) string {
	ClearScreen()

	s := spinner.New()
	s.Spinner = spinner.Dot
	s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))

	p := tea.NewProgram(initialSpinnerModel(s, name, "Creating Nuxt with Pocketbase project"))
	p.Run()

	message := fmt.Sprintf("Created Nuxt with Pocketbase project: %s\n", name)
	fmt.Print(message)
	return message
}
