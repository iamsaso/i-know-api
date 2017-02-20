FROM golang

RUN go get -u gopkg.in/h2non/baloo.v1
RUN go get -u github.com/sasso/i-know-api

ADD ./docker-entrypoint.sh /docker-entrypoint.sh
ADD ./specs ./go/src/github.com/sasso/i-know-api/specs

WORKDIR /go/src/github.com/sasso/i-know-api
ENTRYPOINT ["/docker-entrypoint.sh"]
