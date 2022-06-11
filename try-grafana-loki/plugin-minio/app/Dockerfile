FROM golang:1.17-alpine3.13 as builder

ENV CGO_ENABLED 0
WORKDIR /app

COPY . .

RUN go build -o app main.go


FROM alpine:3.13.6 as production

WORKDIR /src
COPY --from=builder app .

EXPOSE 7000

CMD ["/src/app"]
