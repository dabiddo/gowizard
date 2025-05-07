# GoWizard 🧙‍♂️
---

A modern CLI tool written in Go to create and manage development containers. GoWizard helps developers quickly set up consistent development environments using devcontainers.

## 🚀 Features

- Interactive TUI using charm.sh libraries
- Multiple development container templates
- Easy configuration and customization

## 📦 Installation

### From Source

```bash
# Clone the repository
git clone https://github.com/dabiddo/gowizard.git

# Navigate to project directory
cd gowizard

# Install dependencies
go mod tidy

# Build the binary
go build -o gowizard ./cmd/gowizard

# Make it executable
sudo chmod +x gowizard

# Move to PATH (optional)
sudo mv gowizard /usr/local/bin/
```

## 🎮 Usage

```bash
# Launch GoWizard
gowizard

# Follow the interactive prompts to:
# 1. Select your project type
# 2. Configure container settings
# 3. Generate devcontainer configuration
```

## 🔧 Requirements

- Docker installed and running
- VSCode with Remote Containers extension (for devcontainer support)

## 🤝 Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## 🙏 Acknowledgments

- Original [ContainerWizard](https://github.com/dabiddo/containerwizard) project
- [Charm](https://charm.sh/) for the amazing TUI libraries