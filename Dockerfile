FROM golang:latest AS golang
WORKDIR /build
COPY . /build
RUN go build -o demo-exporter -ldflags "-linkmode 'external' -extldflags '-static'" .

FROM alpine:latest
WORKDIR /opt/demo-exporter
COPY --from=golang /build/demo-exporter /opt/demo-exporter/demo-exporter
ENTRYPOINT [ "/opt/demo-exporter/demo-exporter" ]
