
FROM golang:1.10.3
#ENV GOPATH /Users/leb/project/gowork
#ENV GOROOT /usr/local/go
EXPOSE 8081
COPY . /go/src/golang-book/server
WORKDIR /go/src/golang-book/server
#COPY . .

RUN go install -v ./...

CMD [ "server" ]