#!/bin/bash

#This script will clone a git repository and remove the .git directory
#Usage: run the script with the git url as the first argument

# Get the Git URL from the first command-line argument
git_url="$1"

# Extract the repository name from the Git URL
repo_name=$(basename -s .git "$git_url")

# Clone the repository and remove the .git directory
git clone --depth=1  "$git_url" "$repo_name" && rm -rf "./$repo_name/.git"