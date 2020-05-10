BINARY_LOCATION=bin
APP_NAME=crawler

.PHONY: build
build: clean
	@echo "Building service"
	protoc -I . service/*.proto --go_out=plugins=grpc:${GOPATH}/src

	@echo "Building server"
	cd server && go build -o ../$(BINARY_LOCATION)/$(APP_NAME)

.PHONY: run
start: clean build
	./$(BINARY_LOCATION)/$(APP_NAME)

.PHONY: clean
clean:
	if [ -d $(BINARY_LOCATION) ]; then rm -rf $(BINARY_LOCATION); fi
	if [ -f "crawler/*pb.go" ]; then rm service/*pb.go; fi
