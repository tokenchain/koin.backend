GOCMD=go
GOBUILD=$(GOCMD) build
GOGET=$(GOCMD) get
BINARY_NAME=koikoin

install:
	@$(GOGET) -u github.com/dchest/uniuri
	@$(GOGET) -u github.com/shomali11/xredis
	@$(GOGET) -u github.com/kataras/iris
	@$(GOGET) -u github.com/iris-contrib/middleware/tollboothic
	@$(GOGET) -u github.com/didip/tollbooth
	@$(GOGET) -u gopkg.in/gomail.v2

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
