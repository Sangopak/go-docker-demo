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
```
lsof -i -P | grep LISTEN | grep 8080
```
and then 
```
kill -9 PID
```

## Running the go app in docker
Install docker https://docs.docker.com/desktop/install/mac-install/

### Create the docker file
Please check the Dockerfile for more details

### Build image
```
docker build -t go-docker-demo .
``` 
or if you want to use no cache
```
docker build --progress=plain --no-cache -t go-docker-demo .
```

### Check the image
```
docker image ls
```

### Run the image 
```
docker run go-docker-demo
```

### Run image in detached mode with hostport 8080 and container port 8080
```
docker run -d -p 8080:8080 go-docker-demo
```

### Check all running container
```
docker ps --all
```

### Stop running container 
```
docker stop <container-id>
```

### remove container
note you need to stop a container before removing it
```
docker rm <container-id1> <container-id2> ...
```

### Docker tag
```
docker image tag <container-id> sangopak/go-docker-demo:v1.0
```

### Docker login
```
docker login -u <docker-user> -p <docker-password>
```

### Docker push
```
docker push sangopak/go-docker-demo:v1.0
```
