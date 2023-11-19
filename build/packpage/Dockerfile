# Build stage
FROM golang:1.20-alpine3.17 AS build-stage

WORKDIR /app

COPY ./../../go.mod ./../../go.sum ./
RUN go mod download

COPY ./../../cmd/ ./cmd/
COPY ./../../internal/ ./internal/
RUN CGO_ENABLED=0 GOOS=linux go build -o app ./cmd/factorial_calculate

# Deploy the application binary into a lean image
FROM alpine:latest AS build-release-stage

WORKDIR /

COPY --from=build-stage /app/app /

ENTRYPOINT ["/app"]