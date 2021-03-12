FROM golang:alpine AS builder
ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64
RUN apk update && apk add --no-cache git
WORKDIR $GOPATH/src/worker
COPY worker.go .
COPY go.mod .
RUN go get -d -v
RUN go build -o /go/bin/worker

FROM scratch
ENV DIVISOR=3 OUTPUT_PHRASE=Fizz
COPY --from=builder /go/bin/worker /go/bin/worker
EXPOSE 8080
ENTRYPOINT ["/go/bin/worker"]