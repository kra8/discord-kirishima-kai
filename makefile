BUILD_DIR 	:= bin
NAME 		:= discord-notify-voice-join
VERSION		:= v2.0.0
ZIP_DIR		:= "$(NAME)-$(VERSION)"

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

clean:
	rm -rf ./bin
