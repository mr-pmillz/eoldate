FROM golang:1.24-alpine as builder

ENV GO111MODULE=on
RUN apk add --no-cache build-base

WORKDIR /app
COPY . /app
RUN go mod download
RUN GOOS=linux GOARCH=amd64 go build -v -trimpath -ldflags="-s -w" -o /eoldate .
RUN rm -rf /app

FROM alpine:latest

# Release
COPY --from=builder /eoldate /usr/local/bin/eoldate
RUN apk -U upgrade --no-cache \
    && apk add --no-cache bind-tools ca-certificates
RUN chmod +x /usr/local/bin/eoldate

ENTRYPOINT ["eoldate"]
