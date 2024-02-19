# Bases for building and running the app
FROM golang:1.20-alpine AS builder-base
WORKDIR /go/src/github.com/m-butterfield/heatmap-creator/server
COPY go.* ./
RUN go mod download
ADD . /go/src/github.com/m-butterfield/heatmap-creator/server

FROM alpine:latest AS runner-base
WORKDIR /root

# Run build
FROM builder-base AS server-builder
RUN go build -o bin/server cmd/server/main.go

# Copy the built executable to the runner
FROM runner-base AS server
COPY --from=server-builder /go/src/github.com/m-butterfield/heatmap-creator/server/bin/ ./bin/
CMD ["bin/server"]
