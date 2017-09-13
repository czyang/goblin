#!/bin/sh
# Stop running goblin
kill -SIGKILL `pgrep goblin`

# Update goblin
cd /root/blog/goblin_blog
git pull origin master

# Run goblin
nohup goblin /root/blog/goblin_blog &