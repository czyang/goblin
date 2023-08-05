# ![ICON](./goblin.png "ICON") Goblin: Static Blog Site Generator 

[![Build Status](https://travis-ci.org/gin-gonic/gin.svg)](https://travis-ci.org/czyang/goblin)
[![Go Report Card](https://goreportcard.com/badge/github.com/czyang/goblin)](https://goreportcard.com/report/github.com/czyang/goblin)

Showcase: https://codingmelody.com

All my blog posts managed by Github as a repo, and I can CRUD the articls inside the Github repo, [a script](https://github.com/czyang/goblin-den/blob/main/fire-ball.sh) will sync data at every 5m: https://github.com/czyang/goblin-den


Goblin is a static blog site generator developed using Go. It takes your Markdown posts and transforms them into a complete, static blog site inclusive of HTML/CSS assets and a sitemap, enhancing your SEO capabilities.

## Primary Features

1. Markdown-to-HTML Conversion: Converts your Markdown blog posts into HTML format.
2. Archive Page Creation: Generates an archive page compiling all your posts.
3. Static Site Generation: Produces static HTML pages for each of your blog posts.
4. Sitemap Creation: Automatically generates a sitemap of your blog to optimize SEO.
5. Ping Search Engines: Notifies search engines of updates to your sitemap.


## Core Components
**Posts**: Goblin processes blog posts written in Markdown format and converts them into HTML pages. Each post contains essential metadata such as title, tags, creation date, modification date, categories, permalink, and author information.

**Pages**: In addition to posts, Goblin also manages other pages on your site. It can convert markdown files into static HTML pages.

**Config**: The app leverages a configuration file to set parameters like navigation links, social media links, and host details.

**Usage**:
Execute the Goblin application with the following command line arguments:

* **posts**: The directory path to your blog posts written in Markdown format.
* **template**: The directory path to your HTML templates for the blog layout.
* **output**: The output directory where your static blog site will be generated.
* **config**: The directory path to your configuration file.

``` bash
go run . -posts=/path/to/posts -template=/path/to/templates -output=/path/to/output -config=/path/to/config.json
```

After the execution, you can find your fully generated blog site in the specified output directory.

Note: Make sure all your blog posts are in Markdown format with the appropriate metadata.

_Remember, the paths to posts, templates, output and configuration must be supplied as command-line arguments. Goblin does not provide default paths or handle errors if arguments are not provided._

## Instructions

### Install Go if needed
```bash
sudo apt update
sudo apt install golang-go
```

### Install Goblin
```bash
git clone https://github.com/czyang/goblin.git

cd goblin

go build -o goblin .

# [optional]
sudo mv goblin /usr/local/bin/
```

### Update Goblin
```bash
cd goblin

git pull

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

## Arch
![Goblin drawio](https://github.com/czyang/goblin/assets/830725/3beceae9-7c5e-413e-b269-fac4442bf095)


### License
Golbin is distributed under the [MIT License](./LICENSE.txt)
