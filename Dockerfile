FROM golang:1.24.2-bookworm

WORKDIR /app

COPY main kepler
COPY docs/static/openapi.yml docs/static/openapi.yml
RUN mkdir -p /app/config/data/ && mkdir /app/config/config && echo '' > /app/config/data/priv_validator_state.json && echo '' > /app/config/config/addrbook.json

CMD ["/app/kepler", "start", "--home", "/app/config"]
