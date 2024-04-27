#!/bin/bash

# Define the remote Raspberry Pi address
PI_ADDRESS="pi@192.168.67.8"

# Define the local and remote directories
LOCAL_APP_DIR="./raspberry"
REMOTE_APP_DIR="~/"

# Name of the Go binary
APP_BINARY="weatherApp"

# Build the Go application for Raspberry Pi (ARM architecture)
echo "Building Go application for ARM..."
# chanche to loacl app dir
cd $LOCAL_APP_DIR

# go mod tidy
go mod tidy

env GOOS=linux GOARCH=arm GOARM=7 go build -o $APP_BINARY 

# Check if the build was successful
if [ $? -ne 0 ]; then
    echo "Failed to build the application."
    exit 1
fi

echo "Build successful."
cd ..
# Copy the binary and HTML template to the Raspberry Pi
echo "Copying files to the Raspberry Pi..."
scp $LOCAL_APP_DIR/$APP_BINARY $PI_ADDRESS:$REMOTE_APP_DIR
scp $LOCAL_APP_DIR/index.html $PI_ADDRESS:$REMOTE_APP_DIR

# Check if the copy was successful
if [ $? -ne 0 ]; then
    echo "Failed to copy files."
    exit 1
fi

echo "Files copied successfully."

# Run the application on the Raspberry Pi
echo "Starting the application on the Raspberry Pi..."
ssh $PI_ADDRESS "chmod +x $REMOTE_APP_DIR/$APP_BINARY && $REMOTE_APP_DIR/$APP_BINARY"

# Print the final status
if [ $? -eq 0 ]; then
    echo "Application started successfully."
else
    echo "Failed to start the application."
    exit 1
fi
