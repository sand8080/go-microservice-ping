FROM golang:latest
RUN go get github.com/sand8080/go-microservice-ping/...
ENTRYPOINT ["/go/bin/go-microservice-ping"]
