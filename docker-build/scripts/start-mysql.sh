#!/bin/bash

mysqld --daemonize --user=root

mysql -u root < $MYSQL_HOME/mate/sql/init.sql