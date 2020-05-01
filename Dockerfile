FROM golang:1.14-alpine

WORKDIR /app
COPY . .

RUN apk update && \
    apk add git && \
    go get github.com/cespare/reflex

EXPOSE 8080
CMD ["reflex", "-c", "reflex.conf"]