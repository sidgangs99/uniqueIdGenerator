FROM golang:1.18-alpine

WORKDIR /usr/src/app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./

RUN go build -o /uniqueIdService

EXPOSE 8000

CMD [ "/uniqueIdService" ]