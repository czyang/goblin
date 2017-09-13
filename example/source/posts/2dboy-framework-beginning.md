{
    "title": "初探 2D Boy Framework",
    "author": "Chengzhi Yang",
    "createDate": "2010-10-13",
    "modifyDate": "2010-10-13",
    "permanent":"2dboy-framework-beginning"
}

# 初探 2D Boy Framework

2D Boy Framework 是2D Boy studio（http://2dboy.com/）开发的一个2D跨平台游戏开发框架，粘粘世界(World of Goo) 使用了这个框架，这是一个开源框架你可以在 2D Boy studio 的论坛（ http://2dboy.com/forum/index.php#2）的 2D Boy Framework这个栏目里面下载到他的源码，这个框架在使用上没有任何的许可上的限制。

开发语言：C++

图形API：Dx9，DXUT

声音引擎:：irrKlang  (http://www.ambiera.com/irrklang/docu/index.html)

它也使用了SDL，不过只用了SDL的一些关于系统方面的函数，如设置窗口标题，鼠标指针，时间，多线程编程。这样的封装可以带来很多好处，一方面提高了游戏开发的效率，只需要调用其他库的库函数就能实现很多复杂的功能；另一方面大大增强了代码的健壮性和可读性，因为框架使用的这些库都是经过长时间的使用和维护的。

---

这个框架项目文件自带的是VS2008的，如果需要在VS2005里面编译项目的两个DEMO，可以简单的吧 .sln 文件的 第2，3行改成

Microsoft Visual Studio Solution File, Format Version 9.00

# Visual Studio 2005

注意：.sln文件的第一行是一个空行。

然后把其他的 .vcproj 中的

ProjectType="Visual C++"

Version="9.00"

改成

ProjectType="Visual C++"

Version="8.00"

然后可以尝试打开.sln，如果打不开就按提示修改.vcproj文件。其中的DXUTCore_2008.vcproj，可以在打开.sln后先删掉该项目，然后添加把DXUTCore_2005.vcproj 进解决方案中，这个文件在..\libs\dxsdk\200806\Samples\C++\DXUT\Core

---

之后按照项目中的readme 就可以很容易的设置好，然后就可以运行两个Demo。

这个框架的C++代码非常清晰。



框架的一些特性：

* 封装过的游戏运行过程，让你免于去写初始化窗口和消息循环之类的代码；

* 封装过的资源管理器，支持PNG读取，读取PNG时只需要在一个xml文件 里面定义图片的ID和路径，就可以通过封装过的资源管理器读取和调用图片资源（Framework 通过TinyXML库来读取解析XML文件）；

* Boy Framework 本身的框架只实现了对DirectX 的支持，不过图形接口封装得很好，可以方便的写出用OpenGL来渲染的图形接口；

* 因为粘粘世界在WII平台上面发行过，Boy Framework有对WII平台的支持；（框架里面有不上关于WII平台的宏定义，我对WII平台开发不了解，所以我猜框架对WII平台是有一定的支持的）

* 封装过的键盘Listener 和 鼠标Listener ，并且支持GamePad  和 Conroller；

* 支持ASE(Advanced Encryption Standard) 加密， MD5加密；

* 内部实现一些轻量级的数学类和函数，如二维向量，向量旋转，点线距离，颜色插值等；

* 框架还是实现了几个简单的边界检测类；

* 框架对声音支持，可以非常容易的读取和播放声音，类似于读取图片的过程；

* 框架支持英文字体，可以方便读取英文字体（一张图片），然后方便的显示出来英文， 字体文件使用的是 Popcap Framework 中的格式，Popcap Framework 中附带了制作字体的工具（后面介绍Popcap framework）；



几个缺点

* 没有对精灵动画的直接实现；

* 没有GUI；

* 没有实用的碰撞检测和物理效果的实现；（对于这一点可以自己实现自己的碰撞检测，如AABB，OBB； 物理上来说可以使用 Box 2D（http://box2d.org/） ，或者 ODE(http://www.ode.org/)， 这两个库能够实现非常非常牛逼的物理效果，并且都实现了碰撞检测，粘粘世界使用的就是ODE）

* 没有对中文显示的支持；

* 没有非常完整高效的数学库；（网上找找可以找到一大堆优化过的数学库）

* 这个库没有被作者维护；


从这几个缺点中可以看出来2D Boy Framework  只是一个 原型Framework ，并非是一个游戏引擎，不过他提供的这个框架非常值得学习，而且可以在其基础上面开发出自己的游戏引擎。

我非常非常喜欢这个代码，非常感谢2D Boy 的慷慨分享 :-)


Boy Framework的基础框架的设计参考了不少 Popcap framework  http://sourceforge.net/projects/popcapframework/  （Popcap framework 当然是 Popcap 公司过去发布的一个开源项目），这个框架的功能就非常非常强大了，比Boy Framework 的功能复杂得多。但是Popcap framework发布得非常早，图形接口是 D3D 7… ，不过可以把他修改成D3D 9的来用。
