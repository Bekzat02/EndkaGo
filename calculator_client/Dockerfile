FROM golang

WORKDIR /app

COPY . /app

RUN go build -o app /calculator_client/client.go

EXPOSE 8080

ENTRYPOINT  /app/app
