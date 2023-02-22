#!/bin/bash

echo "> docker build -t forum ."
docker build -t forum .

echo "> docker run -d -p 80:8080 forum"
docker run -d -p 80:8080 forum

echo "Server started on http://localhost"