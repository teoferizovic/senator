# start with base image
FROM mysql:8.0.23

# import data into container
# All scripts in docker-entrypoint-initdb.d/ are automatically executed during container startup

COPY ./provision/mysql-db/*.sql /docker-entrypoint-initdb.d/