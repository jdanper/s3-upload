FROM golang:alpine AS build-env
RUN apk --no-cache add build-base git bzr mercurial gcc make
ADD . /src
RUN cd /src && make build

FROM alpine
WORKDIR /app
COPY --from=build-env /src/audit-sink /app/
ENTRYPOINT ./audit-sink