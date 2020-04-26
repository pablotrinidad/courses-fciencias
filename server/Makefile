all: compile_protos

compile_protos:
	@echo "Compiling crawler protobuf files..."
	rm crawler/*pb.go
	protoc -I . crawler/*.proto --go_out=plugins=grpc:${GOPATH}/src
