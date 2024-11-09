# Set up the Go development environment based on Debian Bookworm
FROM golang:bookworm

# Install dependencies
RUN apt-get update && apt-get install -y \
    curl \
    git \
    build-essential

# Set the default user and workspace
#USER vscode
WORKDIR /workspace
