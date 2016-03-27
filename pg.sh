#!/bin/bash
docker kill bookhouse-db
docker rm bookhouse-db
docker run \
  --name bookhouse-db \
  -e POSTGRES_PASSWORD=bookhousepass \
  -e POSTGRES_USER=bookhouseuser \
  -e POSTGRES_DB=bookhousedb \
  -p 5432:5432 \
  -d postgres:9.5.1

# psql postgres://bookhouseuser:bookhousepass@dockerhost:5432/bookhousedb

# select * from books where host like '%amazon.com';
# select * from books where path like '/Da-Vinci-Code-Dan-Brown/dp/0307474275%'


# function getPathFromUrl(url) {
#   return url.split(/[?#]/)[0];
# }
