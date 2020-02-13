FROM golang:1.13 as builder
WORKDIR /the-app
COPY main.go main.go
RUN CGO_ENABLED=0 go build -a -o the-app main.go

FROM scratch
COPY --from=builder /the-app/the-app /the-app
COPY data.txt /data.txt
WORKDIR /
ENTRYPOINT ["/the-app"]