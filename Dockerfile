FROM alpine
MAINTAINER Georg Kunz, cloudgear.net <contact@cloudgear.net>

COPY go-app    /app/
COPY templates /app/templates/
COPY public    /app/public/

WORKDIR /app
CMD ["/app/go-app"]
