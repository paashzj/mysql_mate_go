#!/bin/bash

export REMOTE_MODE=false

mkdir -p $MYSQL_HOME/logs
nohup $MYSQL_HOME/mate/mysql_mate >>$MYSQL_HOME/logs/mysql_mate.stdout.log 2>>$MYSQL_HOME/mysql_mate.stderr.log

