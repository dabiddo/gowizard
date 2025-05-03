package projects

import (
	"fmt"
	"gowizard/internal/utils"
)

func PhpProject(name string) {
	fmt.Printf("PHP Project '%s' has been successfully created.\n", name)
	utils.CreateDevContainerInCurrent("_laravel.stub")
	utils.CreateDockerfileInCurrent("_laravel.stub")
	utils.CreateDockerComposeInCurrent("_laravel.stub")
}

func FrankenProject(name string) {
	fmt.Printf("Franken Project '%s' has been successfully created.\n", name)
	utils.CreateDevContainerInCurrent("_laravel.stub")
	utils.CreateDockerfileInCurrent("_laravel.stub")
	utils.CreateDockerComposeInCurrent("_laravel.stub")
}

func NodeProject(name string) {
	fmt.Printf("Node Project '%s' has been successfully created.\n", name)
	utils.CreateDevContainerInCurrent("_nuxt.stub")
	utils.CreateDockerfileInCurrent("_nuxt.stub")
}
