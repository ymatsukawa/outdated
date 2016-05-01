# Customized Ver. README

(showing respect to @m-yamashita -san, @kakusuke -san)

# What's this?

Customized [Lodge](https://github.com/lodge/lodge) by @ymatsukawa compliance with MIT.

# What is (will be) featured?

* is featured
  * Notice email will be sent if article or comment is posted.
  * (email send process is delayed job and daemonized)
* will be featured
  * no plans

# how to start

## Given environment

* CentOS 6.6
* Ruby 2.2.1
* bundler 1.9.9


## 1. setup config/database.yml

```
$ cp config/database.yml.example config/database.yml
# enable which to use sqlite3 or mysql
```

## 2. setup .env and {Rails.root}/config/environments/production.rb

```
$ cp .env.example .env
# edit .env

# LODGE_DOMAIN expected 'domain and port'
## ex. LODGE_DOMAIN = localhost:3000
# MAIL_SENDER expected 'noreply notifier'
## ex. MAIL_SENDER = noreply_lodge@example.com

# if plan to use external MTA(ex. gmail)
# check notice > email header's "from" spec
```

## 3. bundle gems

```
$ {Rails.root}/bin/bundle install --path vendor/bundle
```

## 4. migrate db

### 4-1. restore SQLite3 

```
$ cp /path/to/lodge_production {Rails.root}/
```

### 4-2. restore MySQL

* no tested

### 4-3. migarte db

```
$ RAILS_ENV=production {Rails.root}/bin/rake db:migrate
```

## 5. execute delay job daemon process

```
$ RAILS_ENV=production {RAILS.root}/bin/delayed_job start
# if want to stop process
# $ RAILS_ENV=production {RAILS.root}/bin/delayed_job stop
```

## 6. create emoji

```
$ bundle exec rake emoji
```

## 7. launch HTTP server(webrick)

```
$ {Rails.root}/bin/rails s -e production
```

# notice

## email header's "from" spec

* if .env's DELIVER_METHOD is 'smtp'
  * sent email header's "from" will be set as DELIVER_METHOD > SMTP_USER
  * even MAIL_SENDER is set
