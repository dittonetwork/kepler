FROM gcr.io/distroless/base-debian12

WORKDIR /app

COPY main kepler
COPY docs/static/openapi.yml docs/static/openapi.yml

CMD ["/app/kepler", "start", "--home", "/app/config"]