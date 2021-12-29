# start with base image
FROM postgres:10.1

# import data into container
# All scripts in docker-entrypoint-initdb.d/ are automatically executed during container startup
COPY ./database/init.sql /docker-entrypoint-initdb.d/