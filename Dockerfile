### source build ###
FROM golang:1.11-alpine3.8 as build

COPY src /build

WORKDIR /build

RUN go get -d -v -t;\
    CGO_ENABLED=0 GOOS=linux go build -v -o /build/webserver

### Runtime build ###
FROM scratch

COPY --from=build /build/webserver /usr/local/bin/

COPY files /

WORKDIR /wwwroot

EXPOSE 8080

CMD ["/usr/local/bin/webserver", "-p", "8080", "-d", "/wwwroot"]
