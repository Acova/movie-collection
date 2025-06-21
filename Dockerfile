FROM golang:latest

WORKDIR /app

COPY go.mod go.sum ./

RUN go get .

COPY *.go ./

RUN go build -o /movie-collection

EXPOSE 8080

CMD [ "/movie-collection" ]