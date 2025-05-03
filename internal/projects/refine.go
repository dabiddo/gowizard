package projects

import (
	"fmt"
	"gowizard/internal/utils"
	"os"
	"os/exec"
	"os/user"
)

func CreateRefineProject(name string) string {
	utils.ClearScreen()

	// Build the Docker command
	cmd := exec.Command("docker", "run", "--rm",
		"-v", fmt.Sprintf("%s:/app", utils.GetCurrentPath()),
		"-w", "/app",
		"-it",
		"node:lts-alpine",
		"sh", "-c",
		fmt.Sprintf("apk add --no-cache git && npm install create-refine-app && npm create refine-app@latest %s && chown -R $(id -u):$(id -g) %s",
			name, name))

	// Set up pipes for real-time output
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	fmt.Println("Creating Refine.dev project...")

	// Run the command
	if err := cmd.Run(); err != nil {
		fmt.Printf("\nError creating Refine.dev project: %v\n", err)
		return "Failed to create Refine.dev project"
	}

	// Change ownership of the project directory
	currentUser, err := user.Current()
	if err != nil {
		fmt.Printf("\nError getting current user: %v\n", err)
		return "Failed to get current user"
	}

	utils.ChangeOwnership(utils.GetCurrentPath(), currentUser.Username, name)

	//Clean node_modules, package.json & package-lock.json created in root directory
	fmt.Println("Cleaning up temporary files...")

	// Use bash with -c to handle both files and directories with proper error handling
	cleanCmd := exec.Command("bash", "-c", "sudo rm -rf node_modules/ 2>/dev/null; rm -f package.json package-lock.json 2>/dev/null || true")
	if err := cleanCmd.Run(); err != nil {
		fmt.Printf("\nWarning: Failed to clean temporary files: %v\n", err)
		// Continue execution even if cleanup fails
	}

	// Create dev container files
	utils.CreateDevContainer(name, "_astro.stub")
	utils.CreateDockerfile(name, "_nodejs.stub")

	message := fmt.Sprintf("Refine.dev project '%s' created successfully!\n", name)
	fmt.Print(message)
	return message
}
