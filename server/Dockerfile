FROM golang:1.15.7 as builder
RUN mkdir /build
ADD * /build/
WORKDIR /build
ENV GO111MODULE=on

RUN CGO_ENABLED=0 GOOS=linux go build -o server
RUN ls 

# generate clean, final image for end users
FROM alpine:3.11.3
COPY --from=builder /build/server .
ENTRYPOINT [ "./server" ]