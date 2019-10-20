FROM golang:1.11.4

WORKDIR /go/src/github.com/naoki-kishi/go-api-sample
COPY . .
ENV GO111MODULE=on

RUN go get github.com/pilu/fresh

RUN go build main.go
RUN chmod u+x main.go

EXPOSE 8081
