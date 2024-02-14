FROM golang:1.22 as builder

RUN mkdir /app
ADD . /app
WORKDIR /app
ENV GOOS=linux \
    CGO_ENABLED=0
RUN go build -o app cmd/server/main.go

FROM scratch as production
COPY --from=builder /app .
CMD ["./app"]
