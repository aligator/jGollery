FROM golang:1.12.0-alpine3.9

RUN mkdir /go/src/jGollery
ADD . /go/src/jGollery

WORKDIR /go/src/jGollery

RUN apk add git --no-cache

RUN go get -d ./...
RUN go build -o main ./cmd/jGolleryServer/main.go

CMD ["/app/main --addr :9090"]