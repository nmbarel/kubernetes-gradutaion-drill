FROM golang:1.18 as buildstage

WORKDIR /go/src

COPY main.go .

RUN go mod init Frontend && \
    CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o Frontend .

FROM alpine:latest  

WORKDIR /tmp

COPY --from=buildstage /go/src/Frontend ./

RUN chmod +x Frontend

ENTRYPOINT ["./Frontend"] 
