package main

import (
	"flag"
	"fmt"
	"gowizard/internal/menu"
	"gowizard/internal/utils"

	"github.com/charmbracelet/lipgloss"
)

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

func main() {

	// Define command-line flags
	initFlag := flag.Bool("init", false, "Initialize a new project")
	wizardFlag := flag.Bool("wizard", false, "Run the wizard interface")
	flag.Parse()

	var menuBanner string
	menuBanner += printBanner() + "\n" + utils.AccessPath() + "\n\n"
	fmt.Printf(menuBanner)
	// Entry point for the application

	if *initFlag {
		// Run the init menu
		menu.InitMenu()
	} else if *wizardFlag {
		menu.InitWizard()
	} else {
		// Run the regular menu
		menu.InitDefault()
	}
}
