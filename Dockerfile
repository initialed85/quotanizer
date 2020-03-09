FROM golang:1.13

WORKDIR /srv/

COPY cmd /srv/cmd
COPY pkg /srv/pkg
COPY test /srv/test
COPY vendor /srv/vendor
COPY go.mod /srv/go.mod
COPY go.sum /srv/go.sum
COPY _test.sh /srv/_test.sh

RUN go build -o /srv/quotanizer ./cmd/main.go

ENTRYPOINT ["/srv/quotanizer"]
