BUILD			= go build
CLEAN			= go clean -i -cache
TEST			= go test
INSTALL 		= go install
BINARY_NAME 	= audit-sink

all: clean test build install
build: test
	$(BUILD) -o $(BINARY_NAME) -v
test:
	$(TEST) -v ./...
clean:
	$(CLEAN)
	rm -f $(BINARY_NAME)
run: build
	./$(BINARY_NAME)
install: clean test
	$(INSTALL)