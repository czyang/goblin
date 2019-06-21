{ "title": "Go Porject settings in VS Code", "author": "Chengzhi Yang", "createDate": "2019-06-17", "modifyDate": "2019-06-17", "permanent":"vscode-go-dev" }


# Go Porject settings in VS Code

For now VS code is a great tool for develop Go project without much efforts to 
config. For most simple Go project you only need to install the Go plugin to 
turn your VS Code into Go IDE. I have noted some problems I have encountered 
when I setup Go in VS Code. 

## For the project using [Go 1.11 Modules](https://github.com/golang/go/wiki/Modules) 
It's a bit tricky to setup. Because we need set environment virable
`GO111MODULE=on` and build flag: `-mod=vendor`
```javascript
// .vscode/launch.json for Go debug
{
    "version": "0.2.0",
    "configurations": [
        
        {
            "name": "Launch Package",
            "type": "go",
            "request": "launch",
            "mode": "debug",
            "program": "${workspaceFolder}",
            "env": {"GO111MODULE": "on"},
            "buildFlags": "-mod=vendor",
            "args": []
        }
    ]
}
```
Then you can run the Go project by clicking the debug button in vscode.
You can also change the `.vscode/settings.json` and add the `go.buildFlags` and 
`go.toolsEnvVars` . 

For the `go test`, you can set `go.testEnvVars` and `go.testFlags`. But it does 
not work for the test for the project which forked from the original project. 
For this situation please check below article.


## Disable 'format on save'

VS Code has enabled go format on save default. For most case, disable the go 
format on save is not recommended. But if you work on the legacy project or 
some project not put linter and format into workflow in the beginning. You may 
want to disable this feature.

It seems vscode-go or vscode have the bug with current version 
*`Version 1.35.1 (1.35.1)` go plugin version: `0.11.0`*
VS code can't disable "formatOnSave". And here is a alternative solution:

Open Shortcuts editor and change `saveWithoutFormatting` to 
`cmd+s` and change `save` to `cmd+k s`:

```json
{
  "key": "cmd+s",
  "command": "workbench.action.files.saveWithoutFormatting"
}
```

```json
{
  "key": "cmd+k s",
  "command": "workbench.action.files.save"
}
```

And we should using below solution to achive this effect. But it's not working 
for current vscode.

The solution is to add this property to the `.vscode/settings.json`
*May not work for `Version 1.35.1 (1.35.1)` go plugin version: `0.11.0`*
```json
    "[go]": {
        "editor.formatOnSave": false
    }
```

