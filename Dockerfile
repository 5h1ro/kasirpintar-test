FROM golang:1.19-alpine as builder

RUN apk update && apk add --no-cache git build-base

WORKDIR /app
COPY . .

RUN go install -mod=mod github.com/githubnemo/CompileDaemon
RUN go get github.com/swaggo/swag/cmd/swag

ADD https://github.com/ufoscout/docker-compose-wait/releases/download/2.7.3/wait /wait
RUN chmod +x /wait

CMD CompileDaemon --build="go build main.go" --command="./main" --color