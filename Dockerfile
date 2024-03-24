FROM node:20-alpine AS web-build

WORKDIR /build
COPY ./web/package.json .
RUN npm install
COPY ./web .
RUN npm run build

FROM golang:1.22.1-alpine AS build

WORKDIR /build
# TODO: exclude web files using --exclude when available in stable syntax
# https://docs.docker.com/reference/dockerfile/#copy---exclude
COPY . .
COPY --from=web-build /build/dist ./web/dist
RUN go mod download
RUN go build -o undershorts ./cmd/undershorts/main.go

FROM alpine:3 AS final

WORKDIR /app
COPY --from=build /build/undershorts .

CMD [ "/app/undershorts" ]
