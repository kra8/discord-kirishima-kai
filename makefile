BUILD_DIR 	:= bin
NAME 		:= kirishima-kai
VERSION		:= v1.2.0
ZIP_DIR		:= "$(NAME)-$(VERSION)"

.PHONY: build
build-linux: clean
	@rm -rf $(ZIP_DIR)
	@GOOS=linux GOARCH=amd64 go build -o $(BUILD_DIR)/$(NAME)
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
