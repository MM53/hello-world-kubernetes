############################
# build binary
############################
FROM golang:alpine AS builder
RUN apk update && \
    apk add --no-cache git

WORKDIR /tmp/hello-world-kubernetes

COPY . .
RUN go get -d -v hello-world-kubernetes/backend
RUN CGO_ENABLED=0 go build -o bin/backend hello-world-kubernetes/backend

############################
# build image
############################
FROM scratch
COPY --from=builder /tmp/hello-world-kubernetes/bin/backend /go/bin/backend
ENTRYPOINT ["/go/bin/backend"]
