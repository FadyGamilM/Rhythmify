FROM postgres:14
ENV POSTGRES_DB auth
COPY init.sql /docker-entrypoint-initdb.d/
