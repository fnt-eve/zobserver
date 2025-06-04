FROM golang:1.24 AS builder

WORKDIR /build
COPY . .

ENV CGO_ENABLED=0
RUN go build -o  /go/bin/observer github.com/fnt-eve/zobserver/cmd/observer

FROM alpine:3.22

LABEL org.opencontainers.image.source=https://github.com/fnt-eve/zobserver

COPY  --from=builder /go/bin/observer /go/bin/observer

ENV PATH="${PATH}:/go/bin"

ENTRYPOINT [ "/go/bin/observer" ]