FROM golang:1.13.7
WORKDIR /go/src/github.com/muquit/mailsend-go/
COPY . .
RUN make linux

FROM alpine:latest  
RUN apk --no-cache add ca-certificates
COPY --from=0 /go/src/github.com/muquit/mailsend-go/mailsend-go_linux /usr/local/bin/mailsend-go
ENTRYPOINT ["/usr/local/bin/mailsend-go"]
