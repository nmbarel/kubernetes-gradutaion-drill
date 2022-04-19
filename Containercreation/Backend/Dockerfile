FROM golang:1.18 as buildstage

WORKDIR /go/src

COPY main.go .

RUN go mod init Backend &&  \
    CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o Backend .

FROM alpine:latest  

WORKDIR /tmp

COPY --from=buildstage /go/src/Backend .

RUN chmod +x Backend

ENTRYPOINT ["./Backend"]  
