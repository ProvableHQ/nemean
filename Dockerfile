FROM rust:1.57.0-slim as builder

RUN apt-get update -y && \
    apt-get install -y wget make libssl-dev pkg-config

RUN wget https://golang.org/dl/go1.17.3.linux-amd64.tar.gz

RUN rm -rf /usr/local/go && tar -C /usr/local -xzf go1.17.3.linux-amd64.tar.gz

COPY . .

ENV PATH=/usr/local/go/bin:${PATH}

RUN make build

FROM debian:sid-slim as final

COPY --from=builder nemean .
COPY --from=builder /usr/lib/libaleo.so /usr/lib/libaleo.so
COPY --from=builder /usr/include/aleo.h /usr/include/aleo.h

CMD ["/nemean"]
