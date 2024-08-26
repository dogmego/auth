LOCAL_BIN:=$(CURDIR)/bin

# Переменные
PROTOC = protoc
PROTOC_GEN_GO = $(GOPATH)/bin/protoc-gen-go
PROTOC_GEN_GO_GRPC = $(GOPATH)/bin/protoc-gen-go-grpc
PROTO_SRC = api/note_v1/note.proto
OUT_DIR = pkg\note_v1
PROTO_INCLUDE = api\note_v1

# Основное правило
all: generate

install-golangci-lint:
	GOBIN=$(LOCAL_BIN) go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.53.3

lint:
	GOBIN=$(LOCAL_BIN) golangci-lint run ./... --config .golangci.pipeline.yaml

# Правило для генерации Go-кода из .proto файла
generate:
	$(PROTOC) --go_out=$(OUT_DIR) --go_opt=paths=source_relative --go-grpc_out=$(OUT_DIR) --go-grpc_opt=paths=source_relative -I $(PROTO_INCLUDE) $(PROTO_SRC)

# Очистка сгенерированных файлов
clean:
	del /q $(OUT_DIR)\*.pb.go

build:
	GOOS=linux GOARCH=amd64 go build -o service_linux cmd/grpc_server/main.go

copy-to-server:
	scp service_linux root@31.41.155.2:

docker-build-and-push:
	docker buildx build --no-cache --platform linux/amd64 -t cr.selcloud.ru/dogmego/test-server:v0.0.1 .
	docker login -u token -p CRgAAAAA6Piyt8HACTD47uaBx83ULmV31rvgGDRx cr.selcloud.ru/dogmego
	docker push cr.selcloud.ru/dogmego/test-server:v0.0.1

