# ![alt text](./source/attachment/goblin_icon.png "Logo Title Text 1") Goblin 

![travis-ci-status-icon](https://travis-ci.org/czyang/goblin.svg?branch=master "Build Status")

** A Golang static blog generator and holder. **

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
You sould keep the folders stract like the example.
```sh
--- static          # Generated files.
--- tmpl            # Your sites template.
--- config.json     # Config file.
--- source          # Your files.
   \--- attachment  # Images / Downloadable files etc.
   \--- pages       # Your custom pages. like profile page.
   \--- posts       # Your articles write in markdown.
```

Run Goblin like this:
```sh
goblin /path/to/the/template/
```

#### License
Golbin is distributed under the [MIT License](./LICENSE.txt)
