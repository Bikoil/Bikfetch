# Name of the output binary
BINARY = bikfetch

# Go commands
GO = go
GO_BUILD = $(GO) build

# Source file
SRC = bikfetch.go

# Installation directory
INSTALL_DIR = /usr/local/bin

# Default target
all: build

# Build the binary
build:
	$(GO_BUILD) -o $(BINARY) $(SRC)

# Install the binary
install: build
	mkdir -p $(INSTALL_DIR)
	cp $(BINARY) $(INSTALL_DIR)
	chmod +x $(INSTALL_DIR)/$(BINARY)

# Run the binary
run: build
	./$(BINARY)

# Clean the build
clean:
	rm -f $(BINARY)

# Phony targets
.PHONY: all build install run clean

