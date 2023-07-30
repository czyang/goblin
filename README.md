# ![ICON](./goblin.png "ICON") Goblin 

[![Build Status](https://travis-ci.org/gin-gonic/gin.svg)](https://travis-ci.org/czyang/goblin)
[![Go Report Card](https://goreportcard.com/badge/github.com/czyang/goblin)](https://goreportcard.com/report/github.com/czyang/goblin)

**A Golang static blog generator.**

Showcase: https://codingmelody.com
All my blog posts managed by Github as a repo, and I can CRUD the articls inside the Github repo, [a script](https://github.com/czyang/goblin-den/blob/main/fire-ball.sh) will sync data at every 5m: https://github.com/czyang/goblin-den

Goblin is a **Minimal** static pages generator Golang app which generate pages from markdown files with little extra JSON header.

### Install Go if needed
```bash
sudo apt update
sudo apt install golang-go
```

### Install Goblin
```bash
git clone https://github.com/czyang/goblin.git

cd goblin.git

go build -o goblin .

# [optional]
sudo mv goblin /usr/local/bin/
```

### Setup a static folder
```bash
sudo mkdir -p /var/www/yourwebsite.com/public_html

sudo chown -R $USER:$USER /var/www/yourwebsite.com/public_html
sudo chmod -R 755 /var/www/yourwebsite.com

# Create a folder for holding static files
sudo mkdir -p /var/www/codingmelody.com/public_html
# Modify the permission
sudo chown -R $USER:$USER /var/www/codingmelody.com/public_html
sudo chmod -R 755 /var/www/codingmelody.com
```

### Run

Example:
```bash
goblin -posts=./input_folder/posts -template=./input_folder/tmpl  -output=/var/www/codingmelody.com/public_html -config=./config.json
```

Run Goblin for from Go folder:
```bash
go run . -posts=./input_folder/posts -template=./input_folder/tmpl  -output=./output_folder -config=./config.json
```

### Arch
![Goblin drawio](https://github.com/czyang/goblin/assets/830725/3beceae9-7c5e-413e-b269-fac4442bf095)


#### License
Golbin is distributed under the [MIT License](./LICENSE.txt)
