package projects

import (
	"fmt"
	"gowizard/internal/utils"
	"os"
	"os/exec"
	"os/user"

	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// Create a new Nuxt.js project
func CreateNuxtProject(name string) string {
	utils.ClearScreen()

	// Build the Docker command
	cmd := exec.Command("docker", "run", "--rm",
		"-v", fmt.Sprintf("%s:/app", utils.GetCurrentPath()),
		"-w", "/app",
		"-it",
		"node:20.11.1-alpine",
		"sh", "-c",
		fmt.Sprintf("pnpm dlx nuxt init %s --yes --package-manager pnpm --git-init && chown -R $(id -u):$(id -g) %s", name, name))

	// Set up pipes for real-time output
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	fmt.Println("Creating NuxtJs project...")

	// Run the command
	if err := cmd.Run(); err != nil {
		fmt.Printf("\nError creating NuxtJs project: %v\n", err)
		return "Failed to create NuxtJs project"
	}

	// Change ownership of the project directory
	currentUser, err := user.Current()
	if err != nil {
		fmt.Printf("\nError getting current user: %v\n", err)
		return "Failed to get current user"
	}

	utils.ChangeOwnership(utils.GetCurrentPath(), currentUser.Username, name)

	// Clean up temporary files
	cleanCmd := exec.Command("rm", "-rf", ".pnpm-store")
	if err := cleanCmd.Run(); err != nil {
		fmt.Printf("\nWarning: Failed to clean temporary files: %v\n", err)
	}

	// Create dev container files
	utils.CreateDevContainer(name, "_nuxt.stub")
	utils.CreateDockerfile(name, "_nuxt.stub")

	message := fmt.Sprintf("NuxtJs project '%s' created successfully!\n", name)
	fmt.Print(message)
	return message
}

// Create a new Nuxt.js with MySQL project
func CreateNuxtWithMySQL(name string) string {
	utils.ClearScreen()

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
	utils.ClearScreen()

	s := spinner.New()
	s.Spinner = spinner.Dot
	s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))

	p := tea.NewProgram(initialSpinnerModel(s, name, "Creating Nuxt with Pocketbase project"))
	p.Run()

	message := fmt.Sprintf("Created Nuxt with Pocketbase project: %s\n", name)
	fmt.Print(message)
	return message
}
