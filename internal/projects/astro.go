package projects

import (
	"fmt"
	"gowizard/internal/utils"
	"os"
	"os/exec"
	"os/user"
)

func CreateAstroProject(name string) string {
	utils.ClearScreen()

	// Build the Docker command
	cmd := exec.Command("docker", "run", "--rm",
		"-v", fmt.Sprintf("%s:/app", utils.GetCurrentPath()),
		"-w", "/app",
		"-it",
		"dabiddo/larabox",
		"sh", "-c",
		fmt.Sprintf("pnpm create astro %s --template basics --install --git --yes && chown -R $(id -u):$(id -g) %s",
			name, name))

	// Set up pipes for real-time output
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	fmt.Println("Creating Astro project...")

	// Run the command
	if err := cmd.Run(); err != nil {
		fmt.Printf("\nError creating Astro project: %v\n", err)
		return "Failed to create Astro project"
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
	utils.CreateDevContainer(name, "_astro.stub")
	utils.CreateDockerfile(name, "_astro.stub")

	message := fmt.Sprintf("Astro project '%s' created successfully!\n", name)
	fmt.Print(message)
	return message
}

func CreateAstroBlogProject(name string) string {
	utils.ClearScreen()

	// Build the Docker command
	cmd := exec.Command("docker", "run", "--rm",
		"-v", fmt.Sprintf("%s:/app", utils.GetCurrentPath()),
		"-w", "/app",
		"-it",
		"larabox:latest",
		"sh", "-c",
		fmt.Sprintf("apk add --no-cache git && npm install -g pnpm && pnpm create astro@latest %s --template blog --install --git --yes && chown -R $(id -u):$(id -g) %s",
			name, name))

	// Set up pipes for real-time output
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	fmt.Println("Creating Astro Blog project...")

	// Run the command
	if err := cmd.Run(); err != nil {
		fmt.Printf("\nError creating Astro Blog project: %v\n", err)
		return "Failed to create Astro Blog project"
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
	utils.CreateDevContainer(name, "_astro.stub")
	utils.CreateDockerfile(name, "_astro.stub")

	message := fmt.Sprintf("Astro Blog project '%s' created successfully!\n", name)
	fmt.Print(message)
	return message
}
