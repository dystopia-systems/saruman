FROM golang:1.14

WORKDIR saruman/src
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

EXPOSE 3000
ENTRYPOINT "saruman"
VOLUME /var/saruman/