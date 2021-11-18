#write here your username
USERNAME :=
APP_NAME := k8s-go-app
VERSION := latest

#write here path for your project
PROJECT :=
GIT_COMMIT := $(shell git rev-parse HEAD)


all: run

run:
  go install -ldflags="-X '$(PROJECT)/version.Version=$(VERSION)' \
  -X '$(PROJECT)/version.Commit=$(GIT_COMMIT)'" && $(APP_NAME)

build_container:
  docker build --build-arg=GIT_COMMIT=$(GIT_COMMIT) --build-arg=VERSION=$(VERSION)  --build-arg=PROJECT=$(PROJECT)\
  -t docker.io/$(USERNAME)/$(APP_NAME):$(VERSION) .