BINARY_LOCATION=bin
APP_NAME=crawler
DOCKER_IMAGE_NAME=fciencias-crawler

.PHONY: build
build: clean
	@echo "Building service"
	protoc -I . proto/*.proto --go_out=plugins=grpc:${GOPATH}/src

	@echo "Building server"
	cd server && go build -o ../$(BINARY_LOCATION)/$(APP_NAME)

.PHONY: start
start: clean build
	./$(BINARY_LOCATION)/$(APP_NAME)

.PHONY: clean
clean:
	if [ -d $(BINARY_LOCATION) ]; then rm -rf $(BINARY_LOCATION); fi
	if [ -f "proto/*pb.go" ]; then rm proto/*pb.go; fi

.PHONY: build-docker-image
build-docker-image:
	docker build -t $(DOCKER_IMAGE_NAME) .
	@echo "Image build successfully"
	docker images $(DOCKER_IMAGE_NAME):latest