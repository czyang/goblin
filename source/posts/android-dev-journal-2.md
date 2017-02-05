{
    "title": "Android开发资料记录(二)",
    "author": "Chengzhi Yang",
    "createDate": "2011-02-25",
    "modifyDate": "2011-02-25",
    "permanent":"android-dev-journal-2"
}

# Android开发资料记录(二)

如果感兴趣你可以阅读：[Android开发资料记录(一)](./android-dev-journal-1.html)

---
Android Canvas 的抗锯齿(Anti-Alias)设置
http://www.javaeye.com/topic/794505
法一：直接设置Paint   paint.setAntiAlias(true);
法二：不能设置Paint的函数可以用：
canvas.setDrawFilter(new PaintFlagsDrawFilter(, Paint.ANTI_ALIAS_FLAG|Paint.FILTER_BITMAP_FLAG));

原文中给出了两种方法的对比，方法二的效果是最好的。
这两种方法都可以用在文字的渲染上。
我在我自己的程序里面实验了下第一种方法。我的图片是一个64×64的小方块，方块一直在移动，并且有一定的旋转，用了第一种方法后锯齿效果更明显。可能是图片的原因。

Paint mPaint = new Paint(Paint.ANTI_ALIAS_FLAG);
这个方法也是可以开启 Paint的抗锯齿调用起来比较方便。

---

一个Android的详尽教程网站
http://www.helloandroid.com/
Android的教程网站，里面有很多有用的教学文章，涵盖Android编程几乎所有方面。
---

学习Android Canvas编程的非常好的学习示例
Lunar lander 的代码，学习Android Canvas 游戏编程的一个很好的示例。(需番羽墙)
http://developer.android.com/resources/samples/LunarLander/index.html
http://www.anddev.org/2d_tutorial-t3120.html  有人把这个游戏代码简化了，做了一个非常简单的游戏框架。非常不错。

**注意:**这两个框架有个严重的BUG，就是在运行这个框架做的游戏的时候，如果有电话打进来或者你点击率home键，在从新恢复游戏的时候，游戏会崩溃，网上的解决办法都不能够完美解决。

---

还有个更好的框架：看Chris pruett 的简易框架，这个项目里面包括了 简单的opengl es  和 Canvas的框架

http://code.google.com/p/apps-for-android/source/browse/#svn%2Ftrunk%2FSpriteMethodTest
---
Google I/O 2010 – All Android Talk Videos

Google I/O 2010 的全部演讲视频，有简要介绍。(youtube视频)
http://www.anddev.org/announces/google-i-o-2010-all-android-talk-videos-t14798.html

---

让你的Android应用在运行的时候屏幕不变暗。
默认情况下，Android应用会在系统设定的指定时间内没有点击事件后就会把屏幕变暗，然后到一定时间就关闭屏幕。下面的代码可以让你的应用开启的时候屏幕永远处于激活状态。对于好多应用和游戏都非常有用。
代码在：http://www.androidsnippets.org/snippets/53/
不要忘记在AndroidMainfest.xml 里面获取授权：
