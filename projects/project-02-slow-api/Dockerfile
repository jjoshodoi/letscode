FROM golang:1.20-alpine AS build
RUN apk add --no-cache gcc musl-dev
WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=1 go build -o /out/slowapi ./cmd/slowapi

FROM alpine:3.18
RUN adduser -D app
WORKDIR /app
COPY --from=build /out/slowapi /usr/local/bin/slowapi
RUN mkdir -p /app/data && chown -R app:app /app
USER app
ENV DB_PATH=/app/data/products.db
EXPOSE 8080
CMD ["/usr/local/bin/slowapi"]
