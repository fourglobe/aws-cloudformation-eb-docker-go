# AWS Beanstalk Single Container (Golang)

### Relase Verson build
```
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -ldflags '-s' -o app application.go
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

## Cloud Formation 

parameter setting

```json
{
  "ApplicationName":"USER APPLICATION NAME", 
  "EnvironmentName":"UER ENV NAME(TEST, STAGING, PROD ,,, etc)",
  "InstanceType":"インスタンスタイプ",
  "KeyPairName":"EC2 Key Name",
  "VpcId":"VPCID",
  "SubnetId":"SUBNET ID LIST",
  "ELBSubnetId":"ELB SUBNET ID",
  "MinSize":"1",
  "MaxSize":"2",
  "Cooldown":"180"
}
