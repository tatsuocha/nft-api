# Dockerfile for Go application
FROM golang:latest

WORKDIR /go/src/app
COPY . .

RUN go get -d -v ./...

# main.go をビルドして実行可能ファイルを生成
RUN go build -o main .

# 実行可能ファイルに実行権限を付与する
RUN chmod +x ./main

# 生成した実行可能ファイルを実行する
CMD ["./main"]