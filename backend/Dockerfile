
FROM golang:latest AS builder
RUN apt-get update
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64
WORKDIR /go/src
COPY go.mod .
RUN go mod download
COPY . .

EXPOSE 8001

#ENV DB_USER="admin"
#ENV DB_PASSWORD="101250"
#ENV DB_HOST="172.17.0.3"
#ENV DB_PORT="27017"
#ENV DB_NAME="db_books"
#ENV SERVER_PORT="4042"
#ENV GO_ENV="production"

RUN go build main.go
FROM scratch
COPY --from=builder /go/src .
ENTRYPOINT  ["./main"]
