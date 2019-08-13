FROM golang:1.12

ENV GO111MODULE=on
ENV CASSANDRA_URL=cassandra
ENV CASSANDRA_USERNAME=user
ENV CASSANDRA_PASSWORD=user

WORKDIR /go/src/app
COPY . .

EXPOSE 8000

RUN go get -d -v ./...
RUN go install -v ./...
RUN rm -rf /go/src

RUN groupadd -r golang && useradd --no-log-init -r -g golang golang
USER golang

CMD ["go-todo-rest-api-cassandra"]