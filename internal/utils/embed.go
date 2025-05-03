package utils

import "embed"

//go:embed stubs/compose/*.stub
//go:embed stubs/devcontainer/*.stub
//go:embed stubs/dockerfile/*.stub
var devContainerStubs embed.FS
