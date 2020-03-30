all: compile_protos

compile_protos:
	@echo "Compiling crawler protobuf files..."
	rm server/crawler/*pb.go
	protoc -I . server/crawler/*.proto --go_out=plugins=grpc:${GOPATH}/src
