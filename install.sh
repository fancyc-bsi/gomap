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

echo "Installation process completed."
