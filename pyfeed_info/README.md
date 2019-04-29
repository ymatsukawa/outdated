# pyfeed_info

pyfeed_info enables to gather rss feeds and send them to external service.

And already posted feeds are not thrown again.

## Precondition

### Infras

* OS
  * Linux Base System
  * Mac OS
* MiddleWare
  * python 3.6.4
  * MongoDB 3.6.2

### Settings

Edit `/path/to/pyfeed_info/project/settings.py`.

Setting content and how-to is written in the source.

## Run operation

```bash
$ export PYTHONPATH=/path/to/pyfeed_info
$ cd /path/to/pyfeed_info
$ mongod
$ ./venv/bin/python ./project/main.py
```

## LICENSE

MIT
