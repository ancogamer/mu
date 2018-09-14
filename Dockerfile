FROM golang:1.8

WORKDIR /go/src/github.com/fiscaluno/mu
COPY . .

RUN go get -u github.com/kardianos/govendor
RUN govendor sync
