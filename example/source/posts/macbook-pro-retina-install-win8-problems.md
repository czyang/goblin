{
    "title": "Macbook Pro Retina 安装windows 8 / windows 8.1 问题记录",
    "author": "Chengzhi Yang",
    "createDate": "2014-02-20",
    "modifyDate": "2014-02-20",
    "permanent":"macbook-pro-retina-install-win8-problems"
}

# Macbook Pro Retina 安装windows 8 / windows 8.1 问题记录

这里是我在 13′ 的 Macbook Pro Retina（ME866） 安装windows 8 / windows 8.1 遇到的一些问题。我安装时用Mac OSX 自带的Boot Camp 安装的，windows 8的系统版本我选的是 windows 8.1 pro。
1. Macbook Pro Retina 安装windows 8 后不能上网。
默认的windows 8.1 pro貌似没有现成的Macbook Pro的无线网卡。解决方法：1.可以直接靠一个驱动精灵；2，下载最新的Boot Camp Support Software 里面包含了大概1G的MBP可能用到的驱动，然后通过在设备管理器里面更新驱动的方式更新。

2. Macbook Pro Retina 安装windows 8 后系统没有声音。
安装后我发现MBP的扬声器没有声音，耳机孔红色的灯饰亮着的，解决方法：直接下载驱动精灵，省去不少烦恼。

3. 修改IE 11 的搜索引擎成 Google SSL。
IE右上角小齿轮->新窗口的最下面“查找更多搜索提供程序”->添加谷歌作为搜索引擎->win+R regedit->[HKEY_CURRENT_USER\Software\Microsoft\Internet Explorer\SearchScopes\ 在这个字段下，会有对应的几个子字段，分别对应各个搜索引擎的属性，找到google对应的，修改字段”URL”改成”https://www.google.com/search?q={searchTerms}&#8221; 注销再登陆就可以了，其中的SuggestionsURL我还不知道怎么填，没有找到合适的SSL链接URL。
4.诡异的Chrome Retina。
我直接装了个Chrome，发现Chrome不是默认就能支持Retina的。打开Chrome 地址栏输入：chrome://flags/ -> 找到HiDPI 这个项->开启Enable->然后你的Chrome就会出现异常->关掉Chrome，打开Chrome安装目录在Chrome图标上右键（或者在Chrome的快捷方式上），属性->然后 “兼容性”->勾选”高DPI设置时禁用显示缩放”->再打开Chrome就应该相对正常了。


总结，在Macbook Pro Retina上面使用现在的windows，确实很蛋疼，多数软件都还不支持或者支持得不好，包括一些很常用的软件，包括QQ，Evernote，Chrome… 可以观望之后出的windows 9会不会好起来，不是特别爱玩游戏或者必须用windows的话不建议安装。
