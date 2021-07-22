# save feed and create markup

**This repository is an experiment use, and does not assure quality.**

**Do not use in production environment.**

## prepare

```
$ docker-compose build
$ docker-compose up
```

## create feed and markup

`CATEGORY_NAME` is written `FEED_CONST.feeds.name` in `./app/common/feed_category.js`

* get feed and save to DB

```
curl -X POST -H "Content-Type: application/json" -d '{"category":"CATEGORY_NAME"}' http://localhost:3000/feed
```

* create category's data to markup

```
curl -X POST -H "Content-Type: application/json" -d '{"category":"CATEGORY_NAME"}' http://localhost:3000/feed/markup
```

markup file, html is located at `./created_files`

## jenkins

* access to http://localhost:8080
* login as `root / mypass`
* install suggested plugins
* set job and script for periodic execution
* default running scripts are in `./jenkins/home/run_scripts/**/*.bash`
