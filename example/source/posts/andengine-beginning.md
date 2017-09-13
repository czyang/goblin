{
    "title": "AndEngine初步",
    "author": "Chengzhi Yang",
    "createDate": "2010-12-16",
    "modifyDate": "2010-12-16",
    "permanent":"andengine-beginning"
}

# AndEngine初步

AndEngine 是一个Android平台下基于OpenGL ES的免费（LGPL协议）2D游戏引擎。他的作者是德国人*Nicolas Gramlich*

AndEngine 内置了Box2D物理引擎，你可以非常方便的在你的程序里面使用物理效果。

![alt text](../attachment/andengine/logo.png "Logo Title Text 1")

## AndEngine相关资料

AndEngine的官方网站：http://www.andengine.org/

他的源码保存在google code 上面你可以从这里下载：http://code.google.com/p/andengine/

AndEngine的Wiki     http://wiki.andengine.org 你想知道AndEngine可以做哪些游戏的话请看：http://wiki.andengine.org/List_of_Apps_and_Games

注：签出引擎代码的时候你可以在你的Eclipse下安装 Mercurial 插件http://www.vectrace.com/mercurialeclipse/（推荐）或者直接下载Mercurial用命令行签出http://mercurial.selenic.com/downloads/ 。

这里是几篇比较有用的AndEngine英文教程（这些教程来自AndEngine论坛  http://www.andengine.org/forums/tutorials/）：

Getting started with AndEngine  http://www.andengine.org/forums/tutorials/getting-started-with-andengine-t11.html

Eclipse, AndEngine and HelloWorld!!! http://www.andengine.org/forums/tutorials/eclipse-andengine-and-helloworld-t380.html

Mimminito’s Tutorial List http://www.andengine.org/forums/tutorials/mimminito-s-tutorial-list-t417.html

---
AndEngine功能强大并且易于使用，作者非常活跃。虽然AndEngine缺少文档，但是作者提供了非常多并且非常全面的示例程序，AndEngine Examples， http://code.google.com/p/andengineexamples/ 。里面把AndEngine的主要功能都做了演示，这些实例程序非常简洁，100多行的程序就能够实例很复杂的功能，虽然同样缺少文档。“好代码本身就是最好的文档。”(Steve McConnell，《代码大全》的作者) 。AndEngine的代码非常清晰。AndEngine缺少官方的手册和规范的文档，但是你可以从AndEngine 的 论坛http://www.andengine.org/forums/ 里面找到许多非常不错的英文教程，如果你遇到问题你可以先用论坛的搜索功能搜索，一般都能够发现别人之前的解决方案，如果你的问题没有得到解决的话，可以在论坛里面提问，论坛的成员和作者都非常热情并且活跃，能够在短时间内解决你棘手的问题。

在使用AndEngine前你需要具备一些Android平台的编程能力，Java的编程能力，这不是必须的，不过如果你不具备的话很可能你会遇到很多困难（你也可以同时学他们）。

下面的我将写如何使用AndEngine：

使用AndEngine写一个简单的程序，你首先得下载AndEngine的源码（http://code.google.com/p/andengine/）这个不是必须的，因为AndEngine引擎的jar 库你可以在 AndEngine Examples里面找到，你下载AndEngine后你可以自己做一些简单的修改，并且可以获取最新的AndEngine更新，建议下载。

下载AndEngine Examples  http://code.google.com/p/andengineexamples/ AndEngine Examples有apk版的，可以直接下载进手机里面运行。这一步对于刚刚接触AndEngine的人是必须的，你可以吧AndEngine Examples 看成是一个非常好的实例教程，AndEngine论坛里的某人曾经说过，他每天都要在手机里面运行 AndEngine Examples里面的所有例子两次以上以学习AndEngine ! 可见 AndEngine Examples 的重要性。

有了AndEngine Examples后，你就新建一个你自己的Android项目，比如HelloWrold，target platform 一般选1.6，建立一个你自己的activity 你可以从他的lib 目录里面找到 andengine.jar.


把andengine.jar 直接复制到你的项目的lib目录下，如果没有lib目录请自己创建一个，然后右键点击刚刚被复制过来的andengine.jar 然后点击Build Path -> Add To Build Path，这样就成功的把AndEngine添加到你的项目中了。

然后将你的HelloWorld.java 换成如下代码：(这里的代码出处是：http://www.andengine.org/forums/tutorials/eclipse-andengine-and-helloworld-t380.html)

```java
public class HelloWorld extends BaseGameActivity {
// ===========================================================
// Constants
// ===========================================================
private static final int CAMERA_WIDTH = 720;
private static final int CAMERA_HEIGHT = 480;
// ===========================================================
// Fields
// ===========================================================
private Camera mCamera;
// ===========================================================
// Constructors
// ===========================================================
// ===========================================================
// Getter & Setter
// ===========================================================
// ===========================================================
// Methods for/from SuperClass/Interfaces
// ===========================================================
@Override
public Engine onLoadEngine() {
  this.mCamera = new Camera(0, 0, CAMERA_WIDTH, CAMERA_HEIGHT);
  return new Engine(new EngineOptions(true, ScreenOrientation.LANDSCAPE,new   RatioResolutionPolicy(CAMERA_WIDTH, CAMERA_HEIGHT), this.mCamera));
}
@Override
public void onLoadResources() {
}
@Override
public Scene onLoadScene() {
  this.mEngine.registerUpdateHandler(new FPSLogger());
  final Scene scene = new Scene(1);
  scene.setBackground(new ColorBackground(0, 0, 0.8784f));
  return scene;
}
@Override
public void onLoadComplete() {
}
// ===========================================================
// Methods
// ===========================================================
// ===========================================================
// Inner and Anonymous Classes
// ===========================================================
}
```

添加完这段代码之后，别忘了在你项目的AndroidManifest.xml 中添加

```xml
<uses-permission android:name=“android.permission.WAKE_LOCK”/>
```

现在运行这个程序你就可以得到一个蓝色的屏幕，里面除了一个蓝色的背景什么都没有，下面你所要做的是打开Android Examples项目，一面看里面例子的效果，一面看效果的实现代码，相信你很快就能够入门！
