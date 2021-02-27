FROM golang

WORKDIR /app

COPY . /app

RUN go build -o app ./calculator_server

EXPOSE 50051

ENTRYPOINT  /app/app
