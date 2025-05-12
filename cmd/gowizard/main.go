package main

import (
	"bufio"
	"flag"
	"fmt"
	"gowizard/internal/projects"
	"gowizard/internal/utils"
	"os"
	"path/filepath"
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
	"Create Refine.dev Project",
	"Create Tanstack Project",
	"Create NestJs Project",
	"Create Payload CMS Project",
	"Create Rust Project",
	"Create HonoJs Project",
}

var initChoices = []string{
	"Initialize PHP Project",
	"Initialize PHP Project with FrankenPHP",
	"Initialize NodeJs Project",
	"Initialize Golang Project",
}

// Styling
var (
	titleStyle  = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#00BFFF")).Padding(1, 0, 1, 2)
	optionStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#FFA500")).Padding(0, 0, 0, 2)
	activeStyle = optionStyle.Copy().Bold(true).Background(lipgloss.Color("#3333"))
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

type initModel struct {
	cursor int
	choice string
	name   string
}

func (m initModel) Init() tea.Cmd {
	return nil
}

func (m initModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q", "esc":
			return m, tea.Quit

		case "enter":
			m.choice = initChoices[m.cursor]
			return m, tea.Quit

		case "down", "j":
			m.cursor++
			if m.cursor >= len(initChoices) {
				m.cursor = 0
			}

		case "up", "k":
			m.cursor--
			if m.cursor < 0 {
				m.cursor = len(initChoices) - 1
			}
		}
	}
	return m, nil
}

func (m initModel) View() string {
	s := strings.Builder{}
	s.WriteString("Initialize a new project:\n\n")

	for i := 0; i < len(initChoices); i++ {
		if m.cursor == i {
			s.WriteString("(•) ")
		} else {
			s.WriteString("( ) ")
		}
		s.WriteString(initChoices[i])
		s.WriteString("\n")
	}
	s.WriteString("\n(press q to quit)\n")

	return s.String()
}

func main() {
	// Define command-line flags
	initFlag := flag.Bool("init", false, "Initialize a new project")
	flag.Parse()

	var menu string
	menu += printBanner() + "\n" + utils.AccessPath() + "\n\n"
	fmt.Printf(menu)

	var finalModel tea.Model
	var err error

	if *initFlag {
		// Run the init menu
		p := tea.NewProgram(initModel{})
		finalModel, err = p.Run()
	} else {
		// Run the regular menu
		p := tea.NewProgram(model{})
		finalModel, err = p.Run()
	}

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	if *initFlag {
		if m, ok := finalModel.(initModel); ok && m.choice != "" {
			utils.ClearScreen()
			fmt.Printf("\n---\nYou chose %s!\n", m.choice)
			handleInitProject(m)
		}
	} else {
		if m, ok := finalModel.(model); ok && m.choice != "" {
			utils.ClearScreen()
			fmt.Printf("\n---\nYou chose %s!\n", m.choice)
			reader := bufio.NewReader(os.Stdin)
			fmt.Print("Enter Project Name: ")
			projectName, _ := reader.ReadString('\n')
			projectName = strings.TrimSpace(projectName)
			m.name = projectName
			projectChoose(m)
		}
	}
}

func projectChoose(m model) {
	switch m.cursor {
	case 0:
		projects.CreateLaravelComposerProject(m.name)
	case 1:
		projects.CreateLaravelCLIProject(m.name)
	case 2:
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter starter Project (larave/react): ")
		starter, _ := reader.ReadString('\n')
		projects.CreateLaravelStarterProject(m.name, starter)
	case 3:
		projects.CreateLaravelWithMySQL(m.name)
	case 4:
		projects.CreateLaravelWithPostgreSQL(m.name)
	case 5:
		projects.CreateNuxtProject(m.name)
	case 6:
		projects.CreateNuxtWithMySQL(m.name)
	case 7:
		projects.CreateNuxtWithPocketbase(m.name)
	case 8:
		projects.CreateAstroProject(m.name)
	case 9:
		projects.CreateAstroBlogProject(m.name)
	case 10:
		projects.CreateRefineProject(m.name)
	case 11:
		projects.CreateTanstackProject(m.name)
	case 12:
		projects.CreateNestJSProject(m.name)
	case 13:
		projects.CreatePayloadCMSProject(m.name)
	case 14:
		projects.CreateRustProject(m.name)
	case 15:
		projects.CreateHonoJsProject(m.name)
	default:
		fmt.Printf("Invalid option")
	}
}

func handleInitProject(m initModel) {
	currentDir := filepath.Base(utils.GetCurrentPath())
	switch m.cursor {
	case 0:
		// Initialize PHP project
		fmt.Printf("Initializing PHP project: %s\n", m.name)
		projects.PhpProject(currentDir)
	case 1:
		// Initialize PHP project with FrankenPHP
		fmt.Printf("Initializing PHP project with FrankenPHP: %s\n", m.name)
		projects.FrankenProject(currentDir)
	case 2:
		// Initialize NodeJs project
		fmt.Printf("Initializing NodeJs project: %s\n", m.name)
		projects.NodeProject(currentDir)
	case 3:
		// Initialize Golang project
		fmt.Printf("Initializing Golang project: %s\n", m.name)
		projects.GolangProject(currentDir)
	default:
		fmt.Printf("Invalid option")
	}
}
