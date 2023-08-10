FROM golang:alpine AS build-env

WORKDIR /go/src
COPY go.sum go.mod ./
RUN go mod download
COPY . /go/src/nostra-api-product
RUN cd /go/src/nostra-api-product && go build .

FROM alpine
WORKDIR /app
COPY --from=build-env /go/src/nostra-api-product/nostra-api-product /app
# COPY .env /app

ENTRYPOINT [ "./nostra-api-product" ]