package projects

import (
	"fmt"
	"gowizard/internal/utils"
	"os"
	"os/exec"
	"os/user"
)

func CreateNestJSProject(name string) string {
	utils.ClearScreen()

	// Build the Docker command
	cmd := exec.Command("docker", "run", "--rm",
		"-v", fmt.Sprintf("%s:/app", utils.GetCurrentPath()),
		"-w", "/app",
		"-it",
		"node:lts-alpine",
		"sh", "-c",
		fmt.Sprintf("nest new %s --strict --skip-git --package-manager=pnpm && chown -R $(id -u):$(id -g) %s",
			name, name))

	// Set up pipes for real-time output
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	fmt.Println("Creating NestJs project...")

	// Run the command
	if err := cmd.Run(); err != nil {
		fmt.Printf("\nError creating NestJs project: %v\n", err)
		return "Failed to create NestJs project"
	}

	// Change ownership of the project directory
	currentUser, err := user.Current()
	if err != nil {
		fmt.Printf("\nError getting current user: %v\n", err)
		return "Failed to get current user"
	}

	utils.ChangeOwnership(utils.GetCurrentPath(), currentUser.Username, name)

	cleanCmd := exec.Command("bash", "-c", "sudo rm -rf .pnpm-store/ 2>/dev/null")
	if err := cleanCmd.Run(); err != nil {
		fmt.Printf("\nWarning: Failed to clean temporary files: %v\n", err)
		// Continue execution even if cleanup fails
	}

	// Create dev container files
	utils.CreateDevContainer(name, "_astro.stub")
	utils.CreateDockerfile(name, "_nodejs.stub")

	// Add logic to create the NestJS project here
	message := fmt.Sprintf("NestJS project '%s' created successfully!\n", name)
	fmt.Print(message)
	return message
}
