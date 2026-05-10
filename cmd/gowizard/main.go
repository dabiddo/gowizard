package main

import (
	"bufio"
	"fmt"
	"gowizard/internal/projects"
	"gowizard/internal/utils"
	"log"
	"os"

	"charm.land/huh/v2"
	"charm.land/lipgloss/v2"
)

// Styling
var (
	titleStyle  = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#00BFFF")).Padding(1, 0, 1, 2)
	optionStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#FFA500")).Padding(0, 0, 0, 2)
	activeStyle = optionStyle.Copy().Bold(true).Background(lipgloss.Color("#3333"))
	bannerStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#FFD700")).Bold(true)
)

// Banner function to display
func printBanner() string {
	banner := `

    
 ██████╗  ██████╗ ██╗    ██╗██╗███████╗ █████╗ ██████╗ ██████╗ 
██╔════╝ ██╔═══██╗██║    ██║██║╚══███╔╝██╔══██╗██╔══██╗██╔══██╗
██║  ███╗██║   ██║██║ █╗ ██║██║  ███╔╝ ███████║██████╔╝██║  ██║
██║   ██║██║   ██║██║███╗██║██║ ███╔╝  ██╔══██║██╔══██╗██║  ██║
╚██████╔╝╚██████╔╝╚███╔███╔╝██║███████╗██║  ██║██║  ██║██████╔╝
 ╚═════╝  ╚═════╝  ╚══╝╚══╝ ╚═╝╚══════╝╚═╝  ╚═╝╚═╝  ╚═╝╚═════╝ 
                                                                                
    
`
	return bannerStyle.Render(banner)
}

type projectType int

const (
	vanilla projectType = iota
	PHPFramework
	JSFramework
)

func (c projectType) String() string {
	return [...]string{"Vanilla Project", "PHP Framework", "JS Framework"}[c]
}

func main() {
	var category projectType
	type opts []huh.Option[string]

	var choice string
	var projectName string

	menuBanner := printBanner() + "\n" + utils.AccessPath() + "\n"
	// Then ask for a specific food item based on the previous answer.
	form := huh.NewForm(

		huh.NewGroup(

			huh.NewNote().
				Title("Welcome!").
				Description(menuBanner),

			huh.NewInput().
				Title("Enter Project Name").
				Prompt("e.j. My Project: ").Value(&projectName),

			huh.NewSelect[projectType]().
				Title("What Type of Project?").
				Value(&category).
				Options(
					huh.NewOption("Vanilla Project", vanilla),
					huh.NewOption("PHP Framework Project", PHPFramework),
					huh.NewOption("JS Framework Project", JSFramework),
				),

			huh.NewSelect[string]().
				Value(&choice).
				Height(7).
				TitleFunc(func() string {
					return fmt.Sprintf("Okay, what kind of %s do you want to make?", category)
				}, &category).
				OptionsFunc(func() []huh.Option[string] {
					switch category {
					case JSFramework:
						return []huh.Option[string]{
							huh.NewOption("Vuejs", "vuejs"),
							huh.NewOption("React", "react"),
							huh.NewOption("Nuxt 4", "nuxt"),
							huh.NewOption("Astro", "astro"),
						}
					case PHPFramework:
						return []huh.Option[string]{
							huh.NewOption("Laravel", "laravel"),
							huh.NewOption("Laravel CLI", "laravelcli"),
							huh.NewOption("Laravel Starterkit", "laravelstarterkit"),
							huh.NewOption("Symfony", "synfony"),
						}
					default:
						return []huh.Option[string]{
							huh.NewOption("C++", "cpp"),
							huh.NewOption("Golang", "golang"),
							huh.NewOption("PHP", "php"),
							huh.NewOption("Rust", "rust"),
							huh.NewOption("Zig", "zig"),
						}
					}
				}, &category),
		),
	)

	err := form.Run()
	if err != nil {
		log.Fatal(err)
	}

	switch choice {
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
	case "cakephp":
		projects.CreateCakePhpProject(projectName)
	case "vuejs":
		projects.CreateVueProject(projectName)
	case "golang":
		projects.CreateGolangProject(projectName)
	case "update":
		projects.UpdateImages(projectName)
	case "cpp":
		projects.CreateCppProject(projectName)
	case "zig":
		projects.CreateZigProject(projectName)
	default:
		fmt.Println("Invalid option selected.")
	}
}
