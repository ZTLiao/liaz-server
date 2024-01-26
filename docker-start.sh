#!/usr/bin/env bash
ROOT_DIR=/root/golang/liaz-server
APPLICATION_NAME=$1
SERVER_PORT=$2
PROFILES_ACTIVE=prod

if [ -z "$APPLICATION_NAME" ]
then
  echo 'applicationName is null!'
  exit
fi

WORK_DIR=/data/golang/$APPLICATION_NAME

mkdir -p $WORK_DIR

cd $ROOT_DIR && docker build -f $ROOT_DIR/Dockerfile -t $APPLICATION_NAME --build-arg PROFILES_ACTIVE=$PROFILES_ACTIVE --build-arg APPLICATION_NAME=$APPLICATION_NAME --build-arg SERVER_PORT=$SERVER_PORT $WORK_DIR

docker run --rm --net=host \
  -p $SERVER_PORT:$SERVER_PORT \
  --name $APPLICATION_NAME -d -v /etc/timezone:/etc/timezone:ro -v /etc/localtime:/etc/localtime:ro -v $ROOT_DIR:$WORK_DIR/ -v $WORK_DIR/logs:$WORK_DIR/logs \
  $APPLICATION_NAME \
  make $APPLICATION_NAME  || exit
