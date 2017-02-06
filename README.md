# ![alt text](./tmpl/assets/images/header_icon.png "Logo Title Text 1") Goblin

**My personal blog within whole code.** 

Review it on: https://codingmelody.com

Goblin is a simplicity **static** pages generator app within minimal features which generator pages from markdown files with little extra JSON header.

#### Installation
```sh
# To install
go get github.com/czyang/goblin

# To update 
go get -u github.com/czyang/goblin

```

#### Docker Instructions *(optional)*
```sh
docker build -t goblin .

docker run -d -p 8001:8001 --name goblin <ref>
```

#### License
Golbin is distributed under the [MIT License](./LICENSE.txt)
