FROM golang:latest

COPY ./ ./
ENV GOPATH=/

RUN go mod download

RUN go build -o pop-up-museum ./cmd/main.go
CMD ["./pop-up-museum"]