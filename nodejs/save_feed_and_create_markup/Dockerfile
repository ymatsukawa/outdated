FROM ubuntu:20.04
RUN mkdir -p /usr/src/app

RUN apt-get update
RUN apt-get install -y curl
# https://github.com/nodesource/distributions/blob/master/README.md#installation-instructions
RUN curl -fsSL https://deb.nodesource.com/setup_14.x | bash -
RUN apt-get install -y nodejs

RUN node -v
# https://classic.yarnpkg.com/en/docs/install#debian-stable
RUN curl -sS https://dl.yarnpkg.com/debian/pubkey.gpg | apt-key add -
RUN echo "deb https://dl.yarnpkg.com/debian/ stable main" | tee /etc/apt/sources.list.d/yarn.list
RUN apt-get update && apt-get install yarn
ADD . /usr/src/app

RUN apt-get install -y pandoc

WORKDIR /usr/src/app
