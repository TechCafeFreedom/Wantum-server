#!/usr/bin/env bash

mysql -u root -proot -e "source /docker-entrypoint-initdb.d/ddl.sql"
mysql -u root -proot wantum < "/docker-entrypoint-initdb.d/dml.sql"