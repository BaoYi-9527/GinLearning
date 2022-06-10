FROM golang:1.17.7

ENV GOPROXY https://goproxy.io,direct
WORKDIR $GOPATH/src/Gin-Learning
COPY . $GOPATH/src/Gin-Learning
RUN go build .

EXPOSE 8000
ENTRYPOINT ["./GinLearning"]

