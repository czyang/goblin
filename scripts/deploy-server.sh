#!/bin/sh
# Stop running goblin
kill -SIGKILL `pgrep goblin`

# Update goblin
go get -u github.com/czyang/goblin

# Run goblin
cd ~/gohome/src/github.com/czyang/goblin
nohup goblin &