FROM golang:1.19.3-alpine3.15 as builder
COPY go.mod go.sum /go/src/productsAPI/
WORKDIR /go/src/productsAPI
RUN go mod tidy
COPY . /go/src/productsAPI
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o build/productsAPI productsAPI

FROM alpine
RUN apk add --no-cache ca-certificates && update-ca-certificates
COPY --from=builder /go/src/productsAPI/build/productsAPI /usr/bin/productsAPI
EXPOSE 8080 8080
ENTRYPOINT ["/usr/bin/productsAPI"]