FROM ruby:2.6-alpine3.9
LABEL maintainer="Ryota Egusa <gedorinku@yahoo.co.jp>"

WORKDIR /db
COPY Gemfile Gemfile.lock ./

RUN apk add build-base postgresql-dev
RUN bundle install

ENTRYPOINT [ "ridgepole" ]
CMD [ "--version" ]
