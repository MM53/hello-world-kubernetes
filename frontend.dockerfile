############################
# build binary
############################
FROM golang:alpine AS builder
RUN apk update && \
    apk add --no-cache git

WORKDIR /tmp/hello-world-kubernetes

COPY . .
RUN go get -d -v hello-world-kubernetes/frontend
RUN CGO_ENABLED=0 go build -o bin/frontend hello-world-kubernetes/frontend

############################
# build image
############################
FROM scratch
COPY frontend/assets /var/www/assets
COPY frontend/templates /var/www/templates
COPY --from=builder /tmp/hello-world-kubernetes/bin/frontend /go/bin/frontend
ENV ASSET_DIR=/var/www/assets TEMPLATE_DIR=/var/www/templates
ENTRYPOINT ["/go/bin/frontend"]
