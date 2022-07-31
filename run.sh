#!/bin/sh

docker build -t http-request . && docker run -p 9191:8080 http-request