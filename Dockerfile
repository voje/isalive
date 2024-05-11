# TODO: use scratch or alpine as runtime, import CA

FROM docker.io/library/golang:alpine as builder

WORKDIR /project
COPY . .

RUN go build ./cmd/isalive/main.go 


FROM docker.io/library/alpine:latest

RUN apk --no-cache add ca-certificates \
    && update-ca-certificates

COPY --from=builder /project/main /isalive
ENTRYPOINT ["/isalive"]
