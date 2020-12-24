FROM golang

WORKDIR /go/splitter
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...
RUN go build -o splitter *.go

CMD ["./splitter"]