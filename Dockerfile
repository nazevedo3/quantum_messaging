FROM golang:1.16 as construct

COPY go.mod go.sum /go/src/github.com/nazevedo3/quantum_messaging/
WORKDIR /go/src/github.com/nazevedo3/quantum_messaging
RUN go mod download
COPY . /go/src/github.com/nazevedo3/quantum_messaging
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o build/quantum_messaging github.com/nazevedo3/quantum_messaging


FROM alpine

RUN apk add --no-cache ca-certificates && update-ca-certificates
COPY --from=construct /go/src/github.com/nazevdo3/quantum_messaging/build/quantum_messaging /usr/bin/quantum_messaging

EXPOSE 8080 8080

ENTRYPOINT ["/usr/bin/quantum_messaging"]
