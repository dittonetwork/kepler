FROM alpine:3.16

RUN apk --no-cache add ca-certificates

WORKDIR /root

COPY appd .

EXPOSE 1317 26657

CMD ["./appd", "start", "--home", "/root/app_config"]