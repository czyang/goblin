{
    "title": "2D Boy Framework 使用笔记",
    "author": "Chengzhi Yang",
    "createDate": "2010-10-28",
    "modifyDate": "2010-10-28",
    "permanent":"2dboy-framework-note-1"
}

# 2D Boy Framework 使用笔记

如果你不了解2D Boy Framework 你需要先阅读：[]初探 2D Boy Framework]()

如果你想快速上手的话，可以阅读框架自带的两个demo和Game.h中的注释。2D Boy Framework 目前是没有官方文档的，但是代码结构非常清晰，如果有面向对象游戏引擎使用经验的人，只需要看一下这个框架中自带的两个Demo就可以非常快速的上手。我写这篇文章只是为了做一下记录我遇到的问题，或者框架代码中注释不清楚的地方。

程序需要定义一个类继承Game 类，实现Game类的一些必要接口，然后给Environment 的一个实例（该实例为单件）调用。

```cpp
Boy::Environment::instance()->init( Game *game, int screenWidth, int screenHeight,
   bool fullscreen, const char *windowTitle, const UString &persFile,
   unsigned char *persFileKey);
```

Game类中定义的类都在 WinEnvironment.cpp中定义， 类WinEnvironment继承于类Environment。

继承于Game类的主类一般需要实现下面的几个接口，如果你的程序想实现监听键盘操作，同时还需要继承 public Boy::KeyboardListener，并且实现相关的类，游戏手柄同理，后文再叙述键盘相关的东西。先描述一下常用的几个函数，Game还有一些函数会在游戏运行的过程中被调用，可以在Game.h里面查看，里面有详细的注释。

```cpp
virtual void init ();
```
非必需实现，被游戏的框架自动调用，你可以做一些你主类的初始化工作。

```cpp
virtual void load ();
```

非必需实现，在init()之后被调用，在这里你需要做一些健壮的初始化工作（例如，大量资源的载入，大量图片的设置）。这个调用是运行在一个独立的线程中的，使用的时候必须谨慎（例如：如果你的程序需要画一张PNG的话，你在load()中载入你的图片，然后在draw()里面渲染，如果你没有在draw()中加一个判断该图片是否已经成功被载入的标记的话，极有可能会出错）。


```cpp
virtual void loadComplete();
```

这个函数是在主线程中运行的，但是他只能够在load()调用完成后运行！你可以做一些在load()被成功调用之后做的事情。

```cpp
virtual void update(float dt);
```

必须实现。这个函数被框架自动调用，游戏的主循环每循环一次就调用一次，他在draw()之前被调用。其中的参数dt是，这一次调用update()的时间到上一次调用update()的时间间隔，dt的单位是秒！


```cpp
virtual void draw(Boy::Graphics *g);
```

必须实现。被框架自动调用，游戏的主循环每循环一次就被调用一次。这个函数就是你渲染图像的地方。


Boy将Dx9的图形API封装在类Graphics中，具体实现是在类WinGraphics ，类WinGraphics 继承了类Graphics的接口，类WinGraphics在使用的时候已经被定义在WinEnvironment中*mGraphics， 并且已经在WinEnvironment::init中做好了初始化图形接口的工作，你的程序在需要绘图的时候只需要直接通过Environment 的单件实例获取指针：

```cpp
Boy::Graphics *g = Boy::Environment::instance()->getGraphics();
```
