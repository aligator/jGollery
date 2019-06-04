FROM golang:1.12.0-alpine3.9

RUN apk add git shadow --no-cache

# do not use root user inside of docker
RUN groupadd -g 1050 -r appuser && useradd --create-home -u 1050 --no-log-init -r -g appuser appuser \
 && chown -R appuser:appuser /go
USER appuser

# setup jGollery
RUN mkdir /go/src/jGollery
WORKDIR /go/src/jGollery

# add all except static/galllery (.dockerignore) which should be binded with
# docker run --mount type=bind,source="$(pwd)"/static/gallery,target=/go/src/jGollery/static/gallery
ADD . /go/src/jGollery

# commpile
RUN go get -d ./...
RUN go build -o jGollery ./cmd/jGolleryServer/main.go

# remove make deps
USER root
RUN apk del git shadow --no-cache
USER appuser

CMD ["/go/src/jGollery/jGollery",  "--addr", ":9090"]

EXPOSE 9090

# full run-command to expose port and mount gallery
# docker run -p 9090:9090 --mount type=bind,source="$(pwd)"/static/gallery,target=/go/src/jGollery/static/gallery