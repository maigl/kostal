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

# we only use one number for the version, eg. v1 or v2
# it's much easier to handle than a full semver
# for a new release we just increment the version number
# but we do this only if git is clean
release:
	@if [ -z "$(shell git status --porcelain)" ]; then \
		NEW_VERSION=$$(echo $(IMAGE_TAG) | awk -F. '{print $$1+1}'); \
		git tag -a "v$$NEW_VERSION" -m "Release v$$NEW_VERSION"; \
		git push origin "v$$NEW_VERSION"; \
		echo "Released new version: v$$NEW_VERSION"; \
	else \
		echo "Git is not clean, please commit or stash your changes before releasing."; \
	fi
