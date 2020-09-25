FROM golang:1.14

WORKDIR saruman/src
COPY . .

ENV MYSQL_GORM_CONN_STRING=${MYSQL_CONN_STRING}
RUN go get -d -v ./...
RUN go install -v ./...
ENTRYPOINT [ "/go/bin/saruman" ]