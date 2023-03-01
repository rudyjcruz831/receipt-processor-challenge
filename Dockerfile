# start from golang:latest
FROM golang:alpine as builder

# setting enviroment viarable for grpc
ENV GO111MODULE=on

RUN mkdir /fetch_server

# set the curretn working directory inside the container
WORKDIR /fetch_server

# copy current directory to the container   
COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . /fetch_server

# build the go server 
RUN CGO_ENABLED=0 GOOS=linux go build -o serverexec main.go
EXPOSE 50052

CMD ./serverexec