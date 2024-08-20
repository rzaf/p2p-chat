ifeq ($(OS),Windows_NT) 
    DETECTED_OS := Windows
	FILE := chat.exe
	PB_FILE := pb\chat.proto
else
    DETECTED_OS := $(shell uname)
	FILE := chat
	PB_FILE := pb/chat.proto
endif

all: generatePbs build

generatePbs:
	@echo "generating protocol buffer files ... "
	@protoc ${PB_FILE} --go_out=. --go-grpc_out=.

download-modules:
	@echo "downloading required go modules if not cached ....."
	@go mod download

bundle:
ifeq ($(DETECTED_OS),Windows)
	@fyne bundle -o gui\config\bundled.go --pkg config .\assets\chat-icon.png
else
	@fyne bundle -o gui/config/bundled.go --pkg config ./assets/chat-icon.png
endif

build: download-modules
ifeq ($(DETECTED_OS),Windows)
	@if not exist bin mkdir bin
	@echo "building go files ..... (for first time will take a while)"
	@go build -o .\bin\${FILE} cmd\gui\main.go 
else
	@mkdir -p bin
	@echo "building go files ..... (for first time will take a while)"
	@go build -o ./bin/${FILE} cmd/gui/main.go 
endif

clean:
	@echo "Removing bin directory ... "
ifeq ($(DETECTED_OS),Windows)
	@if exist bin rmdir /s /q bin 
else
	@rm -rf bin 
endif
