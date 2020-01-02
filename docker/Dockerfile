FROM golang:1.13.5-alpine

WORKDIR /app

COPY . .

RUN GO11MODULE=on go build -o grpc-example main.go

EXPOSE 8080

CMD ["./grpc-example"]