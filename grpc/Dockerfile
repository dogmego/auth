FROM golang:1.20.3-alpine AS builder

COPY . /github.com/dogmego/microservices_course/week_2/grpc/source/
WORKDIR /github.com/dogmego/microservices_course/week_2/grpc/source/

RUN go mod download
RUN go mod tidy
RUN go build -o ./bin/crud_server cmd/grpc_server/main.go

FROM alpine:latest

WORKDIR /root/
COPY --from=builder /github.com/dogmego/microservices_course/week_2/grpc/source/bin/crud_server .

CMD ["./crud_server"]