.PHONY: arm install all


arm:
	GOOS=linux GOARCH=arm GOARM=5 go build -o kostal_arm ./cmd/frame

install: arm
	scp ./cmd/frame/kostal_arm pi@192.168.0.47:/tmp/
