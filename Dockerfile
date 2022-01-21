FROM golang:1-alpine AS base
WORKDIR /app

RUN apk add --no-cache git gcc musl-dev

FROM base AS dev

RUN go get github.com/cosmtrek/air
RUN go get github.com/vektra/mockery/v2/.../
RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

COPY .air.toml .

CMD ["air"]

