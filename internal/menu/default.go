package menu

import (
	"fmt"
	"gowizard/internal/projects"
	"log"

	"github.com/charmbracelet/huh"
)

func InitDefault() {

	form := huh.NewForm(

		huh.NewGroup(
			huh.NewInput().
				Title("Enter Project Name").
				Prompt("e.j. My Project: ").Value(&projectName),
		),

		huh.NewGroup(
			// Ask the user for a base burger and toppings.
			huh.NewSelect[string]().
				Title("Choose Project Preset").
				Options(
					huh.NewOption("Create Laravel Project with Composer", "laravel_composer"),
					huh.NewOption("Create Laravel Project with Laravel CLI", "laravel_cli"),
					huh.NewOption("Create Laravel Project with MySQ", "laravel_mysql"),
					huh.NewOption("Create Laravel Project with PostgreSQL", "laravel_pgsql"),
					huh.NewOption("Create Nuxt Project", "nuxt"),
					huh.NewOption("Create Nuxt Project with MySQL", "nuxt_mysql"),
					huh.NewOption("Create Nuxt Project with Pocketbase", "nuxt_pocketbase"),
					huh.NewOption("Create Astro Project", "astro_web"),
					huh.NewOption("CCreate Astro Blog Project", "astro_blog"),
					huh.NewOption("Create Refine.dev Project", "refine"),
					huh.NewOption("Create Tanstack Project", "better_stack"),
					huh.NewOption("Create NestJs Project", "nest"),
					huh.NewOption("Create Payload CMS Project", "payload_cms"),
					huh.NewOption("Create Rust Project", "rust"),
					huh.NewOption("Create HonoJs Project", "hono"),
					huh.NewOption("Create HonoJs OpenAPI Project", "hono_openapi"),
				).
				Value(&stackType), // store the chosen option
		),
	)

	err := form.Run()
	if err != nil {
		log.Fatal(err)
	}

	switch stackType {
	case "laravel_composer":
		projects.CreateLaravelProject(projectName)
		fmt.Println("You chose to create a Laravel project with Composer.")
	case "laravel_cli":
		fmt.Println("You chose to create a Laravel project with Laravel CLI.")
	case "laravel_mysql":
		fmt.Println("You chose to create a Laravel project with MySQL.")
	case "laravel_pgsql":
		fmt.Println("You chose to create a Laravel project with PostgreSQL.")
	case "nuxt":
		fmt.Println("You chose to create a Nuxt project.")
	case "nuxt_mysql":
		fmt.Println("You chose to create a Nuxt project with MySQL.")
	case "nuxt_pocketbase":
		fmt.Println("You chose to create a Nuxt project with Pocketbase.")
	case "astro_web":
		fmt.Println("You chose to create an Astro project.")
	case "astro_blog":
		projects.CreateAstroBlogProject(projectName)
	case "refine":
		fmt.Println("You chose to create a Refine.dev project.")
	case "better_stack":
		fmt.Println("You chose to create a Tanstack project.")
	case "nest":
		fmt.Println("You chose to create a NestJs project.")
	case "payload_cms":
		fmt.Println("You chose to create a Payload CMS project.")
	case "rust":
		fmt.Println("You chose to create a Rust project.")
	case "hono":
		fmt.Println("You chose to create a HonoJs project.")
	case "hono_openapi":
		fmt.Println("You chose to create a HonoJs OpenAPI project.")
	default:
		fmt.Println("Invalid option selected.")
	}
}
