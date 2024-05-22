FROM golang:latest AS build
WORKDIR /opt/app/
COPY app/ .
COPY test_all.sh .
#RUN apt update 
#RUN apt install -y golang
#RUN apt-get install -y ca-certificates openssl
RUN cd api/src && rm 'go.mod' && rm 'go.sum' && go mod init src


WORKDIR /opt/app/api/src

# BUILD
RUN go get github.com/go-sql-driver/mysql
RUN go get github.com/gorilla/sessions
RUN go get golang.org/x/crypto/argon2
RUN go get github.com/gorilla/websocket
RUN go build -o ../build/api main.go
# BUILD

RUN ls -la
RUN md5sum /opt/app/api/build/api
CMD ["/opt/app/api/build/api"]
