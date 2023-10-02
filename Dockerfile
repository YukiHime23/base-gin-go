## Developer
FROM golang:1.19-buster as development

WORKDIR /usr/src/base-gin-go

COPY . .
RUN go install github.com/cosmtrek/air@latest
RUN go mod download
CMD air

## Build
FROM golang:1.19-buster as build

WORKDIR /usr/src/base-gin-go

COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY . ./

RUN go build -o /base-gin-go

CMD ["/base-gin-go"]
