package menu

// FormData holds all the form input data
type FormData struct {
	Burger       string
	Toppings     []string
	SauceLevel   int
	Name         string
	Instructions string
	Discount     bool
}

// NewFormData creates a new instance of FormData with default values
func NewFormData() *FormData {
	return &FormData{}
}

// Global instance for backward compatibility during refactoring
var (
	burger       string
	toppings     []string
	sauceLevel   int
	name         string
	instructions string
	discount     bool
)
