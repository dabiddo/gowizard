package projects

import (
	"fmt"
	"gowizard/internal/utils"
	"os"
)

func CreateGolangProject(name string) string {
	utils.ClearScreen()

	//create project dir
	if err := os.MkdirAll(name, 0755); err != nil {
		fmt.Printf("failed to create folder %s:", name)
	}
	fmt.Println("Folder '.devcontainer' created inside '" + name + "'")

	utils.CreateDevContainer(name, "_golang.stub")
	utils.CreateDockerfile(name, "_golang.stub")

	message := fmt.Sprintf("Golang project '%s' created successfully!\n", name)
	fmt.Print(message)
	return message
}
