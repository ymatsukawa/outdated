FROM python:3.9.6

ENV PYTHONBUFFERED=1
RUN mkdir /usr/src/test
WORKDIR /usr/src/test

COPY . /usr/src/test
RUN pip3 install -r requirements.txt
