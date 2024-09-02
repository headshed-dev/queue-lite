#!/usr/bin/env bash

LOG_DIR=$(pwd)/logs
LOG_FILE=${LOG_DIR}/$(date +%Y%m%d).log
DATE_STR=$(date +%Y%m%d%H%M%S)
mkdir -p $LOG_DIR

echo "${DATE_STR} publsh started" >> $LOG_FILE

sleep 5

echo "${DATE_STR} publsh finished" >> $LOG_FILE