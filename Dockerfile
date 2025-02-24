FROM alpine:3.21

RUN apk --no-cache add ca-certificates

WORKDIR /root

COPY keplerd .

EXPOSE 1317 26657

CMD ["./keplerd", "start", "--home", "/root/app_config"]