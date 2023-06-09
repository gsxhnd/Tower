SHELL := /bin/bash
BASEDIR = $(shell pwd)
APP = main
BuildDIR = build

# gitTag = $(shell if [ "`git describe --tags --abbrev=0 2>/dev/null`" != "" ];then git describe --tags --abbrev=0 | sed 's/v//g'; else git log --pretty=format:'%h' -n 1; fi)
# buildDate = $(shell TZ=Asia/Shanghai date +%FT%T%z)
# gitCommit = $(shell git log --pretty=format:'%H' -n 1)
# gitTreeState = $(shell if git status|grep -q 'clean';then echo clean; else echo dirty; fi)
# versionDir = "github.com/gsxhnd/garage"
# ldflags= "-X ${versionDir}.gitTag=${gitTag} \
# -X ${versionDir}.buildDate=${buildDate} \
# -X ${versionDir}.gitCommit=${gitCommit} \
# -X ${versionDir}.gitTreeState=${gitTreeState}"

all: release_linux

release_linux:
	go clean
#	Build for linux
#	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -trimpath  -ldflags ${ldflags} -o ${BuildDIR}/${APP}-linux64-amd64 ./src
#	CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -v -trimpath  -ldflags ${ldflags} -o ${BuildDIR}/${APP}-linux64-arm64 ./src
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -trimpath  -o ${BuildDIR}/${APP} ./src

# release_win:
# 	# Build for win
# 	go clean
# 	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -v -trimpath -ldflags ${ldflags} -o ${BuildDIR}/${APP}-windows-amd64.exe .
# 	CGO_ENABLED=0 GOOS=windows GOARCH=arm go build -v -trimpath -ldflags ${ldflags} -o ${BuildDIR}/${APP}-windows-arm.exe .

di: 
	wire ./src/wire

clean:
	@go clean --cache
	@rm -rvf build/*

.PHONY: release_linux release_win clean