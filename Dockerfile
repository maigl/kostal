FROM golang:1.24 AS builder

# use modules
COPY go.mod .

ENV GO111MODULE=on
RUN go mod download

COPY . .

# Build the binary
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -a -installsuffix cgo -o /go/bin/kostal ./cmd/frame

FROM gcr.io/distroless/base

# Copy our static executable
COPY --from=builder /go/bin/kostal /go/bin/kostal
COPY web /go/bin/web
COPY forecasts.json /go/bin/forecasts.json

WORKDIR /go/bin

ENTRYPOINT ["/go/bin/kostal"]
