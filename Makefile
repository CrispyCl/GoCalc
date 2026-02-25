# Project variables
APP_NAME := GoCalc
MAIN_PATH := ./cmd/calc


.PHONY: run test build-all build-windows build-linux build-android clean deps help

run:
	@echo "Running the application..."
	@go run $(MAIN_PATH)/main.go

test:
	@echo "Running tests..."
	@go test -v ./...


## --- Fyne-Cross Builds ---

build-all: build-windows build-linux build-android

build-windows:
	@echo "Building for Windows..."
	@fyne-cross windows $(MAIN_PATH)
	@echo ""

build-linux:
	@echo "Building for Linux..."
	@fyne-cross linux $(MAIN_PATH)
	@echo ""

build-android:
	@echo "Building for Android..."
	@fyne-cross android $(MAIN_PATH)
	@echo ""


## --- Utilities ---

clean:
	@echo "Cleaning up..."
	@rm -rf fyne-cross

deps:
	@echo "Installing dependencies..."
	go mod tidy
	go install fyne.io/tools/cmd/fyne@latest
	go install github.com/fyne-io/fyne-cross@latest

help:
	@echo "Usage: make [target]"
	@echo ""
	@echo "Targets:"
	@echo "  run             Run the application locally"
	@echo "  test            Run unit tests for internal packages"
	@echo "  build-all       Build for all systems"
	@echo "  build-windows   Build for Windows"
	@echo "  build-linux     Build for Linux"
	@echo "  build-android   Build for Android"
	@echo "  clean           Remove build artifacts"
	@echo "  deps            Install tools"
