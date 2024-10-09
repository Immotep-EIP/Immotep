#!/bin/bash

git pull
if [ $? -ne 0 ]; then
    echo "git pull failed"
    exit 1
fi

cd Server/
go run github.com/steebchen/prisma-client-go migrate deploy
go run github.com/steebchen/prisma-client-go generate
if [ $? -ne 0 ]; then
    echo "Prisma migration failed"
    exit 1
fi
cd ..
