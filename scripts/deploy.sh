#!/bin/sh
ssh -p [port] [user]@[remote_server] "source ~/.profile; ~/deploy-server.sh"
