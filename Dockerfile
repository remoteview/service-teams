FROM golang:1.11.0-alpine3.8 as builder

ADD . /go/src/github.com/remoteview/service-teams/

WORKDIR /go/src/github.com/remoteview/service-teams

RUN set -ex && \
  CGO_ENABLED=0 go build -tags netgo -o service-teams -v -a -ldflags '-extldflags "-static"' && \
  mv ./service-teams /usr/bin/service-teams

# Temporarely using go in order to run migrations on shell.
FROM golang:1.11.0-alpine3.8
COPY --from=builder /usr/bin/service-teams /usr/local/bin/service-teams

WORKDIR /service-teams
COPY --from=builder /go/src/github.com/remoteview/service-teams/VERSION.txt /service-teams/VERSION.txt
COPY --from=builder /go/src/github.com/remoteview/service-teams/database.yml /service-teams/database.yml
COPY --from=builder /go/src/github.com/remoteview/service-teams/migrations /service-teams/migrations

EXPOSE 3001

ENTRYPOINT [ "service-teams" ]
