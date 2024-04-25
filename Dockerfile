FROM golang:latest AS build
WORKDIR /opt/app/
COPY app/ .
RUN cd api/src && ls -la && go build -o ../build/api main.go
RUN ls -la
CMD ["/opt/app/api/build/api"]
