FROM golang:1.18-buster
WORKDIR /ascii-art-web-dockerize
COPY ./ ./
RUN go build -o /docker
EXPOSE 8080
CMD ["/docker"]
LABEL version="1.0"
LABEL description="That's our Ascii-art-web \
We hope you enjoyed!"