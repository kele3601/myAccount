APP = myAccount
TARGET = ./target/app
SOURCE = ./cmd/start.go

clean:
	@go clean
	@rm -rf $(TARGET)
build: clean
	@go build -o $(TARGET)/$(APP) $(SOURCE)
run: build
	@cd $(TARGET) && ./$(APP)