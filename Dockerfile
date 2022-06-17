FROM golang AS builder

#This env var needs to be set for Gomodule
ENV GO111MODULE=on

#Creating the working directory
WORKDIR /work_planner

RUN apt-get install git

#Run the go mod download if only the changes occured in either of these 2 files
# Copy mod and sum file to workdir
COPY go.mod .
COPY go.sum .
RUN go mod download

#Copy your code to this working directory
COPY . .
#This will build the go server binary files
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build

#This is to spin up the base image to run the binaries of go
FROM alpine
COPY --from=builder /work_planner/* /work_planner/

COPY /start.sh /
COPY /.dockerignore /
RUN chmod 777 /.dockerignore
RUN chmod 777 /start.sh

ENTRYPOINT ["sh","/work_planner/start.sh"]

