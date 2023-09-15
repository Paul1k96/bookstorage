FROM golang:1.21-alpine AS builder
COPY . /go/src/github.com/Paul1k96/bookstorage
WORKDIR /go/src/github.com/Paul1k96/bookstorage
RUN go mod download
RUN go build -ldflags="-w -s" -o /go/bin/server /go/src/github.com/Paul1k96/bookstorage/cmd/api

FROM alpine
COPY --from=builder /go/bin/server /go/bin/server
COPY ./static/swagger.json /app/static/swagger.json
COPY ./.env /app/.env

WORKDIR /app

ENTRYPOINT ["/go/bin/server"]
