#!/bin/bash

# Function to print an error message and exit
error_exit() {
    echo "Error: $1" >&2
    exit 1
}

# Ensure the script is run as root
if [ "$(id -u)" -ne 0 ]; then
    error_exit "This script must be run as root."
fi

# Update package list and install dependencies
echo "Updating package list and installing dependencies..."
apt-get update && apt-get install -y curl git build-essential libssl-dev
if [ $? -ne 0 ]; then
    error_exit "Failed to install required packages."
fi

# Install Nmap
if ! command -v nmap &> /dev/null; then
    echo "Installing Nmap..."
    apt-get install -y nmap
    if [ $? -ne 0 ]; then
        error_exit "Failed to install Nmap."
    fi
else
    echo "Nmap is already installed."
fi

# Install Nikto
if ! command -v nikto &> /dev/null; then
    echo "Installing Nikto..."
    apt-get install -y nikto
    if [ $? -ne 0 ]; then
        error_exit "Failed to install Nikto."
    fi
else
    echo "Nikto is already installed."
fi

# Install smtp-user-enum
if ! command -v smtp-user-enum &> /dev/null; then
    echo "Installing smtp-user-enum..."
    apt-get install -y smtp-user-enum
    if [ $? -ne 0 ]; then
        error_exit "Failed to install smtp-user-enum."
    fi
else
    echo "smtp-user-enum is already installed."
fi

# Install feroxbuster
if ! command -v feroxbuster &> /dev/null; then
    echo "Installing feroxbuster..."
    apt-get install -y feroxbuster
    if [ $? -ne 0 ]; then
        error_exit "Failed to install feroxbuster."
    fi
else
    echo "Feroxbuster is already installed."
fi

# Install Go (Golang)
if ! command -v go &> /dev/null; then
    echo "Installing Go (Golang)..."
    GO_VERSION=$(curl -s https://golang.org/VERSION?m=text)
    wget https://golang.org/dl/${GO_VERSION}.linux-amd64.tar.gz
    if [ $? -ne 0 ]; then
        error_exit "Failed to download Go."
    fi
    tar -C /usr/local -xzf ${GO_VERSION}.linux-amd64.tar.gz
    if [ $? -ne 0 ]; then
        error_exit "Failed to extract Go."
    fi
    rm ${GO_VERSION}.linux-amd64.tar.gz

    # Add Go to the PATH
    echo 'export PATH=$PATH:/usr/local/go/bin:$HOME/go/bin' >> /etc/profile
    source /etc/profile
else
    echo "Go is already installed."
fi

# Ensure ~/go/bin is in the PATH for the current session
if [[ ":$PATH:" != *":$HOME/go/bin:"* ]]; then
    export PATH="$HOME/go/bin:$PATH"
    # echo 'export PATH="$HOME/go/bin:$PATH"' >> ~/.bashrc
    echo 'export PATH="$HOME/go/bin:$PATH"' >> ~/.zshrc
    # source ~/.bashrc || true
    source ~/.zshrc || true
fi

# Install the Go project
echo "Installing the Go project..."
go install github.com/fancyc-bsi/gomap/gomap@latest
if [ $? -ne 0 ]; then
    error_exit "Failed to install the Go project."
fi

echo "Installation process completed."
