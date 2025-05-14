package projects

import (
	"fmt"
	"gowizard/internal/utils"
	"os"
	"os/exec"
	"os/user"
)

func CreatePayloadCMSProject(name string) string {
	utils.ClearScreen()

	// Build the Docker command
	cmd := exec.Command("docker", "run", "--rm",
		"-v", fmt.Sprintf("%s:/app", utils.GetCurrentPath()),
		"-w", "/app",
		"-it",
		"larabox:latest",
		"sh", "-c",
		fmt.Sprintf("npx create-payload-app -n %s --use-pnpm && chown -R $(id -u):$(id -g) %s", name, name))

	// Set up pipes for real-time output
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	fmt.Println("Creating Payload CMS project...")

	// Run the command
	if err := cmd.Run(); err != nil {
		fmt.Printf("\nError creating Payload CMS project: %v\n", err)
		return "Failed to create Payload CMS project"
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

	message := fmt.Sprintf("Payload CMS project '%s' created successfully!\n", name)
	fmt.Print(message)
	return message
}
