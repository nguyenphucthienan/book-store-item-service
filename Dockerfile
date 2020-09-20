FROM golang:1.15.2

ENV ES_URL=localhost:9200
ENV LOG_LEVEL=info
ENV REPO_URL=github.com/nguyenphucthienan/book-store-item-service

ENV GOPATH=/app
ENV APP_PATH=$GOPATH/src/$REPO_URL

ENV WORKPATH=$APP_PATH
COPY . $WORKPATH
WORKDIR $WORKPATH

RUN go build -o item-service .

EXPOSE 8081

CMD ["./item-service"]
