FROM scratch

ENV GOPROXY https://goproxy.cn,direct
ENV BASE $GOPATH/src/github.com/siconghe/blog
WORKDIR $GOPATH/src/github.com/siconghe/blog
COPY . $GOPATH/src/github.com/siconghe/blog

EXPOSE 5000
CMD ["./blog"]