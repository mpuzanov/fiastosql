SOURCE=./cmd/fiastosql
APP=fias-to-sql
VERSION=1.0
ARCH= $(shell uname -m)
GOBASE=$(shell pwd)
RELEASE_DIR=$(GOBASE)/bin

.DEFAULT_GOAL = build 

GO_SRC_DIRS := $(shell \
	find . -name "*.go" -not -path "./vendor/*" | \
	xargs -I {} dirname {}  | \
	uniq)
GO_TEST_DIRS := $(shell \
	find . -name "*_test.go" -not -path "./vendor/*" | \
	xargs -I {} dirname {}  | \
	uniq)	

build: 
	@go build -v -o ${APP} ${SOURCE}

lint:
	@goimports -w ${GO_SRC_DIRS}	
	@gofmt -s -w -d ${GO_SRC_DIRS}
	@golint ${GO_SRC_DIRS}
	@go vet ${GO_SRC_DIRS}
	@#golangci-lint run

test:
	go test -v ${GO_TEST_DIRS}

mod:
	go mod verify
	go mod tidy

run:
	@#go run ${SOURCE} -d H:\fias\fias18 -r 18
	@#go run ${SOURCE} -d H:\\fias\\fias18 -r 18 -v --db_url postgres://postgres:123@localhost:5432/fias?sslmode=disable
	@go run ${SOURCE} -d H:\\fias\\fias18 -r 18 -v --db_url sqlserver://sa:123@localhost?database=fias
	@#go run ${SOURCE} -d H:\\fias\\fias18 -r 18 -v --db_url root:dnypr1@/fias
	@#DIR="H:\fias\fias18" REGION="18" go run ${SOURCE} 

release:
	GOOS=windows GOARCH=amd64 go build -ldflags="-H windowsgui" -o ${APP}.exe ${SOURCE}
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o ${APP} ${SOURCE}

.PHONY: build run release lint test mod
