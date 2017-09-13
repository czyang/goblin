{
    "title": "Windows PC可以连VPN但是Mac和iOS设备不可以",
    "author": "Chengzhi Yang",
    "createDate": "2014-03-19",
    "modifyDate": "2014-03-19",
    "permanent":"windows-pc-vpn-mac-ios-cant-connect"
}

# Windows PC可以连VPN但是Mac和iOS设备不可以

遇到个很奇怪的问题，到了个新地方，拨号宽带，用无线路由器共享网络，Windows笔记本可以链接VPN，但是Mac OSX 和 iOS(iPhone, iPad) 都不能链. 用Mac通过有线直接拨号却又可以连，怀疑是路由器的问题，之后查资料发现可能是MTU的问题，随在路由器 WAN口设置（拨号设置）－高级设置 里面把MTU值从1480(TP-Link等路由器默认的)改成1492一切又恢复正常了。如果有人遇到类似的问题可以查阅一下MTU的设置，就应该可以解决这类的问题。
