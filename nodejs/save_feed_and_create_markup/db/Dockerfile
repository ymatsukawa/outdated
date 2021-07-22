FROM nouchka/sqlite3:latest

RUN mkdir -p /usr/src/app/db
WORKDIR /usr/src/app/db

ADD . /usr/src/app/db
RUN sh prepare_sqlite.sh
RUN pwd && ls

ADD . /usr/src/app/db
