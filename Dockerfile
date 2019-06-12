FROM golang:alpine

WORKDIR $GOPATH/src/github.com/LGUG2Z/microfest
COPY . .

RUN apk update && apk add --no-cache git

ENV GO111MODULE=on
RUN CGO_ENABLED=0 GOOS=linux go install ./cmd/microfest-server

ENTRYPOINT ["/go/bin/microfest-server"]
