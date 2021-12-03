FROM golang:1-alpine

RUN apk add --no-cache git gcc musl-dev

# this gives us some nice colors when running our tests so its a bit
# easier to locate the failing and passing tests
RUN go get -u github.com/rakyll/gotest

RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
