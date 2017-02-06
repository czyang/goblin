#!/bin/sh
ssh root@your_server_ip << EOF
ls -a -l
export GOPATH=$HOME/gohome
export PATH=$PATH:/usr/local/go/bin:$GOPATH/bin
"~/deploy-server.sh"
EOF