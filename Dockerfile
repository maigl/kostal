FROM golang:1.22 AS builder

COPY go.mod .

ENV GO111MODULE=on
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -a -installsuffix cgo -o /go/bin/kostal ./cmd/frame

FROM scratch

COPY --from=builder /go/bin/kostal /kostal
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY web /web
COPY forecasts.json /forecasts.json

ENTRYPOINT ["/kostal"]
