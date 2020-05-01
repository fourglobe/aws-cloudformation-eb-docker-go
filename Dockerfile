FROM golang:1.14-alpine
WORKDIR /usr/src/app
COPY ./app app
EXPOSE 8080
CMD ["./app"]