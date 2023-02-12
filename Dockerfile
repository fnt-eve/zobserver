FROM alpine:3.16

COPY ./observer /go/bin/observer

ENTRYPOINT [ "/go/bin/observer" ]