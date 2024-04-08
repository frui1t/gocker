GockerDir:=./gocker
GockerBin:=./build/gocker

VERSION:=1.0
GITCOMMIT:=$(shell git rev-parse --short HEAD)
BUILT:=$(shell date +"%Y-%m-%d %H:%M:%S")

LDFLAGS="\
-X 'gocker/gocker/version.Version=$(VERSION)'\
-X 'gocker/gocker/version.GitCommit=$(GITCOMMIT)'\
-X 'gocker/gocker/version.BuildTime=$(BUILT)'"

build:
	go mod tidy
	@echo "开始编译"
	@echo "编译Gocker"
	GOOS=linux GOARCH=amd64 go build -ldflags $(LDFLAGS) -o $(GockerBin) $(GockerDir)/cmd/gocker

.PHONY: build
