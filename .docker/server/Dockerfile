
FROM golang:1.20.11 as dev

RUN go install github.com/cosmtrek/air@latest

WORKDIR /var/www/html

# COPY --from=golang /go /go
# COPY --from=golang /usr/local/go /usr/local/go

COPY ./server .

# COPY ./.env .

RUN go mod tidy

CMD ["air", "-c", ".air.toml"]
