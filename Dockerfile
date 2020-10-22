FROM golang:1.14 as build
WORKDIR /app
COPY . /app
EXPOSE 8080
RUN go build -o /bin/demo
ENTRYPOINT ["/bin/demo", "&&"]