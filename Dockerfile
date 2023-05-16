FROM golang AS builder
COPY . /go/src
WORKDIR /go/src
RUN git config --global --add safe.directory /go/src
RUN make

FROM alpine
#RUN [ ! -e /etc/nsswitch.conf ] &&
RUN echo 'hosts: files dns' > /etc/nsswitch.conf
COPY --from=builder /go/src/bin/container-web-tty /usr/bin
EXPOSE 8080
CMD [ "container-web-tty" ]
