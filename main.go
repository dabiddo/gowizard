package main

// A simple example that shows how to retrieve a value from a Bubble Tea
// program after the Bubble Tea has exited.

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var choices = []string{
	"Create Laravel Project with Composer",
	"Create Laravel Project with Laravel CLI",
	"Create Laravel StarterKit Project",
	"Create Laravel Project with MySQL",
	"Create Laravel Project with PostgreSQL",
	"Create Nuxt Project",
	"Create Nuxt Project with MySQL",
	"Create Nuxt Project with Pocketbase",
	"Create Astro Project",
	"Create Astro Blog Project",
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

type model struct {
	cursor int
	choice string
	name   string
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q", "esc":
			return m, tea.Quit

		case "enter":
			// Send the choice on the channel and exit.
			m.choice = choices[m.cursor]
			return m, tea.Quit

		case "down", "j":
			m.cursor++
			if m.cursor >= len(choices) {
				m.cursor = 0
			}

		case "up", "k":
			m.cursor--
			if m.cursor < 0 {
				m.cursor = len(choices) - 1
			}
		}
	}

	return m, nil
}

func (m model) View() string {

	s := strings.Builder{}
	s.WriteString("What kind of Project would you like to create?\n\n")

	for i := 0; i < len(choices); i++ {
		if m.cursor == i {
			s.WriteString("(•) ")
		} else {
			s.WriteString("( ) ")
		}
		s.WriteString(choices[i])
		s.WriteString("\n")
	}
	s.WriteString("\n(press q to quit)\n")

	return s.String()
}

func main() {
	p := tea.NewProgram(model{})
	var menu string
	menu += printBanner() + "\n" + AccessPath() + "\n\n"
	fmt.Printf(menu)
	// Run returns the model as a tea.Model.
	m, err := p.Run()
	if err != nil {
		fmt.Println("Oh no:", err)
		os.Exit(1)
	}

	// Assert the final tea.Model to our local model and print the choice.
	if m, ok := m.(model); ok && m.choice != "" {
		ClearScreen()
		fmt.Printf("\n---\nYou chose %s!\n", m.choice)
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter Project Name: ")
		text, _ := reader.ReadString('\n')
		fmt.Println(text)
		projectChoose(m)
	}
}

func projectChoose(m model) {
	switch m.cursor {
	case 0:
		CreateLaravelComposerProject(m.name)
	case 1:
		CreateLaravelCLIProject(m.name)
	case 2:
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter starter Project (larave/react): ")
		starter, _ := reader.ReadString('\n')
		CreateLaravelStarterProject(m.name, starter)
	case 3:
		CreateLaravelWithMySQL(m.name)
	case 4:
		CreateLaravelWithPostgreSQL(m.name)
	case 5:
		CreateNuxtProject(m.name)
	case 6:
		CreateNuxtWithMySQL(m.name)
	case 7:
		CreateNuxtWithPocketbase(m.name)
	case 8:
		CreateAstroProject(m.name)
	case 9:
		CreateAstroBlogProject(m.name)
	default:
		fmt.Printf("Invalid option")
	}
}
