GOFMT=gofmt -w
GOBUILD=go build -v
SERVICE_PATH=${PWD}
SERVICE_NAME=Stock

.PHONY: all server service docker dependency clean

server: dependency
	$(GOFMT) $(SERVICE_PATH)/models/*.go
	$(GOFMT) $(SERVICE_PATH)/common/postgresqldriver/*.go
	$(GOFMT) $(SERVICE_PATH)/handler/*.go
	$(GOFMT) $(SERVICE_PATH)/router/*.go
	$(GOBUILD) $(SERVICE_PATH)/router/*.go
	$(GOBUILD) $(SERVICE_PATH)/models/*.go
	$(GOBUILD) $(SERVICE_PATH)/handler/*.go
	$(GOBUILD) $(SERVICE_PATH)/common/postgresqldriver/*.go

service: server
	$(GOFMT) $(SERVICE_PATH)/main.go
	$(GOBUILD) -o $(SERVICE_NAME) $(SERVICE_PATH)/main.go

all: service

docker:
	docker build -t $(SERVICE_NAME) -f Docker/Dockerfile .
		
dependency:
	go mod tidy

clean:
	rm Stock
