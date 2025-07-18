package menu

import (
	"bufio"
	"fmt"
	"gowizard/internal/projects"
	"log"
	"os"

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
			huh.NewSelect[string]().
				Title("Choose Project Preset").
				Options(
					huh.NewOption("Create Laravel Project with Composer", "laravel_composer"),
					huh.NewOption("Create Laravel Project with Laravel CLI", "laravel_cli"),
					huh.NewOption("Create Laravel Project with Starter Kit", "laravel_starter"),
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
					huh.NewOption("Create Nuxt Project 3", "nuxt_3"),
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
	case "laravel_starter":
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter starter Project (larave/react): ")
		starter, _ := reader.ReadString('\n')
		projects.CreateLaravelStarterProject(projectName, starter)
	case "laravel_cli":
		projects.CreateLaravelCLIProject(projectName)
	case "laravel_mysql":
		projects.CreateLaravelWithMySQL(projectName)
	case "laravel_pgsql":
		projects.CreateLaravelWithPostgreSQL(projectName)
	case "nuxt":
		projects.CreateNuxtProject(projectName)
	case "nuxt_mysql":
		projects.CreateNuxtWithMySQL(projectName)
	case "nuxt_pocketbase":
		projects.CreateNuxtWithPocketbase(projectName)
	case "astro_web":
		projects.CreateAstroProject(projectName)
	case "astro_blog":
		projects.CreateAstroBlogProject(projectName)
	case "refine":
		projects.CreateRefineProject(projectName)
	case "better_stack":
		projects.CreateTanstackProject(projectName)
	case "nest":
		projects.CreateNestJSProject(projectName)
	case "payload_cms":
		projects.CreatePayloadCMSProject(projectName)
	case "rust":
		projects.CreateRustProject(projectName)
	case "hono":
		projects.CreateHonoJsProject(projectName)
	case "hono_openapi":
		projects.CreateHonoOpenApi(projectName)
	case "nuxt_3":
		projects.CreateNuxtThreeProject(projectName)
	default:
		fmt.Println("Invalid option selected.")
	}
}
