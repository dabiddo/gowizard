package utils

import "embed"

//go:embed stubs/compose/*.stub
//go:embed stubs/devcontainer/*.stub
//go:embed stubs/dockerfile/*.stub
//go:embed stubs/configs/*.stub
//go:embed stubs/vscode/*.stub
//go:embed stubs/examples/*.stub
var devContainerStubs embed.FS
