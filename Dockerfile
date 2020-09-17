FROM golang:1.15.2 as builder

WORKDIR /go/src/app
COPY . .

RUN make build

FROM debian:10
RUN apt-get update && \
    apt-get install -y \
        ca-certificates && \
    rm -rf /var/lib/apt/lists/*

COPY --from=builder /go/src/app/bin/miniloops-operator /miniloops

EXPOSE 9100 8080

CMD ["/miniloops"]