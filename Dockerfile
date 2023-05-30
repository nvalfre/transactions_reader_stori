FROM golang:1.19

WORKDIR /app/transactions_summary

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o transactions_summary

EXPOSE 8080

CMD ["./transactions_summary"]
