FROM golang:1.22.3-alpine3.17 as builder

RUN apk update && apk upgrade && apk --update add git make bash build-base

WORKDIR /app
COPY . .
RUN make build

FROM alpine:3.17

RUN apk update && apk upgrade && apk --update --no-cache add tzdata

WORKDIR /app
ARG SERVER_PORT
EXPOSE ${SERVER_PORT}
COPY --from=builder /app/api_server /app/

CMD /app/api_server