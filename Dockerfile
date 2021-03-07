FROM golang:1.16.0-alpine AS builder
WORKDIR /go/src/app
COPY . .
RUN go install -v ./...

FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY --from=builder /go/bin/apigo /app
EXPOSE 8080
ENTRYPOINT ./app