#! /bin/sh
# /etc/init.d/sgoblin
#

USER="root"
GOPATH="/root/gohome"
GOROOT="/usr/local/go"
WORKDIR="/root/gohome/src/github.com/czyang/goblin"
NAME="goblin"

case "$1" in
  start)
    echo "Starting script sgoblin "

    cd "$WORKDIR"
    export GOROOT="$GOROOT"
    export GOPATH="$GOPATH"
    export PATH="${PATH}:${GOROOT}/bin:${GOPATH}/bin"

    nohup goblin &
    ;;
  stop)
    echo "Stopping script sgoblin"
    ;;
  *)
    echo "Usage: /etc/init.d/sgoblin {start|stop}"
    exit 1
    ;;
esac

exit 0
