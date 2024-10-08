#!/bin/bash

go run github.com/steebchen/prisma-client-go migrate dev
if [ $? -ne 0 ]; then
    echo "Prisma migration failed"
    exit 1
fi
