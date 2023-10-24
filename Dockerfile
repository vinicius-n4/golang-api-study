FROM golang:1.21.1

WORKDIR /golang-api-study

COPY . .
RUN go mod download

RUN CGO_ENABLE=0 GOOS=linux go build -o /main

EXPOSE 8000

CMD ["/main"]