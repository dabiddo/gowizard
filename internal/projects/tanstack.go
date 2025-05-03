package projects

import (
	"fmt"
	"gowizard/internal/utils"
	"os"
	"os/exec"
	"os/user"
)

func CreateTanstackProject(name string) string {

	utils.ClearScreen()

	// Build the Docker command
	cmd := exec.Command("docker", "run", "--rm",
		"-v", fmt.Sprintf("%s:/app", utils.GetCurrentPath()),
		"-w", "/app",
		"-it",
		"node:lts-alpine",
		"sh", "-c",
		fmt.Sprintf("apk add --no-cache git && npx create-better-t-stack@latest %s && chown -R $(id -u):$(id -g) %s",
			name, name))

	// Set up pipes for real-time output
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	fmt.Println("Creating Tanstack project...")

	// Run the command
	if err := cmd.Run(); err != nil {
		fmt.Printf("\nError creating Tanstack project: %v\n", err)
		return "Failed to create Tanstack project"
	}

	// Change ownership of the project directory
	currentUser, err := user.Current()
	if err != nil {
		fmt.Printf("\nError getting current user: %v\n", err)
		return "Failed to get current user"
	}

	utils.ChangeOwnership(utils.GetCurrentPath(), currentUser.Username, name)

	// Create dev container files
	utils.CreateDevContainer(name, "_astro.stub")
	utils.CreateDockerfile(name, "_nodejs.stub")

	message := fmt.Sprintf("Tanstack project '%s' created successfully!\n", name)
	fmt.Print(message)
	return message
}
