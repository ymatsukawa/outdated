# README

* Ruby version

2.6.0

* System dependencies

* Configuration

```
cp .env.sample .env
{YOUR EDITOR} .env
```

* Database creation

```
docker-compose run web rake db:create
```

* Database initialization

```
docker-compose run web rake db:migrate
```

* Initializaiton and run

```
docker-compose build
docker-compose up -d
```

* Stop process

```
docker-compose down
```

* How to run the test suite

* Deployment instructions
