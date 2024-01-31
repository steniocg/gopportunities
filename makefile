.PHONY: default run build test docs clean
# Variable
APP_NAME=gopportunities

# Tasks
default: run-with-docs

run:
	@go run main.go
run-with-docs:
	@swag init
	@gorun main.go
build:
	@go build -0 $(APP_NAME) main.go
test:
	@go test ./ ...
docs:
	@swag init
clean:
	@rm -f $(APP_NAME)
	@rm - f ./docs