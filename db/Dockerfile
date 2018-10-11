FROM ruby:2.5.1-alpine3.7 as builder

WORKDIR /tmp
COPY Gemfile Gemfile.lock ./

RUN apk add build-base mysql-dev
RUN bundle install -j2


FROM ruby:2.5.1-alpine3.7
LABEL maintainer="Ryota Egusa <gedorinku@yahoo.co.jp>"

RUN apk --no-cache add mysql-client
COPY --from=builder /usr/local/bundle /usr/local/bundle
WORKDIR /db

ENTRYPOINT [ "ridgepole" ]
CMD [ "--version" ]
