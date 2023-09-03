FROM golang:1.20 AS build

WORKDIR /app

COPY . .

RUN go mod download

# Build the application binary
RUN CGO_ENABLED=0 go build -o ./app cmd/api/main.go

FROM debian:stable-slim

WORKDIR /app

COPY --from=build /app/pkg/config/ ./config
COPY --from=build /app/app .

CMD ["./app"]
