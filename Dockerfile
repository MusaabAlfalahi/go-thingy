FROM golang:1.24-alpine AS build_base

RUN apk add --no-cache git

WORKDIR /tmp/go-thingy

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -buildvcs=false -o ./out/go-thingy github.com/MusaabAlfalahi/go-thingy

FROM alpine:3.21

RUN apk add ca-certificates

COPY --from=build_base /tmp/go-thingy/out/go-thingy /app/go-thingy

EXPOSE 8080

ENTRYPOINT [ "/app/go-thingy" ]
