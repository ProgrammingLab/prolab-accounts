FROM ruby:2.5.1-alpine3.7 as builder

WORKDIR /tmp
COPY Gemfile Gemfile.lock ./

RUN apk add build-base mysql-dev
RUN bundle install -j2


FROM ruby:2.5.1-alpine3.7
LABEL maintainer="Ryota Egusa <gedorinku@yahoo.co.jp>"

COPY --from=builder /usr/local/bundle /usr/local/bundle
COPY --from=builder /usr/lib/libmysqlclient.so.18 /usr/lib
WORKDIR /db

ENTRYPOINT [ "ridgepole" ]
CMD [ "--version" ]
