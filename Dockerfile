FROM golang:1.13-alpine

WORKDIR /app

COPY . .

RUN GO111MODULE=on go build -o grpc-server main.go

EXPOSE 8080
CMD ["./grpc-server"]