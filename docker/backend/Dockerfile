FROM golang:1.20.2
RUN apt-get update && apt-get install -y build-essential make

# vscodeの拡張機能のためにインストール
RUN go install golang.org/x/tools/gopls@latest
RUN go install github.com/ramya-rao-a/go-outline@latest
RUN go install github.com/cosmtrek/air@latest

# go のパッケージを install
WORKDIR /backend
# COPY ./backend/go.mod ./backend/go.sum ./
# RUN go mod download && go mod verify
