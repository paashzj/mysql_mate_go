#!/bin/bash

mysqld --daemonize

mysql -u root < $MYSQL_HOME/mate/sql/init.sql