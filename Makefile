GOCMD=go
GOBUILD=$(GOCMD) build
GOGET=$(GOCMD) get
BINARY_NAME=koikoin

install:
	@$(GOGET) -u github.com/dchest/uniuri
	@$(GOGET) -u github.com/shomali11/xredis

build:
	@$(GOBUILD) -o $(BINARY_NAME)
	@mkdir -p bin
	@mv $(BINARY_NAME) bin/

run: clean build
	@bin/./$(BINARY_NAME)

clean:
	@rm -rf $(BINARY_NAME)
