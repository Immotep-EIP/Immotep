# Makefile for Keyz Server

TEST_PACKAGES	=	$(shell go list ./... | grep -ve prisma -ve docs -ve chatgpt -ve brevo -ve '^keyz/backend$$' -ve '^keyz/backend/services$$')
MODE 			?=	none

all: build

build:
	go run github.com/steebchen/prisma-client-go generate
	swag fmt
	swag init --parseDependency true
	go build

db_update:
	go run github.com/steebchen/prisma-client-go migrate deploy
	go run github.com/steebchen/prisma-client-go generate

db_migrate:
	go run github.com/steebchen/prisma-client-go migrate dev

db_reset:
	go run github.com/steebchen/prisma-client-go migrate reset
	go run prisma/seed.go

lint:
	golangci-lint run

test: lint
	gotestsum -- $(TEST_PACKAGES) -coverprofile cover.out
	@printf "Coverage: "
	@go tool cover -func=cover.out | grep total | awk '{print $$3}'

# test:
# 	@if [ "$(MODE)" = "debug" ]; then \
# 		go test -v $(TEST_PACKAGES) -coverprofile cover.out; \
# 	elif [ "$(MODE)" = "no-interactive" ]; then \
# 		go test $(TEST_PACKAGES) -coverprofile cover.out; \
# 		go tool cover -func=cover.out; \
# 	else \
# 		go test $(TEST_PACKAGES) -coverprofile cover.out; \
# 		go tool cover -func=cover.out; \
# 		go tool cover -html=cover.out; \
# 	fi

coverage:
	go tool cover -func=cover.out
	go tool cover -html=cover.out

run: build lint
	./backend

.PHONY: all db_migrate db_reset lint test coverage old_test run
