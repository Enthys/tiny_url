FROM golang:1.21.5-alpine

WORKDIR /home/app

COPY . .

RUN go install github.com/cosmtrek/air@latest

EXPOSE 80

ENTRYPOINT ["air"]
