GOCMD = go
GOBUILD = $(GOCMD) build
GOCLEAN = $(GOCMD) clean
GOTEST = $(GOCMD) test
BINARY_NAME = ddocs-go

all: clean build

clean:
	@echo "Cleaning..."
	$(GOCLEAN)
	rm -f $(BINARY_NAME)

build:
	@echo "Building..."
	$(GOBUILD) -o $(BINARY_NAME) ./cmd/ddocs-go

test:
	@echo "Running tests..."
	$(GOTEST) -v ./...

run:
	@echo "Running..."
	./$(BINARY_NAME)

help:
	@echo "Available targets:"
	@echo "  - make all:         Build the project (default target)"
	@echo "  - make clean:       Clean the project"
	@echo "  - make build:       Compile the project"
	@echo "  - make test:        Run the tests"
	@echo "  - make run:         Execute the built binary"
	@echo "  - make help:        Show available targets"
