FROM golang:1.21.7-alpine3.19

WORKDIR /app

COPY . .

RUN apk add --no-cache \
    bash \
    git 

RUN go install github.com/cosmtrek/air@v1.49.0
# RUN wget -O- -nv https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s 

ENTRYPOINT [ "./entrypoint.sh" ]