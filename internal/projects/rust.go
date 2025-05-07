package projects

import (
	"fmt"
	"gowizard/internal/utils"
	"os"
	"os/exec"
	"os/user"
)

func CreateRustProject(name string) string {
	utils.ClearScreen()

	// Build the Docker command
	cmd := exec.Command("docker", "run", "--rm",
		"-v", fmt.Sprintf("%s:/app", utils.GetCurrentPath()),
		"-w", "/app",
		"-it",
		"rust:slim-bookworm",
		"sh", "-c",
		fmt.Sprintf("cargo new %s && chown -R $(id -u):$(id -g) %s", name, name))

	// Set up pipes for real-time output
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	fmt.Println("Creating Rust project...")

	// Run the command
	if err := cmd.Run(); err != nil {
		fmt.Printf("\nError creating Rust project: %v\n", err)
		return "Failed to create Rust project"
	}

	// Change ownership of the project directory
	currentUser, err := user.Current()
	if err != nil {
		fmt.Printf("\nError getting current user: %v\n", err)
		return "Failed to get current user"
	}

	utils.ChangeOwnership(utils.GetCurrentPath(), currentUser.Username, name)

	// Create dev container files
	utils.CreateDevContainer(name, "_rust.stub")
	utils.CreateDockerfile(name, "_rust.stub")

	message := fmt.Sprintf("Rust project '%s' created successfully!\n", name)
	fmt.Print(message)
	return message
}
