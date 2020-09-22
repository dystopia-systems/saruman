FROM golang:1.14


WORKDIR saruman/src
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

EXPOSE 4000
CMD ["saruman"]
VOLUME /var/saruman/