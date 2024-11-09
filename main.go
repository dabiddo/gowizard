package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// Define menu options
var options = []string{
	"Create Laravel Project with Composer",
	"Create Laravel Project with Laravel CLI",
	"Create Laravel Project with MySQL",
	"Create Laravel Project with PostgreSQL",
	"Create Nuxt Project",
	"Create Nuxt Project with MySQL",
	"Create Nuxt Project with Pocketbase",
	"Quit",
}

// Styling
var (
	titleStyle  = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#00BFFF")).Padding(1, 0, 1, 2)
	optionStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#FFA500")).Padding(0, 0, 0, 2)
	activeStyle = optionStyle.Copy().Bold(true).Background(lipgloss.Color("#333333"))
	bannerStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#FFD700")).Bold(true)
)

// Banner function to display
func printBanner() string {
	banner := `

	██████╗  ██████╗ ██╗    ██╗██╗███████╗ █████╗ ██████╗ ██████╗ 
	██╔════╝ ██╔═══██╗██║    ██║██║╚══███╔╝██╔══██╗██╔══██╗██╔══██╗
	██║  ███╗██║   ██║██║ █╗ ██║██║  ███╔╝ ███████║██████╔╝██║  ██║
	██║   ██║██║   ██║██║███╗██║██║ ███╔╝  ██╔══██║██╔══██╗██║  ██║
	╚██████╔╝╚██████╔╝╚███╔███╔╝██║███████╗██║  ██║██║  ██║██████╔╝
	 ╚═════╝  ╚═════╝  ╚══╝╚══╝ ╚═╝╚══════╝╚═╝  ╚═╝╚═╝  ╚═╝╚═════╝ 
																   
	
`
	return bannerStyle.Render(banner)
}

// choice model for the app
type choice struct {
	cursor      int
	projectName string
	askingName  bool
}

// Define custom quit message
type quitMsg struct{}

// Add this new type for our completion message
type completionMsg struct {
	message string
}

// Initialize the model
func initialModel() choice {
	return choice{}
}

// Init function
func (m choice) Init() tea.Cmd {
	return nil
}

// Update function
func (m choice) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "up":
			if !m.askingName && m.cursor > 0 {
				m.cursor--
			}
		case "down":
			if !m.askingName && m.cursor < len(options)-1 {
				m.cursor++
			}
		case "enter":
			if m.askingName {
				switch m.cursor {
				case 0:
					//CreateLaravelProject(m.projectName)
					CreateLaravelComposerProject(m.projectName)
				case 1:
					//CreateLaravelWithMySQL(m.projectName)
					CreateLaravelCLIProject(m.projectName)
				case 2:
					CreateLaravelWithMySQL(m.projectName)
				case 3:
					CreateLaravelWithPostgreSQL(m.projectName)
				case 4:
					CreateNuxtProject(m.projectName)
				case 5:
					CreateNuxtWithMySQL(m.projectName)
				case 6:
					CreateNuxtWithPocketbase(m.projectName)
				}
				return m, tea.Quit
			}

			if m.cursor == len(options)-1 {
				return m, tea.Quit
			}

			// Initialize project name input
			m.askingName = true
			m.projectName = ""
			return m, nil
		default:
			// Handle project name input
			if m.askingName {
				switch msg.String() {
				case "backspace":
					if len(m.projectName) > 0 {
						m.projectName = m.projectName[:len(m.projectName)-1]
					}
				default:
					// Append character to projectName if it's a single character
					if len(msg.String()) == 1 {
						m.projectName += msg.String()
					}
				}
			}
		}
	case completionMsg:
		return m, tea.Quit
	}

	return m, nil
}

// View function
func (m choice) View() string {
	if m.askingName {
		return "Enter project name: " + m.projectName + "\n\nPress 'enter' to confirm"
	}

	var menu string
	menu += printBanner() + "\n" + AccessPath() + "\n\n"

	for i, option := range options {
		cursor := "  "
		if m.cursor == i {
			cursor = ">"
		}
		if i == m.cursor {
			menu += activeStyle.Render(cursor+" "+option) + "\n"
		} else {
			menu += optionStyle.Render(cursor+" "+option) + "\n"
		}
	}

	menu += "\nPress 'enter' to select, 'q' to quit."
	return menu
}

func main() {
	// Start the Bubble Tea program
	if err := tea.NewProgram(initialModel()).Start(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
