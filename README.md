# ![alt text](./source/attachment/goblin_icon.png "Logo Title Text 1") Goblin 

![travis-ci-status-icon](https://travis-ci.org/czyang/goblin.svg?branch=master "Build Status")
[![Go Report Card](https://goreportcard.com/badge/github.com/czyang/goblin)](https://goreportcard.com/report/github.com/czyang/goblin)

** A Golang static blog generator. **

Showcase: https://codingmelody.com

Goblin is a **Minimal** static pages generator Golang app which generate pages from markdown files with little extra JSON header.

*If you want a FULL featured and well maintained static site generator you should check [Hugo](https://github.com/spf13/hugo). (At least it has a fancy logo :P)*

#### Installation
```sh
# To install
go get -u github.com/czyang/goblin
```


#### Usage
Copy the example folder as your site template.
Modify config.json. 
You sould keep the folders structure like the example show as below:
```sh
.
├── config.json     # Config file.
├── source          # Your files.
│   ├── attachment  # Images / Downloadable files etc.
│   ├── pages       # Your custom pages. like profile page.
│   └── posts       # Your articles write in markdown.
├── static          # Generated files.
└── tmpl            # Your sites template.
```

Run Goblin like this:
```sh
goblin -input=./input_folder -output=./output_folder
```

#### License
Golbin is distributed under the [MIT License](./LICENSE.txt)
