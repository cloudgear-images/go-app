FROM alpine
MAINTAINER Georg Kunz, cloudgear.net <contact@cloudgear.net>

RUN apk add --no-cache bash git curl bind-tools htop

COPY go-app    /app/
COPY templates /app/templates/
COPY public    /app/public/

WORKDIR /app
CMD ["/app/go-app"]
