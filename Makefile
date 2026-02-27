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


build-windows: build-windows-amd64 build-windows-arm64

build-windows-amd64:
	@echo "Building for Windows/amd64..."
	@fyne-cross windows -arch=amd64 $(MAIN_PATH) 
	@echo ""

build-windows-arm64:
	@echo "Building for Windows/arm64..."
	@fyne-cross windows -arch=arm64 $(MAIN_PATH) 
	@echo ""


build-linux: build-linux-amd64 build-linux-arm64

build-linux-amd64:
	@echo "Building for Linux/amd64..."
	@fyne-cross linux -arch=amd64 $(MAIN_PATH)
	@echo ""

build-linux-arm64:
	@echo "Building for Linux/arm64..."
	@fyne-cross linux -arch=arm64 $(MAIN_PATH)
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
	@echo "  run                   Run the application locally"
	@echo "  test                  Run unit tests"
	@echo "  build-all             Build for all systems (Win, Linux, Android)"
	@echo "  build-windows         Build for Windows (amd64 & arm64)"
	@echo "  build-windows-amd64   Build for Windows x64 only"
	@echo "  build-windows-arm64   Build for Windows ARM only"
	@echo "  build-linux           Build for Linux (amd64 & arm64)"
	@echo "  build-linux-amd64     Build for Linux x64 only"
	@echo "  build-linux-arm64     Build for Linux ARM only"
	@echo "  build-android         Build for Android (apk)"
	@echo "  clean                 Remove fyne-cross artifacts"
	@echo "  deps                  Install Fyne tools and tidy modules"
