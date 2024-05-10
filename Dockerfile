FROM debian:latest AS build
WORKDIR /opt/app/
COPY app/ .
COPY test_all.sh .
RUN apt update 
RUN apt install -y golang
RUN apt-get install -y ca-certificates openssl
RUN cd api/src && rm 'go.mod' && rm 'go.sum' && go mod init src

ARG cert_location=/usr/local/share/ca-certificates
# Get certificate from "github.com"
RUN openssl s_client -showcerts -connect github.com:443 </dev/null 2>/dev/null|openssl x509 -outform PEM > ${cert_location}/github.crt
# Get certificate from "proxy.golang.org"
RUN openssl s_client -showcerts -connect proxy.golang.org:443 </dev/null 2>/dev/null|openssl x509 -outform PEM >  ${cert_location}/proxy.golang.crt
# Update certificates
RUN update-ca-certificates


WORKDIR /opt/app/api/src
RUN go get github.com/go-sql-driver/mysql
RUN go get github.com/gorilla/sessions
RUN go get golang.org/x/crypto/argon2
RUN go build -o ../build/api main.go
RUN ls -la
RUN md5sum /opt/app/api/build/api
CMD ["./test_all.sh && /opt/app/api/build/api"]
