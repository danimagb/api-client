FROM golang:1.17-alpine AS base
ENV CGO_ENABLED=0
RUN apk update && apk add wget
RUN mkdir -p /home/user/app
WORKDIR /app
COPY . /app
COPY go.mod go.sum .


FROM base as unit-tests
COPY ./pkg ./pkg
RUN go mod download

CMD go test -cover ./pkg ./pkg/accounts ./pkg/core

FROM unit-tests as integration-tests
COPY ./tests/integration/ ./tests/integration/
ENTRYPOINT ["sh", "./integration-tests-entrypoint.sh"]
CMD go test -cover ./tests/integration
