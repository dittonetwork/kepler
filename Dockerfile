FROM gcr.io/distroless/base-debian12

WORKDIR /app

COPY main keplerd
COPY docs/static/openapi.yml docs/static/openapi.yml

CMD ["./keplerd", "start", "--home", "/app/config"]