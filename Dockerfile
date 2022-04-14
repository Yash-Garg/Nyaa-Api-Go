FROM golang:1.18.1-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN apk add git

RUN go mod download

COPY . .

RUN go build -o nyaa-api .

EXPOSE 8080

CMD [ "/app/nyaa-api" ]
