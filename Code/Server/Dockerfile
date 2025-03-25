FROM golang:1.22

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

ARG TARGETARCH

RUN go install github.com/swaggo/swag/cmd/swag@latest
RUN CGO_ENABLED=0 GOARCH=$TARGETARCH ./build.sh

# RUN curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.63.4
# RUN ./run_tests.sh no-interactive

ENV PORT='3001'
ENV GIN_MODE=release

EXPOSE 3001

CMD go run github.com/steebchen/prisma-client-go migrate deploy && /app/backend
