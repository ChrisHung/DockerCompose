FROM golang:1.12.7-alpine3.10 AS go-builder
ENV GO111MODULE=on
RUN apk --update add --no-cache git && \
    mkdir /go/src/hello
WORKDIR /go/src/hello
COPY . .
RUN go mod vendor && \
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main .

FROM alpine:latest
COPY --from=go-builder /go/src/hello/main /main
ENTRYPOINT [ "/main" ]