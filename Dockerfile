# get the apline golang image
FROM golang:1.21-alpine

# create a new dir
WORKDIR /app

# copy the go mod file and run go mod in docker
COPY go.mod ./
RUN go mod download

# copy the go files
COPY *.go ./

# build go executable in docker with output
RUN go build

# expose container port 8080
EXPOSE 8080

# execute the go executable 
CMD [ "/app/go-docker-demo" ]