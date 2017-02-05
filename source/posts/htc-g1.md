{
    "title": "HTC G1入手",
    "author": "Chengzhi Yang",
    "createDate": "2011-12-06",
    "modifyDate": "2011-12-06",
    "permanent":"htc-g1"
}

# HTC G1入手

1220元 淘宝上面买了个G1的14天机。卖家第一次发给我的时候帮我刷了个Android 2.1的ROM，那个ROM极品至极，删除了 market, gmail, gtalk 做那个ROM的人怕是跟google 有血海深仇，把谷歌的系统软件都删得差不多。不过也可以理解为国内用户确实对google的服务需求不大，G1的ROM容量也小得可怜。卖家又不帮我刷SPL，只能寄回去叫卖家帮我刷了个1.6的ROM来用，用了半天，感觉很不错，gtalk，market, gmail都健在。

### 关于AndEngine的一个小Tips

我第一次拿到手机的时候运行我的新游戏，效果很好，我换货的过程中继续完善了游戏的其他部分，今天拿到手机，开始调试我的游戏的时候，发现游戏尽然非常的卡。调了G1 cpu频率无果，安装了一大堆游戏，感觉都不卡。然后开始排除法找代码，找了一会发现是 新建Texture时的 Texture参数的问题：

我有几个很大的.png图片，所以新建了几个Texture来存，其中最大的一个有 1024*512 ， 我发现参数为 TextureOptions.BILINEAR_PREMULTIPLYALPHA 就会卡，用TextureOptions.DEFAULT 就不卡了。 看来g1的像素处理能力实在太有限，二次线性插值大一点就顶不住。纹理参数的设置在手机平台上面还是比较敏感的，这可以作为优化AndEngine程序或其他Opengl ES程序的一个方法。
