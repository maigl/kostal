.PHONY: arm install all web


arm:
	GOOS=linux GOARCH=arm GOARM=5 go build -o kostal_arm ./cmd/frame

web:
	scp -r web logpi:~/kostal/

install: arm web
	scp kostal_arm logpi:~/kostal/ && rm kostal_arm

docker:
	docker build -t ghcr.io/maigl/kostal:latest .

docker-push:
	docker push ghcr.io/maigl/kostal:latest
