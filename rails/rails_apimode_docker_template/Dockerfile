FROM ruby:3.1.1-slim
RUN apt-get update -qq && apt-get install -y build-essential libpq-dev curl git

RUN mkdir -p /usr/src/app
WORKDIR /usr/src/app

ADD Gemfile /usr/src/app/Gemfile
ADD Gemfile.lock /usr/src/app/Gemfile.lock
RUN bundle install

ADD . /usr/src/app
