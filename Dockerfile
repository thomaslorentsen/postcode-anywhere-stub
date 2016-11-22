FROM golang:1.6-onbuild

RUN openssl genrsa -out /go/src/app/server.key 2048 && echo "GB\n\n\n\n\n\n" | openssl req -new -x509 -sha256 -key /go/src/app/server.key -out /go/src/app/server.pem -days 3650

EXPOSE 9040
