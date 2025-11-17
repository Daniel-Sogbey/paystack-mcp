FROM golang:1.25-alpine AS build

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

COPY .well-known /app/.well-known

WORKDIR /app/cmd/server

RUN CGO_ENABLED=0 GOOS=linux go build -o /paystack-mcp-server

FROM alpine:3.18

WORKDIR /src

COPY --from=build /app/.well-known /src/.well-known

COPY --from=build /paystack-mcp-server /src/paystack-mcp-server

EXPOSE 8080

CMD ["/src/paystack-mcp-server"]
