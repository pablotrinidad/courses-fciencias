# Crawler microservice

This microservice exposes a gRPC server implementation of the [**FCCrawler**](proto/service.proto) service.

The [**FCCrawler**](proto/service.proto) service handles the content retrieval of UNAM's Faculty
 of Science majors, programs and courses offer listed in the official website [fciencias.unam.mx
 ](http://www.fciencias.unam.mx/).
 
**NOTE:** Please use this service responsibly since the HTTP calls to the official website are
 performed concurrently and might cause unwanted traffic loads. I'm not responsible for the
  usage of this package, this is put together as a learning exercise on microservices.
  
 ## How to use
 * `make start` to start server locally
 * `make build` to build source
 * `make build-docker-image` to build production docker image.
 
 Read more [here](Makefile).
 
 A gRPC web client like [gRPCox](https://github.com/gusaul/grpcox) can be used to call the service.
 
 ## Modules structure
 
 * `proto` contains the gRPC service definition using [**Protocol Buffers**](https://developers.google.com/protocol-buffers), and the corresponding generated files.  
 * `server` contains the TCP server hosting the gRPC service implementation
 * `service` contains an implementation of the [**FCCrawler**](proto/service.proto) service.   