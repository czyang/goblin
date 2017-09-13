{
    "title": "如何相对安全的保存iPhone游戏的分数",
    "author": "Chengzhi Yang",
    "createDate": "2011-08-29",
    "modifyDate": "2011-08-29",
    "permanent":"save-iphone-game-score-data-safely"
}

# 如何相对安全的保存iPhone游戏的分数

介绍一种保存游戏分数（高分）的方法，里面提供了相关的代码。这个方法可以相对安全的保存你的游戏分数，因为现在几乎所有移动平台的游戏都使用了OpenFeint，Game Center一类的社交游戏网络，其中使用最多的功能就是在线排行榜。我在我的两个iPhone游戏（[Voice Fighter](http://itunes.apple.com/us/app/voice-fighter/id449643608?ls=1&mt=8), [Magic Tiles](http://itunes.apple.com/cn/app/magic-tiles/id442161500?mt=8#)）中均使用了这个方法保存游戏的最高分数，防止游戏最高分数被越狱的用户通过修改游戏文件来作弊。但这个方法防不了那些通过截取网络数据包来改分数的手段。因为整个方法的核心就是加密分数然后保存在本地。

代码部分出自：[Cocos2d论坛](http://www.cocos2d-iphone.org/forum)

方法的过程是：

1. 需要保存游戏分数时，有游戏分数H，把H的数据类型转换成字符串：S，在S后加上字符串M作为混淆字符串，现在就得到字符串X(X=S+M).

2. 用MD5或者其他加密算法加密字符串X，得到一串密文，把这串密文和H保存在本地文件中.

3. 用户程序需要获取H时（比如游戏启动时需要在界面中显示最高分），读取保存密文和H的文件，这样就可以获得H和加密后的密文，现在就对密文进行解密，得到一个数字，用这个数字和H进行对比，如果得到的数字不等于H就证明用户修改过分数，可以选择清零数据；结果相等则通过检查。

```objectivec
#import <CommonCrypto/CommonDigest.h> 
//magicString既为前文的字符串M,可以任意修改这个字符串，不过要注意长度不要超过加密函数的上限
static NSString* magicString = @"passward$#@";
//MD5加密，也可以用其他加密方法，或者新写一个加密函数调用N次这个md5算法。
#pragma mark crypto
-(NSString*) md5:(NSString*) str {
    const char *cStr = [str UTF8String];
    unsigned char result[CC_MD5_DIGEST_LENGTH];
    CC_MD5( cStr, strlen(cStr), result );
    return [NSString stringWithFormat:
            @"%02X%02X%02X%02X%02X%02X%02X%02X%02X%02X%02X%02X%02X%02X%02X%02X",
            result[0], result[1], result[2], result[3], result[4], result[5], result[6], result[7],
            result[8], result[9], result[10], result[11], result[12], result[13], result[14], result[15]
            ];
}
```

```objectivec
//这段代码应该在"saveHighScore"保存分数一类的函数中
//highscore既是H,需要加密保存的分数,magicString同上。
NSString *md5Input = [NSString stringWithFormat:@"%d%@", highscore, magicString];
//cryptoString是X,之后只需要保存X即可。
NSString *cryptoString = [self md5:md5Input];
```

```objectivec
//这段代码应该在readHighScore一类的函数里
//highscore是从文件中读取的之前保存的最高分
//cryptoString是从文件中读取的加密字符串
    NSString *md5Input = [NSString stringWithFormat:@"%d%@", highscore, magicString];
    NSString *cryptoString = [self md5:md5Input];
    if ([cryptoString isEqualToString:mima] == NO) {
        //用户手动修改了最高分数，归零最高分数
    }
```

因为好多临时变量的功能其实是相同的，所以我用了相同的名字。

如果你在模拟器里面测试程序时不知到你保存了的文件保存在哪里。请点击 [iOS Simulator File System on Your Mac](http://developer.apple.com/library/ios/#documentation/Xcode/Conceptual/iphone_development/125-Using_iOS_Simulator/ios_simulator_application.html)
