package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

// AccessPath prints the current directory and its name and returns the formatted string
func AccessPath() string {
	// Get the current directory
	currentPath, err := os.Getwd()
	if err != nil {
		return fmt.Sprintf("Error getting current directory: %v", err)
	}

	// Extract the last component of the current directory path (the directory name)
	currentDir := filepath.Base(currentPath)

	// Return the formatted string
	return fmt.Sprintf("Current directory: %s\nCurrent directory name: %s", currentPath, currentDir)
}

func GetCurrentPath() string {
	currentPath, err := os.Getwd()
	if err != nil {
		return fmt.Sprintf("Error getting current directory: %v", err)
	}
	return currentPath
}

func ChangeOwnership(path string, user string, projectName string) {
	cmd := exec.Command("sudo", "chown", "-R", fmt.Sprintf("%s:%s", user, user), projectName)
	if err := cmd.Run(); err != nil {
		fmt.Println("Error changing ownership:", err)
		return
	}
	fmt.Println("Ownership changed to " + user + ":" + user + " for '" + projectName + "'")
}

// CreateDevContainer creates the .devcontainer folder inside the project directory
// and copies the specified stub file content to devcontainer.json
func CreateDevContainer(projectName string, stubName string) error {
	// Define the path for the .devcontainer directory
	devContainerDir := filepath.Join(projectName, ".devcontainer")

	// Create the .devcontainer directory
	if err := os.MkdirAll(devContainerDir, 0755); err != nil {
		return fmt.Errorf("failed to create folder %s: %w", devContainerDir, err)
	}
	fmt.Println("Folder '.devcontainer' created inside '" + projectName + "'")

	// Define paths for devcontainer.json and stub file
	devContainerJSONPath := filepath.Join(devContainerDir, "devcontainer.json")
	stubPath := filepath.Join("stubs", "devcontainer", stubName)

	// Read the stub file content from embedded files
	stubContent, err := devContainerStubs.ReadFile(stubPath)
	if err != nil {
		return fmt.Errorf("failed to read stub file %s: %w", stubPath, err)
	}

	// Replace $DIR with project name
	contentString := strings.ReplaceAll(string(stubContent), "$DIR", projectName)

	// Write the modified content to devcontainer.json
	if err := os.WriteFile(devContainerJSONPath, []byte(contentString), 0644); err != nil {
		return fmt.Errorf("failed to write to devcontainer.json: %w", err)
	}
	fmt.Println("devcontainer.json created successfully inside '" + projectName + "/.devcontainer'")

	return nil
}

// TouchFile creates an empty file at the specified path.
// If the file already exists, it updates the file's access and modification times.
func TouchFile(path string) error {
	file, err := os.OpenFile(path, os.O_RDONLY|os.O_CREATE, 0644)
	if err != nil {
		return fmt.Errorf("failed to touch file %s: %w", path, err)
	}
	defer file.Close()

	currentTime := time.Now().Local()
	return os.Chtimes(path, currentTime, currentTime)
}

// ClearScreen clears the terminal screen across different operating systems
func ClearScreen() {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}

// CreateDockerCompose creates the docker-compose.yml file in the .devcontainer directory
// and copies the specified stub file content to it
func CreateDockerCompose(projectName string, stubName string) error {
	// Define the path for docker-compose.yml inside .devcontainer
	dockerComposePath := filepath.Join(projectName, ".devcontainer", "docker-compose.yml")

	// Read the stub file content from embedded files
	stubPath := filepath.Join("stubs", "compose", stubName)
	stubContent, err := devContainerStubs.ReadFile(stubPath)
	if err != nil {
		return fmt.Errorf("failed to read stub file %s: %w", stubPath, err)
	}

	// Replace $DIR with project name
	contentString := strings.ReplaceAll(string(stubContent), "$DIR", projectName)

	// Write the modified content to docker-compose.yml
	if err := os.WriteFile(dockerComposePath, []byte(contentString), 0644); err != nil {
		return fmt.Errorf("failed to write to docker-compose.yml: %w", err)
	}
	fmt.Println("docker-compose.yml created successfully inside '" + projectName + "/.devcontainer'")

	return nil
}

// CreateDockerfile creates an empty Dockerfile in the .devcontainer directory
// and copies the specified stub file content to it
func CreateDockerfile(projectName string, stubName string) error {
	// Define the path for Dockerfile inside .devcontainer
	dockerfilePath := filepath.Join(projectName, ".devcontainer", "Dockerfile")

	// Read the stub file content from embedded files
	stubPath := filepath.Join("stubs", "dockerfile", stubName)
	stubContent, err := devContainerStubs.ReadFile(stubPath)
	if err != nil {
		return fmt.Errorf("failed to read stub file %s: %w", stubPath, err)
	}

	// Write the content to Dockerfile
	if err := os.WriteFile(dockerfilePath, stubContent, 0644); err != nil {
		return fmt.Errorf("failed to write to Dockerfile: %w", err)
	}
	fmt.Println("Dockerfile created successfully inside '" + projectName + "/.devcontainer'")

	return nil
}
