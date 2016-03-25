FROM golang:alpine

RUN apk update && apk add git

RUN  go get github.com/tools/godep

ADD . /go/src/github.com/joaoevangelista/aetherius

# install dependencies
RUN cd /go/src/github.com/joaoevangelista/aetherius && /go/bin/godep get

RUN go install github.com/joaoevangelista/aetherius

ENTRYPOINT /go/bin/aetherius

EXPOSE 4000
