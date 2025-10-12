# scraping by python

## requirements

* linux OS
* python3 and pip3
  * python v3.9.6 is recommended.

## before run

```
$ cd /path/to/dir_README_exist
$ pip3 install -r src/requirements.txt
$ cp src/.env.example src/.env
```

edit `src/.env`

```
# PROD, DEV or TEST
SCRAPING_ENV=DEV
```

edit `src/setttings.(prod|dev|test).json`

scheme(prod, dev or test) is mapped to `.env`'s PROD, DEV or TEST

<details>
  <summary>settings.*.json edit profile</summary>

[featured json is `src/settings.test.json`](./src/settings.test.json)

**download**

* how delays each `targets` crawl.
* minimum unit of search is `one name, one start_url`

**feed**

* setting is exported from [FEEDS item](https://docs.scrapy.org/en/latest/topics/feed-exports.html#feeds)

**targets**

* `targets.name`
  * unique name of crawl
* `targets.start_url`
  * where to start crawl
* `targets.search.desc`
  * key name of xpath result
  * follows section in `run`, `each_p` and `title` is respond to this
* `targets.search.xpath`
  * value of xpath result
  * also `run` section, found element is written in json value 
* `targets.search.next_xpath_href` (optional)
  * if crawler needs to search page, set xpath of href
* `targets.search.max_crawl_page` (required when `next_xpath_href `is set)
  * set how many crawl page
  * when crawler detect page is over, exit safe automatically

</details>

## run

```
$ python3 src/main.py
```

Result path is written in src/settings.(prod|dev|test).json's `feed`.`/path/to/output_dir/file.json`.

you can change the path. Follows are example output from example.com

```
[
{"each_p": "<p>This domain is for use in illustrative examples in documents. You may use this\n    domain in literature without prior coordination or asking for permission.</p>"},
{"each_p": "<p><a href=\"https://www.iana.org/domains/example\">More information...</a></p>"},
{"title": "<title>Example Domain</title>"}
]
```

## testing

### requirements

* docker >= 20.10.7
* docker-compose >= 1.16.1

```
$ docker-compose build

terminal A
$ docker-compose up

terminal B
edit src/.env 's SCRAPING_ENV as TEST

$ docker-compose run app python3 main.py
```

crawler crawling test web app of djnago at `test` directory

# other

about contribution:

This repository's purpose is personal use. So I don't plan to improvement by contribution.

If you want to customize, please fork. Thanks.
