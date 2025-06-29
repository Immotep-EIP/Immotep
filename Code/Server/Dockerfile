# Build stage
FROM golang:1.24-alpine AS builder

WORKDIR /app

RUN apk add --no-cache make

COPY go.mod .
COPY go.sum .

RUN go mod download

# prefetch the binaries, so that they will be cached and not downloaded on each change
RUN go run github.com/steebchen/prisma-client-go@latest prefetch

COPY . .

ARG TARGETARCH

RUN go install github.com/swaggo/swag/cmd/swag@latest
RUN CGO_ENABLED=0 GOARCH=$TARGETARCH make build

# Final stage
FROM alpine:latest

WORKDIR /app

RUN apk add --no-cache curl

COPY --from=builder /app/backend /app/backend
COPY --from=builder /app/assets /app/assets

ENV PORT='3001'
ENV GIN_MODE=release

EXPOSE 3001

RUN adduser -D myuser
USER myuser

HEALTHCHECK --interval=10s --timeout=5s --retries=5 CMD curl -f http://localhost:3001/ || exit 1

CMD ["/app/backend"]
