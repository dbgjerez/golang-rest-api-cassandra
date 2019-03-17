FROM golang:1.12

ENV GO111MODULE=on

WORKDIR /go/src/app
COPY . .

EXPOSE 8000

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["go-todo-rest-api-cassandra"]