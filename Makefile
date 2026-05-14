TARGET=evilginx
PACKAGES=core database log parser

.PHONY: all build clean install
all: build

build:
	@go build -o ./build/$(TARGET) -mod=vendor main.go

# install target: copies binary to ~/bin for personal use
install: build
	@mkdir -p $(HOME)/bin
	@cp ./build/$(TARGET) $(HOME)/bin/$(TARGET)
	@echo "Installed $(TARGET) to $(HOME)/bin/"

# uninstall target: removes binary from ~/bin
uninstall:
	@rm -f $(HOME)/bin/$(TARGET)
	@echo "Removed $(TARGET) from $(HOME)/bin/"

clean:
	@go clean
	@rm -f ./build/$(TARGET)
	@rm -f ./build/$(TARGET).log

# run target: build and launch with a local phishlets directory for testing
run: build
	@./build/$(TARGET) -p ./phishlets -t ./redirectors -debug
