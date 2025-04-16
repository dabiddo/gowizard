package main

import "fmt"

func phpProject(name string) {
	fmt.Printf("PHP Project '%s' has been successfully created.\n", name)
	CreateDevContainer(name, "_laravel.stub")
}

func frankenProject(name string) {
	fmt.Printf("Franken Project '%s' has been successfully created.\n", name)
	CreateDevContainer(name, "_laravel.stub")
}

func nodeProject(name string) {
	fmt.Printf("Node Project '%s' has been successfully created.\n", name)
	CreateDevContainer(name, "_laravel.stub")
}
