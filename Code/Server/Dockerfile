FROM golang:1.22 AS build-stage

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

ARG TARGETARCH

RUN go install github.com/swaggo/swag/cmd/swag@latest
RUN CGO_ENABLED=0 GOARCH=$TARGETARCH ./build.sh

FROM build-stage AS test-stage

RUN ./run_tests.sh no-interactive

FROM gcr.io/distroless/base-debian11 AS build-release-stage

WORKDIR /

COPY --from=build-stage /app/backend /immotep_backend

ENV PORT='3001'
ENV GIN_MODE=release

EXPOSE 3001

USER nonroot:nonroot

CMD go run github.com/steebchen/prisma-client-go migrate deploy && ./immotep_backend
