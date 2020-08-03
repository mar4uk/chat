FROM golang:alpine

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /server
COPY . .
RUN go mod download
RUN go build -o main ./cmd
EXPOSE 8080
CMD ["/server/main"]