# Cloudgear Example App

Simple example app to output the hostname and environment of a container. Written in Go to keep the image size minimal.

## Run

Run the image with Docker

    docker run -d -p 5000/tcp cloudgear/go-app

## Build

Build the Busybox based image with Docker:

    docker run --rm -v "$(pwd)":/usr/src/github.com/cloudgear-images/go-app -w /usr/src/github.com/cloudgear-images/go-app golang:1.4 /bin/sh -c "go get github.com/go-martini/martini && go get github.com/martini-contrib/render && go build -v"

Or just run 

    rake compile