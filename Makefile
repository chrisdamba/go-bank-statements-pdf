# Target for building the application
build:
    go build -o converter main.go

# Target for running the application
run: build
    ./converter

# Target for cleaning up generated files
clean:
    rm -f converter statement.pdf

# Default target (runs when you just type 'make')
default: run
