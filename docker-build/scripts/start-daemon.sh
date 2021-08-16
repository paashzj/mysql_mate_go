#!/bin/bash

export REMOTE_MODE=false

nohup $MYSQL_HOME/mate/mysql_mate >$MYSQL_HOME/mysql_mate.log 2>$MYSQL_HOME/mysql_mate_error.log