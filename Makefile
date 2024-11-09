APP = myAccount
TARGET = ./target/app
SOURCE = ./cmd/start.go
ARGUMENT = -config=./etc/config.yaml

clean:
	@go clean
	@rm -rf $(TARGET)
build: clean
	@go build -o $(TARGET)/$(APP) $(SOURCE)
run: build
	@./$(TARGET)/$(APP) $(ARGUMENT)