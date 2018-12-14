BUILD_DIR 	:= bin
NAME 		:= kirishima-kai
GOOS 		:= "linux"
GOARCH		:= "amd64"
VERSION		:= v1.1.1
ZIP_DIR		:= "$(NAME)-$(VERSION)"

.PHONY: build
build: clean
	@go build -o $(BUILD_DIR)/$(NAME)
	@mkdir $(ZIP_DIR)
	@cp token.example $(ZIP_DIR)/token.example
	@cp $(BUILD_DIR)/$(NAME) $(ZIP_DIR)/$(NAME)
	@zip $(ZIP_DIR).zip $(ZIP_DIR)/*
	@rm -rf $(ZIP_DIR)

.PHONY: run
run:
	@./$(BUILD_DIR)/$(NAME)

.PHONY: test
test: build run

.PHONY: clean
clean:
	rm -rf ./bin
