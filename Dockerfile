FROM golang:1.10.3-alpine
COPY . /go/src/app
WORKDIR /go/src/app
RUN go get -v -d
RUN go install
CMD ["app"]
