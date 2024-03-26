#!/usr/bin/env bash

CURRENT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
CURRENT_TIME=$(date +"%Y-%m-%d_%H-%M-%S")
SCRIPT_NAME=$(basename $0)

echo "Script name: $SCRIPT_NAME"
echo "Current directory: $CURRENT_DIR"
echo "Current time: $CURRENT_TIME"
