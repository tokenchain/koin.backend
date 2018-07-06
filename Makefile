GOCMD=go
GOBUILD=$(GOCMD) build
GOGET=$(GOCMD) get
BINARY_NAME=koikoin

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

reset_docker:
	docker-compose rm --all && docker-compose pull && docker-compose build --no-cache && docker-compose up -d --force-recreate