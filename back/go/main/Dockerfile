FROM golang:1.9.2-alpine3.6

WORKDIR /go/src/github.com/RepGarage/refrigerator/back/go/main

ADD . /go/src/github.com/RepGarage/refrigerator/back/go/main

RUN apk add --no-cache git

RUN ./install-deps.sh && go build

RUN find . -type f ! -name 'main' -delete \
    rm -rf db \
    apk del git 

RUN ls

EXPOSE 8080

CMD [ "./main" ]