FROM golang:1.12.6-stretch

RUN go get github.com/cespare/reflex

WORKDIR /app

ENTRYPOINT ["reflex", "-c", "/app/reflex.conf"]