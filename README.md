# Go (golang)
## Run Go program
Run the below command from root folder
```
go run src/cmd/main.go
```
## Build Go program
Run the below command from root folder
```
go build -v ./...
```
## Run Tests
Run the below command from root folder
```
go test -v ./...
```

#### If you get TCP port already in use error on Mac
`lsof -i -P | grep LISTEN | grep 8080`
and then `kill -9 PID` 

## Running the go app in docker
Install docker https://docs.docker.com/desktop/install/mac-install/

### Create the docker file
Please check the Dockerfile for more details

### Build image
`docker build -t go-docker-demo .`

### Check the image
`docker image ls`

### Run the image 
`docker run go-docker-demo`

### Docker login
`docker login -u sangopak dckr_pat_ebpyePSwg91BysUjqstYJpUNILQ`