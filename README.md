# go-rdb-sandbox

```sh
# Start-up postgres with Docker
$ docker-compose up -d
# Connect postgres
$ docker-compose exec postgres psql -U grs -d grs_db
# Connect mysql
$ docker-compose exec mysql mysql -u grs --password=grs -D grs_db
```
