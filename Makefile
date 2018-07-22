GOCMD=go
GOBUILD=$(GOCMD) build
GOGET=$(GOCMD) get
BINARY_NAME=koinkoin

install:
	@$(GOGET) -u ./...

build:
	@$(GOBUILD) -o $(BINARY_NAME)
	@mkdir -p bin
	@mv $(BINARY_NAME) bin/

build_win64:
	@GOOS=windows GOARCH=amd64 $(GOBUILD) -o $(BINARY_NAME)
	@mkdir -p bin
	@mv $(BINARY_NAME) bin/

run: clean build
	@bin/./$(BINARY_NAME)

clean:
	@rm -rf $(BINARY_NAME)

docker_build:
	@docker-compose build --no-cache koinkoin

docker_start:
	@docker-compose up -d --build

docker_reset:
	@docker-compose kill koinkoin && docker-compose build --no-cache koinkoin

docker_restart: docker_reset docker_start
