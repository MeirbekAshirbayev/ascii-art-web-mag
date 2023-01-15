# syntax=docker/dockerfile:1

FROM golang:1.17-alpine

WORKDIR /app

COPY . .

RUN go build -o web/main/*

EXPOSE 4000

CMD [ "/ascii-art-web" ]