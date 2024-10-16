#!/bin/bash

swag fmt
swag init --parseDependency true
go build && echo "Build successful" && ./backend
