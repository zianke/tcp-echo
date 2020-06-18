FROM golang:1.13 as builder
WORKDIR /
COPY main.go main.go
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GO111MODULE=on go build -a -o tcp-echo main.go

FROM scratch
COPY --from=builder /tcp-echo .
ENTRYPOINT ["/tcp-echo"]
