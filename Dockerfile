FROM progrium/busybox
MAINTAINER Georg Kunz, cloudgear.co <georg@cloudgear.co>

COPY go-app    /app/
COPY templates /app/templates/
COPY public    /app/public/

EXPOSE 5000
WORKDIR /app
CMD ["/app/go-app"]
