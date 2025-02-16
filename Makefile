GOFMT=gofmt -w
GOBUILD=go build -v
SERVICE_PATH=${PWD}
SERVICE_NAME=stock:v1

.PHONY: all server service docker dependency clean

server: dependency
	@echo "Compiling Go files"
	$(GOFMT) $(SERVICE_PATH)/models/*.go
	$(GOFMT) $(SERVICE_PATH)/common/postgresqldriver/*.go
	$(GOFMT) $(SERVICE_PATH)/handler/*.go
	$(GOFMT) $(SERVICE_PATH)/router/*.go
	$(GOBUILD) $(SERVICE_PATH)/router/*.go
	$(GOBUILD) $(SERVICE_PATH)/models/*.go
	$(GOBUILD) $(SERVICE_PATH)/handler/*.go
	$(GOBUILD) $(SERVICE_PATH)/common/postgresqldriver/*.go

service: server
	@echo "Building service"
	$(GOFMT) $(SERVICE_PATH)/main.go
	$(GOBUILD) -o $(SERVICE_NAME) $(SERVICE_PATH)/main.go

all: service

docker:
	@echo "Building Stocks image"
	docker build -t $(SERVICE_NAME) -f Docker/Dockerfile .
		
dependency:
	@echo "Downloading dependencies"
	go mod tidy

clean:
	@if [ -f "$(SERVICE_NAME)" ]; then \
		rm $(SERVICE_NAME); \
		echo "Previous binary removed"; \
	else \
		echo "Previous binary doesn't exist"; \
	fi
