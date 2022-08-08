FROM golang:1.19
ADD . /go/src/myapp
WORKDIR /go/src/myapp

COPY ./src/go.mod ./
COPY ./src/go.sum ./
RUN go mod download && go mod verify
COPY ./src/ . 
RUN go build -o /go/bin/myapp 
EXPOSE 8081
CMD [ "/go/bin/myapp" ]