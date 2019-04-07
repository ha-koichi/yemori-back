FROM golang:1.11.7

WORKDIR /app
ADD . /app

CMD ["go", "run", "main.go"]

ENV GO111MODULE=on

RUN go get github.com/pilu/fresh
CMD ["fresh"]