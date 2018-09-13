FROM golang:1.10

WORKDIR /go/src/github.com/fiscaluno/dohko
COPY . .

RUN go get -u github.com/kardianos/govendor
RUN govendor sync
