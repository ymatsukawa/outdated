FROM ubuntu:20.04
RUN apt-get update
RUN apt-get install -y openssl git build-essential

RUN git clone https://github.com/ConorOG/xxd.git /usr/src/build/xxd
RUN cd /usr/src/build/xxd && make
RUN cp /usr/src/build/xxd/xxd /usr/local/bin/

WORKDIR /usr/src/bc