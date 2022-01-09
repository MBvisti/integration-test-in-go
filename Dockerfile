FROM golang:1-alpine AS base
WORKDIR /app

RUN apk add --no-cache git gcc musl-dev

FROM base AS dev

RUN go get github.com/cosmtrek/air
RUN go get github.com/vektra/mockery/v2/.../

COPY .air.toml .

CMD ["air"]

