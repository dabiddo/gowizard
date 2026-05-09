package projects

import (
	"fmt"
	"gowizard/internal/utils"
	"os"
	"os/user"
)

func CreateCppProject(name string) string {
	utils.ClearScreen()

	// Use native Go Mkdir
	if err := os.Mkdir(name, 0755); err != nil {
		return fmt.Sprintf("Error creating directory: %v", err)
	}

	fmt.Println("Creating C++ project...")

	// Change ownership of the project directory
	currentUser, err := user.Current()
	if err != nil {
		fmt.Printf("\nError getting current user: %v\n", err)
		return "Failed to get current user"
	}

	utils.ChangeOwnership(utils.GetCurrentPath(), currentUser.Username, name)

	// Create dev container files
	utils.CreateDevContainer(name, "_cpp.stub")
	utils.CreateDockerfile(name, "_cpp.stub")

	// Check for errors here!
	if err := utils.CopyHelloExample(name, "_main.stub", "main.cpp"); err != nil {
		fmt.Println(err)
	}
	if err := utils.CopyHelloExample(name, "_cmakelists.stub", "CMakeLists.txt"); err != nil {
		fmt.Println(err)
	}

	message := fmt.Sprintf("C++ project '%s' created successfully!\n", name)
	fmt.Print(message)
	return message

}
