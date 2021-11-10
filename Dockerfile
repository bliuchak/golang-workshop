FROM golang:1.17.2-alpine3.13 as builder
WORKDIR /go/src/app
ARG VERSION="n/a"

RUN go get github.com/cespare/reflex

COPY . .

RUN CGO_ENABLED=0 \
  GOOS=linux \
  GOARCH=amd64 \
  go build \
    -o /api \
    -ldflags "-X 'main.Version=$VERSION'" \
    -v \
    ./cmd/api

FROM scratch
COPY --from=builder /api /
USER 9000
ENTRYPOINT [ "/api" ]
