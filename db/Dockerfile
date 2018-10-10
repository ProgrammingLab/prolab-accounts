FROM ruby:2.5.1-alpine3.7
LABEL maintainer="Ryota Egusa <gedorinku@yahoo.co.jp>"

WORKDIR /tmp
COPY Gemfile Gemfile.lock ./

RUN apk --no-cache add build-base mysql-dev mysql-client && \
    bundle install -j2 && \
    ridgepole --version

WORKDIR /db
ENTRYPOINT echo "（´・ω・｀）"
