FROM golang:alpine AS builder
ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64
RUN apk update && apk add --no-cache git
WORKDIR /go/src/worker
COPY worker.go .
COPY go.mod .
COPY worker_test.go .
RUN go get -d -v
RUN go build -o /go/bin/worker
RUN go test -v

FROM scratch AS prod
ENV DIVISOR=3 OUTPUT_PHRASE=Fizz
COPY --from=builder /go/bin/worker /go/bin/worker
EXPOSE 80
ENTRYPOINT ["/go/bin/worker"]