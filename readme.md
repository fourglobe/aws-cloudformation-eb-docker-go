# AWS Beanstalk Single Container (Golang)

### Relase Verson build
```
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -ldflags '-s' -o app main.go
```

### Docker Container Clean Up
```
docker rm -f `docker ps -a -q` 
```

### Docker Image Clean Up
```
docker rmi `docker images -q`
```

### Database Clean UP
```
rm -rf database && docker-compose restart db
```