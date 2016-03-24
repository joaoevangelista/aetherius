FROM golang:alpine

RUN apk update && apk add git

RUN  go get github.com/tools/godep

ADD . /go/src/github.com/joaoevangelista/geocode

# install dependencies
RUN cd /go/src/github.com/joaoevangelista/geocode && /go/bin/godep get

RUN go install github.com/joaoevangelista/geocode

ENTRYPOINT /go/bin/geocode

EXPOSE 4000