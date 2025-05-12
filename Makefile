.PHONY: arm install all web

IMAGE_TAG := $(shell git describe --tags --abbrev=0)

build:
	go build -o kostal ./cmd/frame

arm:
	GOOS=linux GOARCH=arm GOARM=5 go build -o kostal_arm ./cmd/frame

web:
	scp -r web logpi:~/kostal/

install: arm web
	scp kostal_arm logpi:~/kostal/ && rm kostal_arm

docker:
	docker build -t ghcr.io/maigl/kostal:$(IMAGE_TAG) .

docker-push: docker
	docker push ghcr.io/maigl/kostal:$(IMAGE_TAG)
