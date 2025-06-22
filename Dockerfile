FROM golang:latest

WORKDIR /app

COPY . .

RUN go get .

RUN go build -o /movie-collection

CMD [ "/movie-collection" ]