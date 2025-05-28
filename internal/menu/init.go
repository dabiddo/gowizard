package menu

import (
	"fmt"
	"gowizard/internal/projects"
	"gowizard/internal/utils"
	"log"
	"path/filepath"

	"github.com/charmbracelet/huh"
)

func InitMenu() {

	fmt.Println("Initializing Init menu...")

	form := huh.NewForm(
		huh.NewGroup(
			// Ask the user for a base burger and toppings.
			huh.NewSelect[string]().
				Title("Choose your burger").
				Options(
					huh.NewOption("Initialize PHP Project", "vanilla_php"),
					huh.NewOption("Initialize PHP Project with FrankenPHP", "franken_php"),
					huh.NewOption("Initialize NodeJs Projec", "nodejs"),
					huh.NewOption("Initialize Golang Project", "golang"),
				).
				Value(&stackType), // store the chosen option in the "burger" variable
		),
	)

	err := form.Run()
	if err != nil {
		log.Fatal(err)
	}

	currentDir := filepath.Base(utils.GetCurrentPath())
	switch stackType {
	case "vanilla_php":
		projects.PhpProject(currentDir)
	case "franken_php":
		projects.FrankenProject(currentDir)
	case "nodejs":
		projects.NodeProject(currentDir)
	case "golang":
		projects.GolangProject(currentDir)
	default:
		fmt.Println("Invalid option selected.")
	}

}
