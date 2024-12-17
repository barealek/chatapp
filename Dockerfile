FROM golang:alpine AS build
WORKDIR /app

COPY go.* .

RUN go mod download -x

COPY . .
RUN go build -o bin/server cmd/server/main.go


FROM node:alpine AS node
WORKDIR /web
COPY web/package*.json .
RUN npm i
COPY web .
RUN npm run build

FROM alpine:latest

WORKDIR /app
COPY --from=build /app/bin /app/bin
COPY --from=node /web/build /app/dist

EXPOSE 3000
CMD [ "bin/server" ]
