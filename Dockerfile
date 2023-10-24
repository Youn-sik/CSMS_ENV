FROM golang:1.20.6

WORKDIR /data/CSMS_ENV

COPY .	.

RUN go mod download

RUN go build -o CSMS_ENV

EXPOSE 9080

CMD ["./CSMS_ENV"]
