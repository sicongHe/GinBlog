FROM golang:latest

ENV GOPROXY https://goproxy.cn,direct
ENV BASE $GOPATH/src/github.com/siconghe/blog
WORKDIR $GOPATH/src/github.com/siconghe/blog
COPY . $GOPATH/src/github.com/siconghe/blog
RUN go build .

EXPOSE 5000
ENTRYPOINT ["./blog"]