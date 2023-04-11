# build stage
FROM golang:1.19-alpine3.16 AS builder
#maintainer info
LABEL maintainer="fazil muhammed <fazilkp2000@gmail.com>"
#installing git
RUN apk update && apk add --no-cache git

WORKDIR /Job-Portal

COPY . .

RUN apk add --no-cache make

RUN make deps
RUN go mod vendor
RUN make build



# Run stage
FROM alpine:3.16

WORKDIR /Job-Portal
COPY go.mod .
COPY go.sum .
COPY --from=builder /Job-Portal/build/bin/api .


CMD [ "/Job-Portal/api"] 