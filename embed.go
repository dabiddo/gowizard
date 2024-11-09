package main

import "embed"

//go:embed stubs/devcontainer/*.stub stubs/compose/*.stub stubs/dockerfile/*.stub
var devContainerStubs embed.FS
