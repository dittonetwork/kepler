FROM golang:1.23.2 AS build-env

WORKDIR /app

COPY cmd/explorer-graphql .

RUN CGO_ENABLED=0 GOOS=linux go build -o explorer-graphql server.go


FROM public.ecr.aws/docker/library/alpine:3.20

COPY --from=build-env /app/explorer-graphql /usr/bin/explorer-graphql

EXPOSE 8080

ENTRYPOINT ["explorer-graphql"]
