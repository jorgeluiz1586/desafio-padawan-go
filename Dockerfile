FROM golang:latest

RUN apt update && apt upgrade -y && \
    apt install git nano -y

COPY . /app

WORKDIR /app

RUN go get

CMD [ "go", "run", "main.go" ]