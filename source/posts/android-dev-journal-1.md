{
    "title": "Android开发资料记录(一)",
    "author": "Chengzhi Yang",
    "createDate": "2011-01-13",
    "modifyDate": "2011-01-13",
    "permanent":"android-dev-journal-1"
}

# Android开发资料记录(一)

如果感兴趣你可以阅读：[Android开发资料记录(二)](./android-dev-journal-2.html)

---
### Android XML Parser Performance

http://www.developer.com/ws/article.php/10927_3824221_2/Android-XML-Parser-Performance.htm
DOM Pull SAX 方法在Android下的对比。（文中结论：SAX在Android下优于其余两种方法）
下面是一篇关于如何使用sax 从xml文件中载入游戏关卡文件的教程。这个AndEngine论坛里面的文章，但是跟AndEngine关系不大，它适用于所有用sax从xml文件中载入数据。
http://www.andengine.org/forums/tutorials/xml-parsing-within-andengine-t372.html

---

### OpenGL ES Tutorial for Android – Part I – Setting up the view

Per-Erik Bergman 关于Android下OpenGL ES 的系列教程，图文并茂。
http://blog.jayway.com/2009/12/03/opengl-es-tutorial-for-android-part-i/

Why a lot of android game  has own GLSurfaceView.
在很多开源的Android游戏里面你都能找到游戏自己的GLSurfaceView.java，而且跟原生的那个差不多，这个问题困扰了我很久，直到在AndEngine自己的GLSurfaceView.java里面我发现了如下的注释才稍微明白了一点：
// This class is a replacement for the original GLSurfaceView, due to issue:
// http://code.google.com/p/android/issues/detail?id=2828
// Reason: Two sequential Activities using a GLSurfaceView leads to a deadlock in the GLThread!
详细的原因我不太清楚，毕竟没有实践过，但是大概是因为某种原因原生的GLSurfaceView在同一时间只能有一个实例存在，多线程的情况下使用原生的GLSurfaceView可能会出现死锁。这个问题很致命，很多Android游戏运行的时候都会把主逻辑和渲染分开，这样可以提高渲染速度。Replica Island的作者在下面这文章中有详细描述。

---

### Rendering With Two Threads

http://replicaisland.blogspot.com/2009/10/rendering-with-two-threads.html  （需番羽墙）
作者将他游戏的渲染过程独立在一个单独的线程中执行以后，游戏的fps提高了30-40%，这样做的主要原因作者在文章描述如下：
for Replica Island I took a different approach. The problem with the single GLSurfaceView thread is that eglSwapBuffers() must block on the hardware until the previous frame finishes drawing. That means that even if you have nothing to draw, a call to eglSwapBuffers() takes 16.67ms to complete. (And of course, if you have a lot to draw, it could take a lot longer).
并且作者还在文中描述了一种使用双缓冲（double buffer ）的方法来解决线程间同步的问题。
这篇文章可以作为 优化 Android 游戏的一个参考。
Replica Island
http://replicaisland.blogspot.com
这个博客是Replica Island这个开源游戏的开发日志，其中包括了很多非常好的Android游戏开发的技术文章和游戏设计文章，作者是Chris Pruett，他在09 10界的Google I/O 会议上面都做了关于Android游戏开发的讲座：2009 http://www.youtube.com/watch?v=U4Bk5rmIpic  2010 http://www.youtube.com/watch?v=7-62tRHLcHk&feature=channel  （两个都需番羽墙）
