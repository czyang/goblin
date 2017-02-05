{
    "title": "在MacOS下搭建ejoy2D编译环境",
    "author": "Chengzhi Yang",
    "createDate": "2014-01-23",
    "modifyDate": "2014-01-23",
    "permanent":"macos-setup-ejoy2d-environment"
}

# 在MacOSX下搭建ejoy2D编译环境

Ejoy2D是Ejoy公司的一个主要针对iOS的开源跨平台2D游戏引擎。项目地址：(https://github.com/cloudwu/ejoy2d)

在mingw下面搭建的时候除了注意把依赖的库拷到对应的位置还是很容易编译的。

但是昨天我发现在MacOSX 10.9(mavericks)下面好像不太搞，按照说明用默认路径安装各个依赖库，然后直接make是编译不过的。下面记录下我在MacOSX下面编译的大概过程，这个肯定不是最佳解决方案，更像一个奇葩方案，不过还是能够解决问题。

针对MacOSX官网文档说：

For Mac OS ,

  * Install Lua 5.2.3
  * Install glew 1.9
  * Install freetype 2
  * Install XQuartz
  * make or make macosx
  * ./ej2d examples/ex01.lua to test

对于新版的XQuartz是自带freetype 2的，所以可以只安装Lua和glew。

如果你以默认路径安装完成这些依赖库后你运行make可能会遇到下面几种错误：

<!--more-->

**错误1

** ld: library not found for -lglew 或者

ld: library not found for -llua

这个些错误应该是库的搜索路径不对造成的，makefile里 LDFLAGS += -L/usr/X11R6/lib -lGLEW -lGL -lX11 -lfreetype -llua -lm -Wl,-E -ldl

所以一种解决办法是吧 glew 和 lua的库文件直接复制到：/usr/X11R6/lib的文件夹里面（这里可能是 /usr/X11/lib）。这样如果这几个依赖库的文件没有缺失，那么编译就不会再报找不到库的路径的错误。

**错误2**

/usr/X11R6/include/ft2build.h:56:10: fatal error: &#8216;freetype/config/ftheader.h&#8217; file not found

这个错误可以通过在makefile的CFLAGS的那行中加入： -I/usr/X11R6/include/freetype2

macosx : CFLAGS += -I/usr/X11R6/include -I/usr/include -I/usr/local/include <span style="color: #ff0000;">-I/usr/X11R6/include/freetype2</span> $(shell freetype-config &#8211;cflags) -D __MACOSX

**错误3

** ld: unknown option: -E

这个错误可以把makefile中的-WL,-E 改成 -Wl,-no\_compact\_unwind

macosx : LDFLAGS += -L/usr/X11R6/lib -lGLEW -lGL -lX11 -lfreetype -llua -lm <span style="color: #ff0000;">-Wl,-no_compact_unwind</span> -ldl

**Mac OS X下安装freetype2(这个步骤非必须)**

freetype 2安装的时候可能需要先安装Autoconf, Automake & Libtool，这三个库，命令如下(来源(http://www.jattcode.com/installing-autoconf-automake-libtool-on-mac-osx-mountain-lion/)：

```sh
curl -OL http://ftpmirror.gnu.org/autoconf/autoconf-2.69.tar.gz
tar -xzf autoconf-2.69.tar.gz
cd autoconf-2.69
./configure && make && sudo make install

curl -OL http://ftpmirror.gnu.org/automake/automake-1.14.tar.gz
tar -xzf automake-1.14.tar.gz
cd automake-1.14
./configure && make && sudo make install

curl -OL http://ftpmirror.gnu.org/libtool/libtool-2.4.2.tar.gz
tar -xzf libtool-2.4.2.tar.gz
cd libtool-2.4.2
./configure && make && sudo make install
```

之后就按照docs/INSTALL.UNIX的步骤安装

```sh
sudo sh autogen.sh
./configure
make
sudo make install
```
