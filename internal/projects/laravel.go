package projects

import (
	"fmt"
	"gowizard/internal/utils"
	"os"
	"os/exec"
	"os/user"
	"time"

	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// Laravel project functions
func CreateLaravelProject(name string) string {
	utils.ClearScreen()

	// Initialize spinner
	s := spinner.New()
	s.Spinner = spinner.Dot
	s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))

	p := tea.NewProgram(initialSpinnerModel(s, name, "Creating Laravel project"))
	p.Run()

	// After spinner completes
	message := fmt.Sprintf("Created Laravel project: %s\n", name)
	fmt.Print(message)
	return message
}

func CreateLaravelCLIProject(name string) string {
	utils.ClearScreen()

	// Get current working directory
	currentPath := utils.GetCurrentPath()

	// Build the Docker command with the new format
	cmd := exec.Command("docker", "run", "--rm", "-it",
		"-v", fmt.Sprintf("%s:/app", currentPath),
		"dabiddo/larabox",
		"laravel", "new", name)

	// Set up pipes for real-time output
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin // Important for interactive prompts

	fmt.Println("\nCreating Laravel project with CLI...")

	// Run the command
	if err := cmd.Run(); err != nil {
		fmt.Printf("\nError creating Laravel project: %v\n", err)
		return "Failed to create Laravel project"
	}

	// Change ownership of the project directory
	currentUser, err := user.Current()
	if err != nil {
		fmt.Printf("\nError getting current user: %v\n", err)
		return "Failed to get current user"
	}

	utils.ClearScreen()
	utils.ChangeOwnership(utils.GetCurrentPath(), currentUser.Username, name)
	utils.CreateDevContainer(name, "_laravel.stub")
	utils.CreateDockerfile(name, "_laravel.stub")
	utils.CreateDockerCompose(name, "_laravel.stub")
	utils.CopyXdebug(name, "_xdebug.stub")
	utils.CopyXdebugLauch(name, "_xdebug.stub")

	message := fmt.Sprintf("Laravel project '%s' created successfully!\n", name)
	fmt.Print(message)
	return message
}

func CreateLaravelStarterProject(name string, using string) string {
	utils.ClearScreen()

	// Get current working directory
	currentPath := utils.GetCurrentPath()

	// Prepare command arguments
	dockerArgs := []string{"run", "--rm", "-it",
		"-v", fmt.Sprintf("%s:/app", currentPath),
		"dabiddo/larabox",
		"laravel", "new", name}

	// Add --using parameter if provided
	if using != "" {
		dockerArgs = append(dockerArgs, fmt.Sprintf("--using=%s", using))
	}

	// Build the Docker command with the arguments
	cmd := exec.Command("docker", dockerArgs...)

	// Set up pipes for real-time output
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin // Important for interactive prompts

	fmt.Println("\nCreating Laravel project with Stertkit...")

	// Run the command
	if err := cmd.Run(); err != nil {
		fmt.Printf("\nError creating Laravel project: %v\n", err)
		return "Failed to create Laravel project"
	}

	// Change ownership of the project directory
	currentUser, err := user.Current()
	if err != nil {
		fmt.Printf("\nError getting current user: %v\n", err)
		return "Failed to get current user"
	}

	utils.ClearScreen()
	utils.ChangeOwnership(utils.GetCurrentPath(), currentUser.Username, name)
	utils.CreateDevContainer(name, "_laravel.stub")
	utils.CreateDockerfile(name, "_laravel.stub")
	utils.CreateDockerCompose(name, "_laravel.stub")
	utils.CopyXdebug(name, "_xdebug.stub")
	utils.CopyXdebugLauch(name, "_xdebug.stub")

	message := fmt.Sprintf("Laravel project '%s' created successfully!\n", name)
	fmt.Print(message)
	return message
}

func CreateLaravelWithMySQL(name string) string {
	utils.ClearScreen()

	s := spinner.New()
	s.Spinner = spinner.Dot
	s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))

	p := tea.NewProgram(initialSpinnerModel(s, name, "Creating Laravel with MySQL project"))

	// Create a channel to signal completion
	done := make(chan bool)

	// Run Docker command in a goroutine
	go func() {
		cmd := exec.Command("docker", "run", "--rm", "-v", fmt.Sprintf("%s:/app", utils.GetCurrentPath()), "composer", "create-project", "--prefer-dist", "laravel/laravel", name)

		if err := cmd.Run(); err != nil {
			fmt.Printf("\nError creating Laravel project: %v\n", err)
			done <- false
			return
		}

		// Change ownership of the project directory
		currentUser, err := user.Current()
		if err != nil {
			fmt.Printf("\nError getting current user: %v\n", err)
			done <- false
			return
		}
		utils.ClearScreen()
		utils.ChangeOwnership(utils.GetCurrentPath(), currentUser.Username, name)
		utils.CreateDevContainer(name, "_laravel.stub")
		utils.CreateDockerfile(name, "_laravel.stub")
		utils.CreateDockerCompose(name, "_laravel_mysql8.stub")
		utils.CopyXdebug(name, "_xdebug.stub")
		utils.CopyXdebugLauch(name, "_xdebug.stub")
		done <- true
	}()

	// Run the spinner program
	p.Run()

	// Wait for the background process to complete
	success := <-done

	if success {
		message := fmt.Sprintf("Laravel project '%s' created successfully!\n", name)
		fmt.Print(message)
		return message
	}

	return "Failed to create Laravel project"
}

func CreateLaravelWithPostgreSQL(name string) string {
	utils.ClearScreen()

	s := spinner.New()
	s.Spinner = spinner.Dot
	s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))

	p := tea.NewProgram(initialSpinnerModel(s, name, "Creating Laravel with MySQL project"))

	// Create a channel to signal completion
	done := make(chan bool)

	// Run Docker command in a goroutine
	go func() {
		cmd := exec.Command("docker", "run", "--rm", "-v", fmt.Sprintf("%s:/app", utils.GetCurrentPath()), "composer", "create-project", "--prefer-dist", "laravel/laravel", name)

		if err := cmd.Run(); err != nil {
			fmt.Printf("\nError creating Laravel project: %v\n", err)
			done <- false
			return
		}

		// Change ownership of the project directory
		currentUser, err := user.Current()
		if err != nil {
			fmt.Printf("\nError getting current user: %v\n", err)
			done <- false
			return
		}
		utils.ClearScreen()
		utils.ChangeOwnership(utils.GetCurrentPath(), currentUser.Username, name)
		utils.CreateDevContainer(name, "_laravel.stub")
		utils.CreateDockerfile(name, "_laravel.stub")
		utils.CreateDockerCompose(name, "_laravel_pgsql.stub")
		utils.CopyXdebug(name, "_xdebug.stub")
		utils.CopyXdebugLauch(name, "_xdebug.stub")
		done <- true
	}()

	// Run the spinner program
	p.Run()

	// Wait for the background process to complete
	success := <-done

	if success {
		message := fmt.Sprintf("Laravel project '%s' created successfully!\n", name)
		fmt.Print(message)
		return message
	}

	return "Failed to create Laravel project"
}

func CreateLaravelComposerProject(name string) string {
	utils.ClearScreen()

	// Initialize spinner
	s := spinner.New()
	s.Spinner = spinner.Dot
	s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))

	p := tea.NewProgram(initialSpinnerModel(s, name, "Creating Laravel project with Composer"))

	// Create a channel to signal completion
	done := make(chan bool)

	// Run Docker command in a goroutine
	go func() {
		cmd := exec.Command("docker", "run", "--rm", "-v", fmt.Sprintf("%s:/app", utils.GetCurrentPath()), "composer", "create-project", "--prefer-dist", "laravel/laravel", name)

		if err := cmd.Run(); err != nil {
			fmt.Printf("\nError creating Laravel project: %v\n", err)
			done <- false
			return
		}

		// Change ownership of the project directory
		currentUser, err := user.Current()
		if err != nil {
			fmt.Printf("\nError getting current user: %v\n", err)
			done <- false
			return
		}
		utils.ClearScreen()
		utils.ChangeOwnership(utils.GetCurrentPath(), currentUser.Username, name)
		utils.CreateDevContainer(name, "_laravel.stub")
		utils.CreateDockerfile(name, "_laravel.stub")
		utils.CreateDockerCompose(name, "_laravel.stub")
		utils.CopyXdebug(name, "_xdebug.stub")
		utils.CopyXdebugLauch(name, "_xdebug.stub")
		done <- true
	}()

	// Run the spinner program
	p.Run()

	// Wait for the background process to complete
	success := <-done

	if success {
		message := fmt.Sprintf("Laravel project '%s' created successfully!\n", name)
		fmt.Print(message)
		return message
	}

	return "Failed to create Laravel project"
}

// Spinner model
type spinnerModel struct {
	spinner spinner.Model
	name    string
	message string
	done    bool
}

func initialSpinnerModel(s spinner.Model, name string, message string) spinnerModel {
	return spinnerModel{
		spinner: s,
		name:    name,
		message: message,
		done:    false,
	}
}

func (m spinnerModel) Init() tea.Cmd {
	return tea.Batch(
		m.spinner.Tick,
		func() tea.Msg {
			time.Sleep(time.Second * 4)
			return "done"
		},
	)
}

func (m spinnerModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "ctrl+c" {
			return m, tea.Quit
		}
	case string:
		if msg == "done" {
			m.done = true
			return m, tea.Quit
		}
	default:
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd
	}
	return m, nil
}

func (m spinnerModel) View() string {
	if m.done {
		return ""
	}
	return fmt.Sprintf("\n %s %s: %s\n", m.spinner.View(), m.message, m.name)
}
