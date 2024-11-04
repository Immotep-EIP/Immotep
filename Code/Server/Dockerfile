FROM golang:1.22

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

ARG TARGETARCH

ENV PORT='3001'
ENV GIN_MODE=release

RUN go install github.com/swaggo/swag/cmd/swag@latest
RUN CGO_ENABLED=0 GOARCH=$TARGETARCH ./build.sh

RUN ./run_tests.sh no-interactive

EXPOSE 3001

CMD go run github.com/steebchen/prisma-client-go migrate deploy ; ./backend
