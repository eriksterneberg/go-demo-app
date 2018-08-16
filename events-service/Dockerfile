# Multi-stage build file that discards the image used for building binary.
# Result is a smaller image.

FROM golang:1.10.3 as goimage
RUN mkdir -p /src/
WORKDIR /src/
COPY src/. .
RUN go get -d -v
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

FROM alpine:latest
COPY --from=goimage /src/app ./
