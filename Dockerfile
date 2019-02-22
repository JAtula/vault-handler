FROM golang:1.11.5-stretch AS build

LABEL Maintainer="Juhani Atula <juhani.atula@polarsquad.com"

WORKDIR /opt/build

COPY . .

ENV GOOS=linux
ENV GOARCH=amd64
ENV CGO_ENABLED=0

RUN go build .

FROM golang:1.11.5-alpine

COPY --from=build /opt/build/vault-handler /

ENTRYPOINT [ "/vault-handler" ]