{
    "title": "网络俄罗斯方块演示程序",
    "author": "Chengzhi Yang",
    "createDate": "2010-09-25",
    "modifyDate": "2010-09-25",
    "permanent":"network-tetris"
}

# 网络俄罗斯方块演示程序

这个程序是我学习socket套接字编程时的练习程序，程序用C++和DirectX编写，附件提供程序的vs2005项目文件。

这个程序是一个简单的可以两人联机的俄罗斯方块程序，没有记分和任何对战的设定，也就是这个程序只是一个原型，不能用来娱乐。

image文件夹下存放了图片素材文件。

Config.txt 中的第一行是服务端的ip地址，必须和Client.exe在同一级目录。

Client.exe 是客服端文件。

Server.exe 是服务端文件。

如果需要在单机执行演示程序的话，首先确保Config.txt中的ip地址必须是127.0.0.1，然后需要先打开Server.exe，然后打开 两个 Client.exe，这样就可以开始游戏了。

如果在不同的电脑上面运行程序的话，就需要修改客户机Config.txt文件中记录的ip地址，改成运行服务端的ip。

注意 bk.wav，是程序的背景文件，因为我用的素材太大，所以没有传上来，如果你运行程序出现问题，可以自己找一个wav文件来替代(具体的采样率同其他wav)。

[下载程序及其源码](../attachment/nettetris.zip)
