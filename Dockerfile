FROM debian:slim as copy
RUN echo '' > /app/config/data/priv_validator_state.json && echo '' > /app/config/config/addrbook.json

FROM gcr.io/distroless/base-debian12

WORKDIR /app

COPY main kepler
COPY docs/static/openapi.yml docs/static/openapi.yml
COPY --from=copy /app/config/data/priv_validator_state.json /app/config/data/priv_validator_state.json
COPY --from=copy /app/config/config/addrbook.json /app/config/config/addrbook.json

CMD ["/app/kepler", "start", "--home", "/app/config"]
