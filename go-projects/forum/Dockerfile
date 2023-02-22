FROM golang:1.19.2-bullseye

LABEL name="forum"
LABEL description="Forum using sqlite and go"
LABEL url="https://github.com/01-edu/public/tree/master/subjects/forum"

WORKDIR /app

copy src .

RUN go build .

EXPOSE 8080

CMD ["/app/forum"]