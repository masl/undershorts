FROM golang:1.16.6-alpine AS build

WORKDIR /build
COPY . .
RUN go mod download
RUN go build -o undershorts ./cmd/undershorts/main.go

FROM alpine:3 AS final

WORKDIR /app
COPY --from=build /build/undershorts .
CMD [ "/app/undershorts" ]