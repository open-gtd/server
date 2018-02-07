FROM debian:9

WORKDIR /go/src/app
COPY ./build/open-gtd-server .

RUN mkdir -p /etc/open-gtd
COPY docker/server.conf.json /etc/open-gtd/

CMD ["./open-gtd-server"]