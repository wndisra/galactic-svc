FROM golang:alpine AS build
WORKDIR /app
COPY . .
RUN go build -o galactic-svc ./cmd/server

FROM alpine:3.11.3
WORKDIR /app
RUN cd /app
COPY --from=build /app/.env .
COPY --from=build /app/galactic-svc /app/galactic-svc

ENTRYPOINT [ "/app/galactic-svc" ]