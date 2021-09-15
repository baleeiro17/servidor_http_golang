FROM golang:1.14.4-stretch

WORKDIR /workspace

RUN git clone https://github.com/baleeiro17/servidor_http_golang.git \
    && cd servidor_http_golang \
    && go mod download


# Move to the binary path
WORKDIR /workspace/servidor_http_golang/cmd

RUN go build -o server

CMD [ "./server" ]