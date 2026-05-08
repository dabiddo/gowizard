package projects

import (
	"fmt"
	"gowizard/internal/utils"
	"os"
	"os/exec"
)

func UpdateImages(name string) string {
	utils.ClearScreen()

	//Run CMD docker to pull latest dabiddo/larabox image
	cmd := exec.Command("docker", "pull", "dabiddo/larabox:latest")

	// Set up pipes for real-time output
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	fmt.Println("Updating to latest Larabox Image...")

	// Run the command
	if err := cmd.Run(); err != nil {
		fmt.Printf("\nError Updating Image: %v\n", err)
		return "Failed to Update Image"
	}

	message := "Larabox Image Updated"
	fmt.Print(message)
	return message
}
