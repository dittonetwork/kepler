FROM golang:1.23.1-bullseye AS build-env

# Set working directory for the build
WORKDIR /app

# Add source files
COPY . .

RUN mkdir -p ~/.ignite/
RUN echo "{\"name\":\"gnrsrivjtjiymush\",\"doNotTrack\":true}" > ~/.ignite/anon_identity.json
#
# install ignite
RUN curl https://get.ignite.com/cli! | bash

# Make the binary
RUN make build


# Final image
FROM public.ecr.aws/docker/library/alpine:3.20

# Install deps
RUN apk add --no-cache tzdata ca-certificates libc6-compat
RUN wget https://truststore.pki.rds.amazonaws.com/global/global-bundle.pem -O /etc/ssl/certs/global-bundle.pem && update-ca-certificates

# Copy over binaries from the build-env
COPY --from=build-env /go/bin/keplerd /usr/bin/keplerd

EXPOSE 26656 26657 1317 9090 2345

# Run ethermintd by default
ENTRYPOINT ["keplerd"]