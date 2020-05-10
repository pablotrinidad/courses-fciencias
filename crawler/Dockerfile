FROM golang:1.14 as base
ENV APP_PORT=8080
WORKDIR /go/src/crawler
COPY . .

FROM base as dev
RUN go get -u github.com/cosmtrek/air
ENTRYPOINT "air"

