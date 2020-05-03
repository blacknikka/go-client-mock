
# init influx
docker-compose up -d
docker-compose exec influx influx -execute "CREATE USER myuser WITH PASSWORD 'secret' WITH ALL PRIVILEGES" && \
docker-compose exec influx influx -execute "CREATE DATABASE telegraf"
docker-compose down
docker-compose up -d
