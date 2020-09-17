FROM golang:1.15 as builder

WORKDIR /go/src/app
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...
RUN GOOS=linux CGO_ENABLED=0 go build main.go

FROM scratch
COPY --from=builder /go/src/app/main /go/bin/main

EXPOSE 9000

CMD ["/go/bin/main"]
