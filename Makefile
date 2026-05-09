TARGET=evilginx
PACKAGES=core database log parser

.PHONY: all build clean
all: build

build:
	@go build -o ./build/$(TARGET) -mod=vendor main.go

# install target: copies binary to ~/bin for personal use
install: build
	@cp ./build/$(TARGET) $(HOME)/bin/$(TARGET)
	@echo "Installed $(TARGET) to $(HOME)/bin/"

clean:
	@go clean
	@rm -f ./build/$(TARGET)
