FROM golang

WORKDIR /go/src/github.com/RepGarage/refrigerator/back/go/main

ADD . /go/src/github.com/RepGarage/refrigerator/back/go/main

RUN ./install-deps.sh && go build

EXPOSE 8080

CMD [ "./main" ]