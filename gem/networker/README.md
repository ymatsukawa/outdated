# Networker

Network test tool.

## Installation

```shell
$ rake build
$ gem install pkg/networker-x.x.x.gem
```

## Usage

In `irb`

```
require 'networker'
Networker::Http.req(:file, '/path/to/request_file.json')

# ex.)
Networker::Http.req(:file, './spec/FILES/http_get.json')
=>
{:status=>200,
 :headers=>
  {"content-type"=>"text/html;charset=utf-8",
   "content-length"=>"197",
   .
   .
   },
 :body=>
  "<!DOCTYPE html>\n<html lang=\"en\">...</body>\n</html>\n"}
```

example file format is in [spec/FILES](./spec/FILES)

* `params` equals `query params` or `body of post, put or delete`

## Testing

```shell
$ rake test:server

# on another console
$ rspec spec
```

When source code is changed on server running, please restart.

Hot-reload is not supported.

## License

[MIT License](https://opensource.org/licenses/MIT).
