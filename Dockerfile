FROM debian:bookworm-slim

RUN apt update && \
    apt install -y curl

WORKDIR /tmp

RUN curl https://dl.google.com/go/go1.21.1.linux-amd64.tar.gz -o go.tar.gz && \
    tar -C /usr/local -xzf go.tar.gz && \
    rm -rf go.tar.gz

ENV GOPATH /usr/local/go
ENV PATH $PATH:$GOPATH/bin

WORKDIR /golang-api-study

COPY . .

CMD ["go", "run", "main.go"]
