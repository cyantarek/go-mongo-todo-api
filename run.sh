#!/usr/bin/env bash
export CGO_ENABLED=0
cd ./src
go build -a -installsuffix cgo -ldflags '-s' -o ../bin/crud-api
cd ..
chmod +x ./bin/crud-api
docker build -t cyantarek/go-mongo-todo-api .
docker run -d -p 8078:8078 cyantarek/go-mongo-todo-api