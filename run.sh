#!/bin/bash

pushd `dirname $0` > /dev/null
BASEDIR=`pwd`
popd > /dev/null


start() {
    nohup $BASEDIR/im_monitor  >/data/logs/im_monitor/im_monitor.log  2>&1 &
}

stop() {
    killall im_monitor
}

case "$1" in
    start)
        start
        ;;

    stop)
        stop
        ;;


    *)
        echo $"Usage: $0 {start|stop}"
        exit 2
esac
