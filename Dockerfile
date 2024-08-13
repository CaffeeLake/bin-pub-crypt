# workspace
FROM golang:1.23 AS workspace

COPY . /bin-pub-crypt

WORKDIR /bin-pub-crypt

RUN go mod download \
  && CGO_ENABLED=0 go build -buildmode pie -o /bin-pub-crypt/bin-pub-crypt

# production
FROM gcr.io/distroless/base:debug AS production

RUN ["/busybox/sh", "-c", "ln -s /busybox/sh /bin/sh"]
RUN ["/busybox/sh", "-c", "ln -s /bin/env /usr/bin/env"]

COPY --from=workspace /bin-pub-crypt/bin-pub-crypt /bin/bin-pub-crypt

ENTRYPOINT ["/bin/sh", "-c", "tail -f /dev/null"]
