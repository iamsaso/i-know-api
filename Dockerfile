FROM golang

RUN go get -u gopkg.in/h2non/baloo.v1
RUN go get -u github.com/sasso/i-know-api

WORKDIR /go/src/github.com/sasso/i-know-api
ADD ./spec ./spec
