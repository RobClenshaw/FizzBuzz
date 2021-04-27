FROM golang:alpine AS builder
ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64
RUN apk update && apk add --no-cache git
WORKDIR $GOPATH/src/fizzbuzz
COPY fizzbuzz.go .
COPY go.mod .
RUN go get -d -v
RUN go build -o /go/bin/fizzbuzz

FROM scratch
COPY --from=builder /go/bin/fizzbuzz /go/bin/fizzbuzz
EXPOSE 80
ENTRYPOINT ["/go/bin/fizzbuzz"]