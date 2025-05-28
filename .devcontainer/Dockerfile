# Set up the Go development environment based on Debian Bookworm
FROM golang:bookworm

# Create a non-root user and set permissions
ARG USERNAME=vscode
ARG USER_UID=1000
ARG USER_GID=$USER_UID

# Install dependencies and create the user
RUN apt-get update && apt-get install -y \
    curl \
    git \
    openssh-client \
    build-essential && \
    groupadd --gid $USER_GID $USERNAME && \
    useradd --uid $USER_UID --gid $USER_GID --create-home $USERNAME && \
    apt-get clean && rm -rf /var/lib/apt/lists/*


# Set the default user and workspace
USER $USERNAME
WORKDIR /workspaces
