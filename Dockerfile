FROM golang:alpine AS build
WORKDIR /app

COPY go.* .

RUN go mod download -x

COPY . .
RUN go build -o bin/server cmd/server/main.go


FROM alpine:latest

COPY --from=build /app/bin /bin

CMD [ "/bin/server" ]
