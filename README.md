# ![ICON](./goblin.png "ICON") Goblin 

![travis-ci-status-icon](https://travis-ci.org/czyang/goblin.svg?branch=master "Build Status")
[![Go Report Card](https://goreportcard.com/badge/github.com/czyang/goblin)](https://goreportcard.com/report/github.com/czyang/goblin)

** A Golang static blog generator. **

Showcase: https://codingmelody.com/blog/posts/index.html

Goblin is a **Minimal** static pages generator Golang app which generate pages from markdown files with little extra JSON header.

#### Installation
```sh
# To install
go get -u github.com/czyang/goblin
```

#### Usage

```sh
.
├── goblin          # goblin program
├── config.js       # config for program
├── posts           # Your files.
│   ├── attachment  # Images / Downloadable files etc.
│   └── *.md        # Posts
└── tmpl            # Your sites template.
```
`

Run Goblin like this:
```sh
go run . -posts=./input_folder/posts -template=./input_folder/tmpl  -output=./output_folder
```

#### License
Golbin is distributed under the [MIT License](./LICENSE.txt)
