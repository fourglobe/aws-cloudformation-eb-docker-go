FROM scratch
WORKDIR /usr/src/app
COPY ./app app
EXPOSE 8080
CMD ["./app"]