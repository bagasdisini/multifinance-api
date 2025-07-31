# Step #1
FROM golang:1.24 AS firststage
LABEL description="Multifinance API"
LABEL maintainer="Bagas <mbagas221@gmail.com>"
WORKDIR /build/
COPY . /build
ENV CGO_ENABLED=0
RUN go get
RUN go build -o multifinance-api

# Step #2
FROM alpine:latest
WORKDIR /app/
RUN apk update
RUN apk upgrade
RUN apk add ca-certificates && update-ca-certificates
RUN apk add --no-cache tzdata gcompat
ENV TZ=Asia/Jakarta
COPY --from=firststage /build/multifinance-api .
CMD ["./multifinance-api"]
